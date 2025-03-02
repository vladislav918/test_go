-- +goose Up
CREATE TABLE measures (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INT NOT NULL,
    unit_cost DECIMAL(10,2) NOT NULL,
    measure_id INT NOT NULL,
    FOREIGN KEY (measure_id) REFERENCES measures(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE products;
DROP TABLE measures;
