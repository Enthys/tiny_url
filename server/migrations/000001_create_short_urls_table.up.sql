CREATE TABLE IF NOT EXISTS short_urls(
	id SERIAL PRIMARY KEY,
	short_url VARCHAR(255) NOT NULL,
	url TEXT NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS short_urls_short_url_idx ON short_urls USING BTREE(url);

