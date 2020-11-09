DELETE FROM "discount_rules";
DELETE FROM "coupons";
DELETE FROM "discounts";
DELETE FROM "cart_items";
DELETE from "products";
DELETE FROM "carts";
DELETE FROM "users";


INSERT INTO products("id","name","details","amount","currency","image")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD','apple.jpeg'),
(2, E'Bananas',E'Bananas Details',200,E'SGD','banana.jpg'),
(3, E'Pears',E'Pears Details',300,E'SGD','pears.jpg'),
(4, E'Oranges',E'Oranges Details',100,E'SGD','orange.jpeg');

INSERT INTO "users"("id","username","password","firstname","lastname","created_at","updated_at")
VALUES
(1,E'test',E'test',NULL,NULL,E'2020-11-03 04:17:31.084258+00',E'2020-11-03 04:17:31.084258+00'),
(2,E'test2',E'test',NULL,NULL,E'2020-11-03 04:17:31.084258+00',E'2020-11-03 04:17:31.084258+00');

-- Apples offer and Pear Banana discount entry
INSERT INTO discounts("id","name","discount_type","discount","created_at","updated_at")
VALUES
(1,E'Apple 10 % discount on 7 or more Apples',E'PERCENTAGE',10,E'2020-11-07 07:05:50.608799+00',E'2020-11-07 07:05:50.608799+00'),
(2,E'Combo discount on 4Pears and 2 Banana',E'PERCENTAGE',30,E'2020-11-07 07:13:59.881972+00',E'2020-11-07 07:13:59.881972+00');

-- Apple and Pears|banana discount rules
INSERT INTO discount_rules("id","product_id","product_quantity","product_quantity_fn","created_at","updated_at","discount_id")
VALUES
(1,1,7,E'GTE',E'2020-11-07 07:16:18.80056+00',E'2020-11-07 07:16:18.80056+00',1),
(2,3,4,E'EQ',E'2020-11-07 07:16:40.367592+00',E'2020-11-07 07:16:40.367592+00',2),
(3,2,2,E'EQ',E'2020-11-07 07:17:07.020461+00',E'2020-11-07 07:17:07.020461+00',2);

-- Orange Discount entry and rules
INSERT INTO discounts("id","name","discount_type","discount","created_at","updated_at")
VALUES (3,E'30% coupon discount on oranges',E'PERCENTAGE',30,E'2020-11-07 17:37:25.873064+00',E'2020-11-07 17:37:25.873064+00');
INSERT INTO discount_rules("id","discount_id","product_id","product_quantity","product_quantity_fn","created_at","updated_at")
VALUES (4,3,4,1,E'GTE',E'2020-11-07 17:48:39.443301+00',E'2020-11-07 17:48:39.443301+00');
