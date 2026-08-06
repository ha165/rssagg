-- +goose Up
CREATE TABLE feeds (
  ID UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  name TEXT NOT NULL,
);

--+goose Down
DROP TABLE feeds;