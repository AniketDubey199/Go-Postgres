BEGIN;

CREATE TABLE users (
  user_id serial primary key,
  name    varchar(40) not null,
  email   varchar(40) not null unique
);

CREATE TABLE Orders (
  order_id serial primary key,
  user_id int not null references users(user_id) on delete cascade,
  order_date Timestamp not null default current_timestamp,
  status varchar(20) not null default '*pending'
);

CREATE TABLE Products(
  product_id serial primary key,
  product_name varchar(40) not null,
  unit_price int not null ,
  stock_qty int not null default 0
);

CREATE TABLE order_items(
  order_item_id serial primary key,
  order_id int not null references Orders(order_id) on delete cascade,
  product_id int not null references Products(product_id) on delete restrict,
  quantity int not null default 1 check (quantity > 0),
  unit_price int not null
);

COMMIT;