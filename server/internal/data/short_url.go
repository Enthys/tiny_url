package data

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base32"
	"errors"
	"time"
)

type ShortUrlModel struct {
	DB *sql.DB
}

func (u ShortUrlModel) New(url string) (*ShortUrl, error) {
	shortUrl, err := generateShortUrl(url)
	if err != nil {
		return nil, err
	}

	return u.Insert(shortUrl)
}

// generateShortUrl generates a new short url for the given url. It uses the provided url as the seed for the random
// number generator. If the random number generator fails to generate a random number then an error is returned.
func generateShortUrl(url string) (*ShortUrl, error) {
	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	hash := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	shortUrl := string([]byte(hash[:20]))
	return &ShortUrl{
		Url:      url,
		ShortUrl: shortUrl,
	}, nil
}

// Insert inserts a new record into the short_urls table. If a conflict occurs then the existing record values will be
func (u ShortUrlModel) Insert(url *ShortUrl) (*ShortUrl, error) {
	stmt := `
		INSERT INTO short_urls (short_url, url)
			VALUES ($1, $2)
			On CONFLICT (url) DO UPDATE SET url = EXCLUDED.url
			RETURNING id, short_url, url
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	inserted := ShortUrl{}
	if err := u.DB.QueryRowContext(ctx, stmt, url.ShortUrl, url.Url).Scan(&inserted.ID, &inserted.ShortUrl, &inserted.Url); err != nil {
		return nil, err
	}

	return &inserted, nil
}

func (u ShortUrlModel) GetByShortUrl(shortUrl string) (*ShortUrl, error) {
	stmt := `
		SELECT id, url, short_url
		FROM short_urls
		WHERE short_url = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	url := ShortUrl{}
	err := u.DB.QueryRowContext(ctx, stmt, shortUrl).Scan(&url.ID, &url.Url, &url.ShortUrl)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &url, nil
}

type ShortUrl struct {
	ID       int    `json:"id"`
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

func (u ShortUrlModel) GetAll(page, pageSize int) ([]ShortUrl, error) {
	stmt := `
		SELECT id, url, short_url
		FROM short_urls
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, stmt, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	shortUrls := []ShortUrl{}
	for rows.Next() {
		url := ShortUrl{}
		if err := rows.Scan(&url.ID, &url.Url, &url.ShortUrl); err != nil {
			return nil, err
		}
		shortUrls = append(shortUrls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shortUrls, nil
}

func (u ShortUrlModel) Get(id int) (*ShortUrl, error) {
	stmt := `
		SELECT id, url, short_url
		FROM short_urls
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	url := ShortUrl{}
	err := u.DB.QueryRowContext(ctx, stmt, id).Scan(&url.ID, &url.Url, &url.ShortUrl)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &url, nil
}

func (u ShortUrlModel) Delete(id int) error {
	stmt := `
		DELETE FROM short_urls
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := u.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound

	}

	return nil
}
