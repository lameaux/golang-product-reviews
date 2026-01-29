-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    price INT NOT NULL
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL,
    review TEXT NOT NULL,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5)
);

CREATE INDEX idx_reviews_product_id ON reviews (product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reviews;

DROP TABLE products;
-- +goose StatementEnd