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

CREATE INDEX cart_items_cart_id_idx ON cart_items(cart_id);

CREATE TABLE product_combo_discount ( 
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    product_id BIGINT NOT NULL REFERENCES products(id),
    product_quantity BIGINT NOT NULL,
    product_quantity_fn TEXT NOT NULL,
    discount_type TEXT NOT NULL,
    discount BIGINT NOT NULL,
    packaged_with_product_id BIGINT NOT NULL REFERENCES products(id),
    packaged_with_product_quantity BIGINT NOT NULL,
    packaged_with_product_quantity_fn TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE product_discounts ( 
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    product_id BIGINT NOT NULL REFERENCES products(id),
    discount_type TEXT NOT NULL,
    quantity BIGINT NOT NULL,
    quantity_fn TEXT NOT NULL,
    discount BIGINT NOT NULL,
    effective_start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    effective_end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
