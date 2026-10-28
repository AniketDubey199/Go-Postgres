BEGIN;

ALTER table Orders drop constraint if exists fk_orders_users;

ALTER table users drop constraint users_pkey;
ALTER table users drop rename column id to to old_uuid;
ALTER table users add column id serial primary key;
ALTER table Orders add column old_users_id int;

update Orders set old_users_id = users.id
from users
where Orders.user_id = users.old_uuid

ALTER TABLE Orders ALTER old_users_id set not null;
ALTER TABLE Orders DROP column if exists user_id;
ALTER TABLE Orders RENAME column old_users_id to user_id;
ALTER DROP users column if exists old_uuid;

ALTER TABLE Orders add constraint Orders_users_id_fkey
foreign key (user_id) references users(id) on delete cascade;

COMMIT;