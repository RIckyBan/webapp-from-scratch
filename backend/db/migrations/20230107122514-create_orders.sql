
-- +migrate Up
CREATE TABLE orders (
  id VARCHAR(26) PRIMARY KEY,
  user_id VARCHAR(26) NOT NULL,
  phone_id VARCHAR(26) NOT NULL,
  quantity INT NOT NULL,
  order_date DATE NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (phone_id) REFERENCES phones(id)
);

-- +migrate Down
DELETE TABLE orders;