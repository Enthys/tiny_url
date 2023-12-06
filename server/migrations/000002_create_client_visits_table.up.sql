CREATE TABLE client_visits (
	ip VARCHAR(15) NOT NULL,
	request_url TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
