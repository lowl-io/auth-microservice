CREATE TYPE status AS ENUM ('active', 'deleted', 'blocked');

CREATE TABLE Users
(
  id             SERIAL NOT NULL PRIMARY KEY,
  name           VARCHAR(60),
  email          VARCHAR(255),
  password       VARCHAR(64),
  current_status status
);

