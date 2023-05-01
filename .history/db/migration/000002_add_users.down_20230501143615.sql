    ALTER TABLE IF EXISTS "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency"); 
    DROP TABLE IF EXISTS users;