
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

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE worklogs (
    worklog_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    task_id INT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    hours_spent INTERVAL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (task_id) REFERENCES tasks(task_id)
);