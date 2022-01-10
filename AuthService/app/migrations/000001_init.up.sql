CREATE TABLE users (
    id SERIAL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(60) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE brewery (
    id SERIAL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(60) NOT NULL,
    PRIMARY KEY (id)
);