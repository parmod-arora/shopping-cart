INSERT INTO products("id","name","details","amount","currency","image")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD','apple.jpeg'),
(2, E'Bananas',E'Bananas Details',200,E'SGD','banana.jpg'),
(3, E'Pears',E'Pears Details',300,E'SGD','pears.jpg'),
(4, E'Oranges',E'Oranges Details',100,E'SGD','orange.jpeg');

-- Apples offer
INSERT INTO "product_discounts"("id","name","discount_type","discount","created_at","updated_at")
VALUES
(1,E'Apple 10 % discount on 7 or more Apples',E'PERCENTAGE',10,E'2020-11-07 07:05:50.608799+00',E'2020-11-07 07:05:50.608799+00'),
(2,E'Combo discount on 4Pears and 2 Banana',E'PERCENTAGE',30,E'2020-11-07 07:13:59.881972+00',E'2020-11-07 07:13:59.881972+00');

INSERT INTO "product_discount_rules"("id","product_id","product_quantity","product_quantity_fn","created_at","updated_at","product_discount_id")
VALUES
(1,1,7,E'GTE',E'2020-11-07 07:16:18.80056+00',E'2020-11-07 07:16:18.80056+00',1),
(2,3,4,E'EQ',E'2020-11-07 07:16:40.367592+00',E'2020-11-07 07:16:40.367592+00',2),
(3,2,2,E'EQ',E'2020-11-07 07:17:07.020461+00',E'2020-11-07 07:17:07.020461+00',2);