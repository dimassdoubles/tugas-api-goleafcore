CREATE TABLE product (
	product_id              bigserial ,
    product_code            varchar(50),
	product_name            varchar(50),
    price                   numeric,
	version                 bigint,

	CONSTRAINT product_pkey PRIMARY KEY(product_id)
);