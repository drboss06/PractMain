CREATE TABLE users
(
id serial not null unique,
name varchar(255) not null,
username varchar(255) not null unique,
password_hash varchar(255) not null
);

CREATE TABLE teams
(
id serial not null unique,
name varchar(255) not null,
team_description varchar(255)
);

CREATE TABLE users_team
(
id serial not null unique,
user_id int references users (id) on delete cascade not null,
team_id int references teams (id) on delete cascade not null
);