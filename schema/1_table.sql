CREATE TABLE IF NOT EXISTS loan (
    id SERIAL PRIMARY KEY,
    borrower_id BIGINT NOT NULL,
    amount FLOAT NOT NULL,
    rate FLOAT NOT NULL,
    approval_proof_url VARCHAR NOT NULL,
    agreement_letter_url VARCHAR NOT NULL,
    status INT NOT NULL,
    created_by BIGINT NOT NULL,
    approved_by BIGINT NOT NULL,
    disbursed_by BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    approved_at TIMESTAMP,
    invested_at TIMESTAMP,
    disbursed_at TIMESTAMP
);
CREATE INDEX idx_loan_status ON loan(status);

CREATE TABLE IF NOT EXISTS investment (
    id SERIAL PRIMARY KEY,
    investor_id BIGINT NOT NULL,
    loan_id BIGINT NOT NULL,
    amount FLOAT NOT NULL,
    roi FLOAT NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE INDEX idx_investment_loan_id ON investment(loan_id);

CREATE TABLE IF NOT EXISTS employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS borrower (
    id SERIAL PRIMARY KEY,
    identification_number VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS investor (
    id SERIAL PRIMARY KEY,
    identification_number VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);