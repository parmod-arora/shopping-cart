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

CREATE TABLE product_combo_packages ( 
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE product_combo_packaged_with ( 
    id BIGSERIAL PRIMARY KEY,
    product_combo_package_id BIGINT NOT NULL,
    product_id bigint NOT NULL REFERENCES products(id),
    packaged_with_product_id bigint NOT NULL REFERENCES products(id),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE product_discounts ( 
    id BIGSERIAL PRIMARY KEY,
    product_id bigint NOT NULL REFERENCES products(id),
    type text NOT NULL,
    min_quantity bigint NOT NULL,
    max_quantity bigint NOT NULL,
    combo_package_id bigint REFERENCES product_combo_packages (id),
    discount bigint,
    effective_start_date timestamp with time zone NOT NULL,
    effective_end_date timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);
