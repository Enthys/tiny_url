package data

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUpdateConflict = errors.New("update conflict")
)

type Models struct {
	ShortUrl     ShortUrlModel
	ClientVisits ClientVisitModel
}
