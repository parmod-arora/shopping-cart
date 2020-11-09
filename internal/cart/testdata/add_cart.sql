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