CREATE TYPE "public"."gender" AS ENUM ('male', 'female', 'other');

CREATE TABLE "user"."users_profiles" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "display_name" varchar(20) NOT NULL COLLATE "pg_catalog"."default",
  "profile_image" varchar(200) DEFAULT NULL,
  "birthday" date DEFAULT NULL,
  "gender" "public"."gender" NOT NULL DEFAULT 'other',
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz,
  FOREIGN KEY (user_id) REFERENCES "user"."users" (user_id)
);