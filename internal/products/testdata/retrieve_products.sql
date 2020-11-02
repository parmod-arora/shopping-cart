DELETE from products;
INSERT INTO products("id","name","details","amount","currency","created_at","updated_at")
VALUES
(1,E'Apples',E'Apples Details',1000,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(2,E'Bananas',E'Bananas Details',200,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(3,E'Pears',E'Pears Details',300,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00'),
(4,E'Oranges',E'Oranges Details',100,E'SGD',E'2020-11-02 07:52:39+00',E'2020-11-02 07:52:39+00');