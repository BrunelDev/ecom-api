-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
id BIGSERIAL PRIMARY KEY,
name TEXT NOT NULL,
price_in_cent INT NOT NULL CHECK (price_in_cent >= 0),
quantity INT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
