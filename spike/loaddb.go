package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    database, _ := sql.Open("sqlite3", "./superman")
    statement, _ := database.Prepare(`
      CREATE TABLE IF NOT EXISTS ips (
        id INTEGER PRIMARY KEY, 
        geoname_id TEXT, 
        registered_country_geoname_id INTEGER,
        represented_country_geoname_id INTEGER,
        postal_code INTEGER,
        latitude REAL,
        longitude REAL,
        accuracy_radius TEXT
      )
    `)
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO ips (geoname_id) VALUES (?)")
    statement.Exec(42)
    rows, _ := database.Query("SELECT id, geoname_id FROM ips")
    var id, geoname_id int
    for rows.Next() {
        rows.Scan(&id, &geoname_id)
        fmt.Printf("%d:%d", id, geoname_id)
    }
}
