CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    details text NOT NULL,
	amount bigint NOT NULL,
    currency text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE stocks (
    id BIGSERIAL PRIMARY KEY,
    product_id bigint NOT NULL REFERENCES products(id),
    quantity bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users(id),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE product_combo_discount ( 
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    product_id bigint NOT NULL REFERENCES products(id),
    product_quantity bigint NOT NULL,
    product_quantity_fn text NOT NULL,
    discount_type text NOT NULL,
    discount bigint,
    packaged_with_product_id bigint NOT NULL REFERENCES products(id),
    packaged_with_product_quantity bigint NOT NULL,
    packaged_with_product_quantity_fn text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE product_discounts ( 
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    product_id bigint NOT NULL REFERENCES products(id),
    discount_type text NOT NULL,
    quantity bigint NOT NULL,
    quantity_fn text NOT NULL,
    discount bigint,
    effective_start_date timestamp with time zone NOT NULL,
    effective_end_date timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);
