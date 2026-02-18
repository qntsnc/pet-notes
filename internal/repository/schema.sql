CREATE TABLE notes (
  user_id BIGINT NOT NULL,
  id   BIGSERIAL PRIMARY KEY,
  title text     ,
  body  text
);
