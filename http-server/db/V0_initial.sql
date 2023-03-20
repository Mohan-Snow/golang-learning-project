CREATE table if not exists test_users
(
    user_id    serial PRIMARY KEY,
    user_name  VARCHAR(255) NOT NULL,
    created_on TIMESTAMP    NOT NULL,
    changed_on TIMESTAMP    NOT NULL
);