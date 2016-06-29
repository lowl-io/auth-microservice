CREATE TYPE current_status AS ENUM ('active', 'deleted', 'blocked');

CREATE TABLE Users
(
  id       SERIAL      NOT NULL PRIMARY KEY,
  name     VARCHAR(60) NOT NULL,
  password VARCHAR(60) NOT NULL,
  email    VARCHAR(255),
  status   current_status
);

INSERT INTO Users VALUES
  ('0', 'Alexander', 'password', 'example@gmail.com', 'active'),
  ('1', 'Dmitry', 'password1', 'example1@mail.ru', 'blocked'),
  ('2', 'Daniel', 'password2', 'example2@gmail.com', 'deleted');