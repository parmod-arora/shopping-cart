CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    details TEXT NOT NULL,
	amount BIGINT NOT NULL,
    currency TEXT NOT NULL,
    "image" TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE stocks (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    reference text NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (user_id)
);
CREATE INDEX carts_user_id_idx ON carts(user_id);

CREATE TABLE cart_items (
    id BIGSERIAL PRIMARY KEY,
    cart_id BIGINT NOT NULL REFERENCES carts(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (cart_id, product_id)
);
CREATE INDEX cart_items_product_id_idx ON cart_items(product_id);
CREATE INDEX cart_items_cart_id_idx ON cart_items(cart_id);

CREATE TABLE discounts (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    discount_type text NOT NULL,
    discount bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE discount_rules (
    id BIGSERIAL PRIMARY KEY,
    discount_id bigint NOT NULL REFERENCES discounts(id),
    product_id bigint NOT NULL REFERENCES products(id),
    product_quantity bigint NOT NULL,
    product_quantity_fn text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE INDEX discount_rules_product_id_idx ON discount_rules(product_id);
CREATE INDEX discount_rules_discount_id_idx ON discount_rules(discount_id);
