
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
