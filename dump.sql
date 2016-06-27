CREATE TYPE current_status AS ENUM ('active', 'deleted', 'blocked');

CREATE TABLE Users
(
  id       SERIAL      NOT NULL PRIMARY KEY,
  name     VARCHAR(60) NOT NULL,
  password VARCHAR(64) NOT NULL,
  email    VARCHAR(255),
  status   current_status
);

