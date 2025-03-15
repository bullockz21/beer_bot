-- +goose Up
CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    delivery_time TIMESTAMP,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_delivery_order_id ON delivery(order_id);

-- +goose Down
DROP TABLE delivery;
