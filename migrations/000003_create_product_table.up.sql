CREATE SCHEMA "product"

CREATE TABLE "product"."products" (
    "id" BIGSERIAL PRIMARY KEY,
    "product_id" uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "amount" INTEGER NOT NULL, 
    "price" DECIMAL NOT NULL,
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz
);

INSERT INTO product.products (product_id,"name",amount,price,status,created_at,updated_at,deleted_at) VALUES
	 ('0f972ce0-9f6a-4fa2-83b7-c4aca617e935','Item 1',10,20.4,'active','2023-07-22 17:08:09.944','2023-07-22 17:08:09.944',NULL),
	 ('377b862d-b0cb-4299-9441-884d36153817','Item 2',11,15,'active','2023-07-22 17:08:09.954','2023-07-22 17:08:09.954',NULL),
	 ('7cb97b03-b086-4260-a68a-fa634eb6d10c','Item 3',0,14,'deleted','2023-07-22 17:08:09.957','2023-07-22 17:08:09.957','2023-07-22 17:07:53.062'),
	 ('a5f60303-e6b7-4afd-a68e-35abf76f5cf6','Item 4',2,100,'active','2023-07-22 17:08:09.960','2023-07-22 17:08:09.960',NULL);
