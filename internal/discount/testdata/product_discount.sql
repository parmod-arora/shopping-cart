DELETE from products;
INSERT INTO products("id","name","details","amount","currency","created_at","updated_at")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(2,E'Bananas',E'Bananas Details',200,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(3,E'Pears',E'Pears Details',300,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(4,E'Oranges',E'Oranges Details',100,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00');

DELETE from "product_discounts";
INSERT INTO "product_discounts"("id","name","product_id","quantity","quantity_fn","discount_type","discount","effective_start_date","effective_end_date","created_at","updated_at")
VALUES
(1,E'10% Discount on 7+ Apples',1,7,E'GTE',E'PERCENTAGE',10,E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00',E'2020-11-02 12:01:24.977785+00');

DELETE from "product_combo_discount";
INSERT INTO "product_combo_discount"("id","name","product_id","product_quantity","product_quantity_fn","discount_type","discount","packaged_with_product_id","packaged_with_product_quantity","packaged_with_product_quantity_fn","created_at","updated_at")
VALUES
(1,E'Banana + Pears',2,2,E'EQ',E'PERCENTAGE',30,3,3,E'EQ',E'2020-11-02 12:13:28.775947+00',E'2020-11-02 12:13:28.775947+00'),
(2,E'Banana + Oranges',4,4,E'EQ',E'PERCENTAGE',30,2,2,E'EQ',E'2020-11-02 12:13:28.775947+00',E'2020-11-02 12:13:28.775947+00');