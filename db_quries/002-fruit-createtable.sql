CREATE TABLE fruits (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(50),
    image VARCHAR(255),
    quantity INT DEFAULT 0,
    price NUMERIC(10,2) NOT NULL,
    description TEXT
);
