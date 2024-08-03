CREATE TABLE "orderes" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "address_id" INTEGER NOT NULL,
  "total_amount" DECIMAL(10, 2) NOT NULL,
  "status" VARCHAR(100) NOT NULL DEFAULT 'pending',
  "payment_method" VARCHAR(100),
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP ,

  CONSTRAINT "fk_user_id"
    FOREIGN KEY ("user_id")
    REFERENCES "users" ("id")
    ON DELETE CASCADE,
  CONSTRAINT "fk_address_id"
    FOREIGN KEY ("address_id")
    REFERENCES "addresses" ("id")
    ON DELETE CASCADE
);