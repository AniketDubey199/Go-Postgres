BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER table Products drop constraint if exists Products_pkey cascade;
ALTER table Products add column new_product_uuid UUID default uuid_generate_v4();
ALTER table Products rename column product_id to old_product_id ;
ALTER table Products rename column new_product_uuid to product_id;
ALTER table Products add primary key (product_id);

ALTER table order_items add column new_p_id UUID;

update order_items
set new_p_id = products.product_id
from products
where order_items.product_id = products.old_product_id;

ALTER TABLE order_items alter column new_p_id set not null; 
ALTER TABLE order_items drop column if exists product_id;
ALTER TABLE Products drop column if exists old_product_id;
ALTER TABLE order_items rename column new_p_id to product_id;

alter table order_items 
add constraint fk_order_items_products
foreign key (product_id) references products(product_id) on delete cascade;

COMMIT;