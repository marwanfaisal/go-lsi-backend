CREATE TABLE destination_products (
  id serial PRIMARY KEY,
  product_name TEXT,
	qty int,
  selling_price float,
  promo_price float,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);