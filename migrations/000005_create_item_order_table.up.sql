CREATE TABLE "order"."order_items" (
    "id" BIGSERIAL PRIMARY KEY,
    "order_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "amount" INTEGER NOT NULL,
    "total" DECIMAL NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (order_id) REFERENCES "order"."orders" (order_id),
    FOREIGN KEY (product_id) REFERENCES "product"."products" (product_id)
);