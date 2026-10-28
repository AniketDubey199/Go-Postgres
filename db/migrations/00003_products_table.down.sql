BEGIN;

ALTER table order_items drop constraint if exists fk_order_items_products;

ALTER table Products drop constraint if exists Products_pkey;
ALTER table Products rename column product_id to old_product_uuid;
ALTER table Products add column product_id serial primary key;
ALTER table order_items add column old_p_int_id int;

update order_items set old_p_int_id = products.product_id
from products
where order_items.product_id = Products.old_product_uuid;

ALTER table Products drop column if exists old_product_uuid;
ALTER table order_items alter old_p_int_id set not null;
ALTER table order_items drop column if exists product_id;
ALTER table order_items rename column old_p_int_id to product_id;

ALTER table order_items add constraint order_items_products_products_id
foreign key (product_id) references Products(product_id) on delete cascade;

COMMIT;