CREATE TABLE "refresh_tokens" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "token" VARCHAR(255) UNIQUE NOT NULL,
    "expires_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    
    CONSTRAINT "fk_user_id"
    FOREIGN KEY ("user_id")
    REFERENCES "users" ("id")
    ON DELETE CASCADE
);
