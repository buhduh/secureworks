/*
  insert or ignore into users (name) values ("foo");

  insert into events(user_id, event_uuid, unix_timestamp, ip_address, latitude, longitude, radius) select id, 'adsfa', 123, 'adsfa', 23.3, 323.4, 100 from users where name = 'foo';

  select 
    e.ip_address, e.unix_timestamp, u.name, e.event_uuid, 
    e.latitude, e.longitude, e.radius from events e 
  join users u on e.user_id=u.id 
  where name = 'bob' 
  order by e.unix_timestamp desc;

*/

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY, 
  user_id INTEGER,
	event_uuid TEXT UNIQUE,
	unix_timestamp INTEGER,
	ip_address TEXT,
	latitude REAL,
	longitude REAL,
	radius INTEGER,
  FOREIGN KEY(user_id) REFERENCES users(id)
);
