
-- +migrate Up
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    admin BOOLEAN default false NOT NULL
);

-- +migrate Down
DROP TABLE users;
