-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS stores_id_seq;

-- Table Definition
CREATE TABLE "public"."stores" (
    "id" int4 NOT NULL DEFAULT nextval('stores_id_seq'::regclass),
    "name" varchar(63),
    "address" varchar(255),
    "latitude" numeric(13,10),
    "longitude" numeric(13,10),
    "total_sales" int4,
    "rating" int4,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" text,
    "password" text,
    "deposit_amount" numeric,
    PRIMARY KEY ("id")
);

