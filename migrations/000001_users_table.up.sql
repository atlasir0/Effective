CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    encrypted_passport_series TEXT NOT NULL,
    encrypted_passport_number TEXT NOT NULL,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100),
    address TEXT,
    UNIQUE (encrypted_passport_series, encrypted_passport_number)
);