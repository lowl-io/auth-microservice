CREATE TYPE status AS ENUM ('active', 'deleted', 'blocked');

CREATE TABLE Users
(
  id       SERIAL       NOT NULL PRIMARY KEY,
  email    VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(60)  NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  name     VARCHAR(60)  NOT NULL,
  status   current_status
);
