CREATE TYPE status AS ENUM ('active', 'deleted', 'blocked');

CREATE TABLE Users
(
  id       SERIAL       NOT NULL PRIMARY KEY,
  name     VARCHAR(60)  NOT NULL UNIQUE,
  password VARCHAR(60)  NOT NULL,,
  email    VARCHAR(255) UNIQUE,
  name     VARCHAR(60)  NOT NULL,
  status   current_status
);
