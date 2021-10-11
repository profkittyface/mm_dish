CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    username text,
    password text,
    email text,
    first_name text,
    last_name text,
    lastip text,
    last_login timestamp with time zone,
    date_created timestamp with time zone DEFAULT now(),
    last_updated timestamp with time zone DEFAULT now()
);

CREATE TABLE IF NOT EXISTS location
(
    id serial PRIMARY KEY,
    location geometry,
    userid integer NOT NULL,
    last_updated timestamp with time zone DEFAULT now(),
    CONSTRAINT location_userid FOREIGN KEY (userid) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS venue
(
    id serial PRIMARY KEY,
    name text,
    address text,
    hours text,
    date_created timestamp with time zone DEFAULT now(),
    last_updated timestamp with time zone DEFAULT now()
);

CREATE TABLE IF NOT EXISTS event
(
    id serial PRIMARY KEY,
    name text,
    venueid int,
    meeting_time timestamp with time zone,
    attended text,
    date_created timestamp with time zone DEFAULT now(),
    last_updated timestamp with time zone DEFAULT now(),
    CONSTRAINT event_venueid FOREIGN KEY (venueid) REFERENCES venue(id)
);

CREATE TABLE IF NOT EXISTS location_key
(
    id serial PRIMARY KEY,
    userid int,
    key text,
    date_created timestamp with time zone DEFAULT now(),
    last_updated timestamp with time zone DEFAULT now(),
    CONSTRAINT location_key_userid FOREIGN KEY (userid) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS auth_key
(
    id serial PRIMARY KEY,
    userid int,
    cookie_key text,
    expires timestamp with time zone,
    CONSTRAINT auth_key_userid FOREIGN KEY (userid) REFERENCES users(id)
);

insert into users (username,password,email,first_name,last_name,lastip,last_login) values ('ahunt','d9251693dfcb10694f3eaf9c7c4cfafbac33104688a947a7956e647be10e8e0c','ahunt@ahunt.com','Andrew','Hunt','1.1.1.1',now());
insert into venue (name,address,hours) values ('The Club of Phil','500 Broadway Avenue, San Francisco, CA 94103','9-5');
insert into location (location,userid) values (st_makepoint(37.7646207,-122.4127467), 1);
insert into event (name,venueid,meeting_time,attended) values ('October Foxtrot',1,'2021-10-02 00:00:00-07','1');
insert into location_key (userid,key) values (1,'d79b4bc10a3f1fcda1645c5ab4f9ff41b4135d61');
