
-- +migrate Up
CREATE TABLE phones (
  id VARCHAR(26) PRIMARY KEY,
  brand VARCHAR(255) NOT NULL,
  model VARCHAR(255) NOT NULL,
  operating_system VARCHAR(255) NOT NULL,
  storage_gb INT NOT NULL,
  ram_gb INT NOT NULL,
  color VARCHAR(255) NOT NULL,
  screen_size_inch DECIMAL(5, 1) NOT NULL,
  weight_g DECIMAL(5, 1) NOT NULL,
  price INT NOT NULL,
  release_date DATE NOT NULL,
  description TEXT
);

-- +migrate Down
DROP TABLE phones;
