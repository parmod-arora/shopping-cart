DELETE from "products";
INSERT INTO "products"("id","name","details","amount","currency","created_at","updated_at")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(2,E'Bananas',E'Bananas Details',200,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(3,E'Pears',E'Pears Details',300,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(4,E'Oranges',E'Oranges Details',100,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00');

DELETE FROM "users";
INSERT INTO "users"("id","username","password","firstname","lastname","created_at","updated_at")
VALUES
(1,E'test',E'test',NULL,NULL,E'2020-11-03 04:17:31.084258+00',E'2020-11-03 04:17:31.084258+00'),
(2,E'test2',E'test',NULL,NULL,E'2020-11-03 04:17:31.084258+00',E'2020-11-03 04:17:31.084258+00');

DELETE from "product_discounts";
INSERT INTO "product_discounts"("id","name","product_id","quantity","quantity_fn","discount_type","discount","effective_start_date","effective_end_date","created_at","updated_at")
VALUES
(1,E'10% Discount on 7+ Apples',1,7,E'GTE',E'PERCENTAGE',10,E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00');


DELETE FROM "carts";
INSERT INTO "carts"("id","user_id","reference","created_at","updated_at")
VALUES
(1,1,E'1',E'2020-11-03 04:18:23.683425+00',E'2020-11-03 04:18:23.683425+00'),
(2,2,E'2',E'2020-11-03 04:18:23.683425+00',E'2020-11-03 04:18:23.683425+00');

DELETE FROM "cart_items";
INSERT INTO "cart_items"("id","cart_id","product_id","quantity","created_at","updated_at")
VALUES
(1,2,1,10,E'2020-11-03 04:19:02.694029+00',E'2020-11-03 04:19:02.694029+00');

-- conmbo rows
INSERT INTO "product_combo_discount"("id","name","product_id","product_quantity","product_quantity_fn","discount_type","discount","packaged_with_product_id","packaged_with_product_quantity","packaged_with_product_quantity_fn","created_at","updated_at")
VALUES
(1,E'4 Pears +  2 Banana - 30 %',3,4,E'EQ',E'PERCENTAGE',30,2,2,E'EQ',E'2020-11-03 14:08:22.942855+00',E'2020-11-03 14:08:22.942855+00');

INSERT INTO "cart_items"("id","cart_id","product_id","quantity","created_at","updated_at")
VALUES
(3,2,2,11,E'2020-11-03 14:13:10.275312+00',E'2020-11-03 14:13:10.275312+00'),
(4,2,3,21,E'2020-11-03 14:13:10.275312+00',E'2020-11-03 14:13:10.275312+00');
