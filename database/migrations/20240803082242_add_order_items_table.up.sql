CREATE TABLE "order_items" (
  "id" SERIAL PRIMARY KEY,
  "order_id" INTEGER NOT NULL,
  "product_id" INTEGER NOT NULL,
  "quantity" INTEGER NOT NULL DEFAULT 1,
  "price_at_order" DECIMAL(10, 2) NOT NULL,

  CONSTRAINT "fk_order_id"
    FOREIGN KEY ("order_id")
    REFERENCES "orderes" ("id")
    ON DELETE CASCADE,
  CONSTRAINT "fk_product_id"
    FOREIGN KEY ("product_id")
    REFERENCES "products" ("id")
    ON DELETE CASCADE
);