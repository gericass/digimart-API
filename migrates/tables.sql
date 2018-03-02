-- +migrate Up
CREATE TABLE instruments (
  `id`            BIGINT UNSIGNED       AUTO_INCREMENT PRIMARY KEY,
  `name`          VARCHAR(256) NOT NULL,
  `category`      VARCHAR(64),
  `price`         INT          NOT NULL,
  `condition`     VARCHAR(64)  NOT NULL,
  `status`        BOOL         NOT NULL,
  `url`           TEXT         NOT NULL,
  `image`         TEXT         NOT NULL,
  `description`   TEXT,
  `register_date` DATETIME     NOT NULL,
  `created_at`    TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
  id         BIGINT UNSIGNED      AUTO_INCREMENT PRIMARY KEY,
  name       VARCHAR(26) NOT NULL,
  password   VARCHAR(26) NOT NULL,
  salt       VARCHAR(6)  NOT NULL,
  created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP   NOT NULL DEFAULT NOW()
);

-- instrumentsとusersの中間テーブル
CREATE TABLE subscribes (
  id            BIGINT UNSIGNED          AUTO_INCREMENT PRIMARY KEY,
  user_id       BIGINT UNSIGNED NOT NULL,
  instrument_is BIGINT UNSIGNED NOT NULL,
  created_at    TIMESTAMP       NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE instruments;
DROP TABLE users;
DROP TABLE subscribes;