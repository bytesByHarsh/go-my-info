-- +goose Up
ALTER TABLE users ADD COLUMN is_active bool NOT NULL DEFAULT TRUE;

CREATE TYPE bank_type AS ENUM(
    'central', 'cooperative', 'commercial',
    'regional', 'local', 'specialized',
    'small_finance', 'payments'
);

CREATE TABLE banks (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    name VARCHAR(100) NOT NULL UNIQUE,
    contact_phone VARCHAR(15),
    contact_email VARCHAR(255) UNIQUE NOT NULL,
    address TEXT,
    type bank_type NOT NULL,
    established_year INT NOT NULL
);

CREATE TYPE bank_account_type AS ENUM('savings', 'checking');
CREATE TABLE bank_accounts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    bank_id UUID NOT NULL REFERENCES banks(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    account_number VARCHAR(50) UNIQUE NOT NULL,
    account_type bank_account_type NOT NULL,
    balance NUMERIC(20, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL
);

CREATE TYPE card_type AS ENUM('credit', 'debit');
CREATE TABLE cards (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bank_id UUID NOT NULL REFERENCES banks(id) ON DELETE CASCADE,
    bank_account_id UUID REFERENCES bank_accounts(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    nickname VARCHAR(100) NOT NULL,
    number VARCHAR(19) UNIQUE NOT NULL,
    type card_type NOT NULL,
    expiration_date DATE NOT NULL,
    cvv VARCHAR(4) NOT NULL,
    total_limit NUMERIC(20, 2) NOT NULL,
    bill_date DATE NOT NULL
);

-- +goose Down
DROP TABLE cards;
DROP TYPE card_type;
DROP TABLE bank_accounts;
DROP TYPE bank_account_type;
DROP TABLE banks;
DROP TYPE bank_type;
ALTER TABLE users DROP COLUMN is_active;