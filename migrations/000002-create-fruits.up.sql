-- +migrate Up
CREATE TABLE IF NOT EXISTS fruits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(50),
    image TEXT,
    quantity INT DEFAULT 0,
    price NUMERIC(10,2),
    description TEXT
);