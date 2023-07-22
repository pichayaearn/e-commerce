CREATE SCHEMA "order";

CREATE TABLE "order"."orders" (
    "id" BIGSERIAL PRIMARY KEY,
    "order_id" uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    "user_id" uuid NOT NULL,
    "total" DECIMAL NOT NULL,
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "canceled_at" timestamptz,

    FOREIGN KEY (user_id) REFERENCES "user"."users" (user_id)
);
