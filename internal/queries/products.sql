-- name: getProductsList :many
SELECT * FROM products;

-- name: getProduct :one
SELECT * FROM products WHERE id = $1;
