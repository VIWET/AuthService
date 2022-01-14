CREATE TABLE roles (
    id SERIAL,
    role_name VARCHAR(10) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE users (
    id SERIAL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(60) NOT NULL,
    role_id int,
    PRIMARY KEY (id),
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);

CREATE TABLE users_profiles (
  	id SERIAL,
	user_id INT NOT NULL UNIQUE,
  	PRIMARY KEY (id),
  	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE users_breweries (
  	id SERIAL,
	user_id INT NOT NULL UNIQUE,
  	PRIMARY KEY (id),
  	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO roles (role_name) VALUES ('admin'), ('user'), ('brewery');