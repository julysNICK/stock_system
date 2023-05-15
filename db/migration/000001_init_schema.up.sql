CREATE TABLE "stores" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "contact_email" varchar NOT NULL,
  "contact_phone" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" decimal NOT NULL,
  "quantity" int NOT NULL,
  "store_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "suppliers" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "email" varchar NOT NULL,
  "contact_phone" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sales" (
  "id" int PRIMARY KEY,
  "product_id" int NOT NULL,
  "sale_date" date NOT NULL,
  "quantity_sold" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "stock_alerts" (
  "id" int PRIMARY KEY,
  "product_id" int,
  "supplier_id" int,
  "alert_quantity" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "products" ("store_id");

CREATE INDEX ON "sales" ("product_id");

CREATE INDEX ON "stock_alerts" ("product_id");

CREATE INDEX ON "stock_alerts" ("supplier_id");

CREATE INDEX ON "stock_alerts" ("product_id", "supplier_id");

ALTER TABLE "products" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "sales" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "stock_alerts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "stock_alerts" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");
