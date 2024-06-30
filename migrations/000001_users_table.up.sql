CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    passport_series TEXT NOT NULL,
    passport_number TEXT NOT NULL,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100),
    address TEXT,
    UNIQUE (passport_series, passport_number)
);