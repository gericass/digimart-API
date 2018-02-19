CREATE TABLE instruments (
  id   BIGINT UNSIGNED PRIMARY KEY,
  name VARCHAR(256) NOT NULL,
  category VARCHAR(64) NOT NULL,
  price INT NOT NULL,
  `condition` VARCHAR(64) NOT NULL,
  status BOOL NOT NULL,
  url TEXT NOT NULL,
  register_date DATETIME NOT NULL
);