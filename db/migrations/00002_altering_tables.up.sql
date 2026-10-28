BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER TABLE users DROP COLUMN IF EXISTS new_id;
ALTER TABLE users ADD COLUMN new_id UUID DEFAULT uuid_generate_v4();
alter table users drop constraint users_pkey cascade;
alter table users rename column user_id to old_int_id;
alter table users drop column if exists user_id;
alter table users rename column new_id to user_id ;
alter table users add primary key (user_id);
-- Clean up users table

ALTER TABLE Orders ADD COLUMN new_user_id UUID ;

update Orders 
set new_user_id = users.user_id
from users
where Orders.user_id  = users.old_int_id;

alter table Orders alter column new_user_id set not null;
alter table Orders drop column user_id;
alter table Orders rename column new_user_id to user_id;

ALTER TABLE users DROP COLUMN if exists old_int_id;

alter table Orders 
add constraint fk_orders_users
foreign key (user_id) references users(user_id) on delete cascade;
COMMIT;