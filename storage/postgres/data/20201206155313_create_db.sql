-- +goose Up
CREATE TABLE IF NOT EXISTS proverbs (
	id SERIAL,
	maori_name TEXT,
	translation TEXT,
	explanation TEXT
);
CREATE TABLE IF NOT EXISTS placenames (
	id SERIAL,
	maori_name TEXT,
	translation TEXT,
	explanation TEXT
);

-- +goose Down
DROP TABLE proverbs;
DROP TABLE placenames;
