-- Example queries for sqlc
CREATE TABLE authors
(
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    age integer NOT NULL,
    bio text
);

-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = $1
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM authors;
