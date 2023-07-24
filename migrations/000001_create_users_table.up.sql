CREATE SCHEMA "user";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "user"."users" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "email" VARCHAR(255) NOT NULL,
    "password" varchar(200) NOT NULL, 
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,

    CONSTRAINT "uniq_user_id" UNIQUE ("user_id")
);
