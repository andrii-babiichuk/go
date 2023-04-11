CREATE TABLE IF NOT EXISTS orders
(
    uuid UUID PRIMARY KEY,
    date TIMESTAMP,
    description TEXT,
    price INT,
    status INT8
);