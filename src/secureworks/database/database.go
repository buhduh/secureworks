package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"secureworks/constants"
	"secureworks/event"
)

const (
	//The "correct" approach for this would be to verify error codes
	DUPE_ERROR string = "UNIQUE constraint failed: events.event_uuid"
)

type Database interface {
	NewEvent(*event.Event) error
	//ordered by timestamp ascending, oldest at index 0
	GetOrderedEventsForUser(string) ([]*event.Event, error)
}

type realDatabase struct {
	db *sql.DB
}

func (r *realDatabase) NewEvent(e *event.Event) error {
	statement, err := r.db.Prepare(`
    INSERT OR IGNORE INTO users (name)
    VALUES (?)
  `)
	if err != nil {
		return err
	}
	_, err = statement.Exec(e.UserName)
	if err != nil {
		return err
	}
	statement, err = r.db.Prepare(`
	    INSERT INTO events(
	      user_id, event_uuid, unix_timestamp,
	      ip_address, latitude, longitude, radius
	    )
	    SELECT
	      id, ?, ?, ?, ?, ?, ?
	    FROM users
	    WHERE name = ?
	  `)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		e.EventUUID, e.Timestamp, e.IP,
		e.Location.Latitude, e.Location.Longitude,
		e.Location.Radius, e.UserName,
	)

	if err != nil && err.Error() != DUPE_ERROR {
		return err
	}
	return nil
}

func (r *realDatabase) GetOrderedEventsForUser(name string) ([]*event.Event, error) {
	statement, err := r.db.Prepare(`
    SELECT 
      e.ip_address, e.unix_timestamp, u.name, e.event_uuid, 
      e.latitude, e.longitude, e.radius from events e 
    JOIN users u ON e.user_id=u.id 
    WHERE name = ?
    ORDER BY e.unix_timestamp ASC;
  `)
	if err != nil {
		return nil, err
	}
	rows, err := statement.Query(name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	//why can't rows have a length?
	toRet := make([]*event.Event, 0)
	var ip, uName, eUUID string
	var lat, lon float64
	var rad, time uint
	for rows.Next() {
		rows.Scan(&ip, &time, &uName, &eUUID, &lat, &lon, &rad)
		newEvent, err := event.NewEvent(
			ip, name, eUUID, time, rad, lat, lon,
		)
		if err != nil {
			return nil, err
		}
		toRet = append(toRet, newEvent)
	}
	return toRet, nil
}

func NewDatabase() (Database, error) {
	database, err := sql.Open("sqlite3", constants.SQL_DB)
	if err != nil {
		return nil, err
	}
	return &realDatabase{
		db: database,
	}, nil
}
