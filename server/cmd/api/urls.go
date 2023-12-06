package main

import (
	"errors"
	"net/http"

	"github.com/Enthys/book-tracker/internal/data"
	"github.com/Enthys/book-tracker/internal/validator"
	"github.com/julienschmidt/httprouter"
)

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

func validateCreateShortUrlRequest(v *validator.Validator, req *CreateShortUrlRequest) {
	v.Check(req.Url != "", "url", "must be provided")
	v.Check(validator.ValidateUrl(req.Url), "url", "must be a valid url")
}

func (a *application) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req CreateShortUrlRequest
	if err := a.readJSON(w, r, &req); err != nil {
		a.badRequestError(w, err)
	}

	v := validator.New()
	if validateCreateShortUrlRequest(v, &req); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	shortUrl, err := a.models.ShortUrl.New(req.Url)
	if err != nil {
		a.serverError(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusCreated, envelope{"short_url": shortUrl}, nil); err != nil {
		a.serverError(w, r, err)
	}
}

func (a *application) RedirectToUrl(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortUrl := params.ByName("shortUrl")

	url, err := a.models.ShortUrl.GetByShortUrl(shortUrl)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundError(w, r)
		default:
			a.serverError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
}

type GetShortUrlsRequest struct {
	Page     int
	PageSize int
}

func validateGetShortUrlRequest(v *validator.Validator, req *GetShortUrlsRequest) {
	v.Check(req.Page > 0, "page", "must be greater than 0")
	v.Check(req.Page <= 10_000_000, "page", "must be a maximum of 10,000,000")

	v.Check(req.PageSize >= 0, "page_size", "must be greater than 0")
	v.Check(req.PageSize <= 100, "page_size", "must be a maximum of 100")
}

func (a *application) GetShortUrls(w http.ResponseWriter, r *http.Request) {
	var req GetShortUrlsRequest
	req.Page = readIntQueryParam(r, "page", 1)
	req.PageSize = readIntQueryParam(r, "page_size", 20)

	v := validator.New()
	if validateGetShortUrlRequest(v, &req); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	shortUrls, err := a.models.ShortUrl.GetAll(req.Page, req.PageSize)
	if err != nil {
		a.serverError(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"short_urls": shortUrls}, nil); err != nil {
		a.serverError(w, r, err)
	}
}

func (a *application) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	id, err := readIntRouteParam(r, "id")
	if err != nil {
		a.badRequestError(w, errors.New("invalid id parameter"))
		return
	}

	err = a.models.ShortUrl.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundError(w, r)
		default:
			a.serverError(w, r, err)
		}
		return
	}

	if err := a.writeJSON(w, http.StatusNoContent, envelope{}, nil); err != nil {
		a.serverError(w, r, err)
	}
}
