BEGIN;
alter table Orders drop constraint if exists fk_orders_users;
alter table c
alter table users drop constraint users_pkey;
ALTER TABLE users ADD COLUMN old_id SERIAL;
ALTER TABLE users DROP COLUMN IF EXISTS id;
ALTER TABLE users RENAME COLUMN old_id TO id;
alter table users add primary key (id);
drop extension if exists "uuid-ossp";

COMMIT;1