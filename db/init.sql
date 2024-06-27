CREATE SCHEMA userms;

CREATE TABLE userms.user
(
    uid char(36) NOT NULL,
    username varchar(64),
    first_name varchar(32),
    last_name varchar(32),
    email varchar(320) UNIQUE,
    phone_number varchar(15) UNIQUE,
    gender char,

    PRIMARY KEY (uid)
);
