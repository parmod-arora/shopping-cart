CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    reference text NOT NULL,
    total_amount bigint NOT NULL,
    sub_total_amount bigint NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX orders_user_id_idx ON orders(user_id);

CREATE TABLE order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE order_coupons (
    id BIGSERIAL PRIMARY KEY,
    order_id bigint NOT NULL REFERENCES orders(id),
    coupon_id bigint NOT NULL REFERENCES coupons(id),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE order_discount (
    id BIGSERIAL PRIMARY KEY,
    order_id bigint NOT NULL REFERENCES orders(id),
    discount_id bigint NOT NULL REFERENCES discounts(id),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);
