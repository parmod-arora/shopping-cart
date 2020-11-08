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
