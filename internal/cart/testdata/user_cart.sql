DELETE FROM "discount_rules";
DELETE FROM "coupons";
DELETE FROM "discounts";
DELETE FROM "cart_items";
DELETE from "products";
DELETE FROM "carts";
DELETE FROM "users";

INSERT INTO products("id","name","details","amount","currency","image","created_at","updated_at")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD','apple.jpeg',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(2, E'Bananas',E'Bananas Details',200,E'SGD','banana.jpg',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(3, E'Pears',E'Pears Details',300,E'SGD','pears.jpg',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(4, E'Oranges',E'Oranges Details',100,E'SGD','orange.jpeg',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00');

INSERT INTO "users"("id","username","password","firstname","lastname","created_at","updated_at")
VALUES
(1,E'test',E'test',NULL,NULL,E'2020-11-03 04:17:31+00',E'2020-11-03 04:17:31+00'),
(2,E'test2',E'test',NULL,NULL,E'2020-11-03 04:17:31+00',E'2020-11-03 04:17:31+00');

INSERT INTO "carts"("id","user_id","reference","created_at","updated_at")
VALUES
(1,1,E'1',E'2020-11-03 04:18:23+00',E'2020-11-03 04:18:23+00'),
(2,2,E'2',E'2020-11-03 04:18:23+00',E'2020-11-03 04:18:23+00');

INSERT INTO "cart_items"("id","cart_id","product_id","quantity","created_at","updated_at")
VALUES
(2,2,1,8,E'2020-11-07 11:21:39+00',E'2020-11-07 11:21:39+00'),
(3,2,2,5,E'2020-11-07 11:21:51+00',E'2020-11-07 11:22:02+00'),
(4,2,3,9,E'2020-11-07 11:21:52+00',E'2020-11-07 11:22:07+00'),
(5,2,4,1,E'2020-11-07 11:21:54+00',E'2020-11-07 11:21:54+00');

-- Apples offer
INSERT INTO "discounts"("id","name","discount_type","discount","created_at","updated_at")
VALUES
(1,E'Apple 10 discount on 7 or more Apples',E'PERCENTAGE',10,E'2020-11-07 07:05:50+00',E'2020-11-07 07:05:50+00'),
(2,E'Combo discount on 4Pears and 2 Banana',E'PERCENTAGE',30,E'2020-11-07 07:13:59+00',E'2020-11-07 07:13:59+00'),
(3,E'Coupon discount on oranges 30%',E'PERCENTAGE',30,E'2020-11-07 17:37:25+00',E'2020-11-07 17:37:25+00');

INSERT INTO "discount_rules"("id","discount_id","product_id","product_quantity","product_quantity_fn","created_at","updated_at")
VALUES
(1,1,1,7,E'GTE',E'2020-11-07 07:16:18+00',E'2020-11-07 07:16:18+00'),
(2,2,3,4,E'EQ',E'2020-11-07 07:16:40+00',E'2020-11-07 07:16:40+00'),
(3,2,2,2,E'EQ',E'2020-11-07 07:17:07+00',E'2020-11-07 07:17:07+00'),
(4,3,4,1,E'GTE',E'2020-11-07 17:48:39+00',E'2020-11-07 17:48:39+00');

INSERT INTO "coupons"("id","name","product_id","discount_id","expire_at","redeemed_at")
VALUES
(1,E'COUPON_30',4,3,E'2020-11-08 07:22:54+00',NULL);

INSERT INTO "cart_coupons"("id","cart_id","coupon_id","created_at","updated_at")
VALUES
(1,2,1,E'2020-11-08 09:21:44+00',E'2020-11-08 09:21:44+00');