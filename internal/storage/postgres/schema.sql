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

CREATE TABLE worklogs (
    worklog_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    hours_spent INTERVAL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
