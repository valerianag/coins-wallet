BEGIN;
CREATE TABLE accounts
(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    balance NUMERIC NOT NULL CHECK(balance >= 0),
    currency TEXT NOT NULL
);
CREATE TABLE payments
(
    account_from integer NOT NULL,
    account_to integer NOT NULL,
    amount NUMERIC NOT NULL CHECK(amount > 0),
    CHECK(account_to != account_from),
    CONSTRAINT fk_account_from
      FOREIGN KEY(account_from)
	  REFERENCES accounts(id),
    CONSTRAINT fk_account_to
      FOREIGN KEY(account_to)
	  REFERENCES accounts(id)
);
COMMIT;
