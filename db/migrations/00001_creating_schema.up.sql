CREATE TABLE users (
  user_id serial primary key,
  name    varchar(40) not null,
  email   varchar(40) not null
);