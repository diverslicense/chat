CREATE SCHEMA chatschema;

CREATE TABLE IF NOT EXISTS chatschema.users (
	userid serial PRIMARY KEY,
	username varchar(30) not null, 
	created timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chatschema.rooms (
	roomid serial PRIMARY KEY,
	roomname varchar(40),
	created timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chatschema.roommembers (
	roomid integer,
	userid integer
);

CREATE TABLE IF NOT EXISTS chatschema.messages (
	messageid serial PRIMARY KEY,
	messagebody text,
	roomid integer,
	senderid integer,
	created timestamp default CURRENT_TIMESTAMP
);