    ALTER TABLE IF EXISTS "accounts" DROP "owner_currency_key" UNIQUE ("owner", "currency"); 
    DROP TABLE IF EXISTS users;