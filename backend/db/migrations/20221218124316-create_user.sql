
-- +migrate Up
CREATE TABLE users (
  id VARCHAR(26) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE users;
