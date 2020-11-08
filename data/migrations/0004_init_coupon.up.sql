CREATE TABLE coupons (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    product_id bigint NOT NULL REFERENCES products(id),
    discount_id bigint NOT NULL REFERENCES discounts(id),
    expire_at timestamp with time zone NOT NULL,
    redeemed_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    UNIQUE (name)
);

CREATE TABLE cart_coupons (
    id BIGSERIAL PRIMARY KEY,
    cart_id bigint NOT NULL REFERENCES carts(id),
    coupon_id bigint NOT NULL REFERENCES coupons(id),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE INDEX cart_coupons_cart_id_idx ON cart_coupons(cart_id);

INSERT INTO "coupons"("id","name","product_id","discount_id","expire_at","redeemed_at")
VALUES
(1,E'COUPON_30',4,3,E'2021-11-08 17:30:53.717317+00',NULL),
(2,E'EXPIRE',4,3,E'2020-11-08 17:30:53.717317+00',NULL);