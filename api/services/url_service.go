package services

import (
	"database/sql"
	"errors"
	"fmt"
	"smol/core"
	"smol/models"
	"strconv"

	"github.com/jxskiss/base62"
)

type IUrlService interface {
	GetLongUrl(shortUrl string) (string, error)
	ShortenUrl(longUrl string) (string, error)
}

type urlService struct {
	app *core.App
}

var service IUrlService
var ErrInvalidURLID = errors.New("invalid URL ID")

func InitUrlService(app *core.App) IUrlService {
	if service == nil {
		service = &urlService{app}
	}

	return service
}

const (
	queryCreateURL = `
        INSERT INTO url (url) values ($1) RETURNING *`
	queryGetURLByID = `
        SELECT * FROM url WHERE id = $1`
	queryGetURLByLongURL = `
        SELECT * FROM url WHERE url = $1`
)

func (s *urlService) GetLongUrl(urlId string) (string, error) {
	db := s.app.DB
	actualId, err := decodeURLId(urlId)
	if err != nil {
		return "", fmt.Errorf("decode URL ID: %w", err)
	}

	var url models.URLModel
	if err := db.Get(&url, queryGetURLByID, actualId); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("URL not found: %w", err)
		}

		return "", fmt.Errorf("query URL: %w", err)
	}

	return url.Url, nil
}

func (s *urlService) ShortenUrl(longURL string) (string, error) {
	db := s.app.DB
	var url models.URLModel
	err := db.Get(&url, queryGetURLByLongURL, longURL)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("check existing URL: %w", err)
	}

	if err == sql.ErrNoRows {
		if err := db.Get(&url, queryCreateURL, longURL); err != nil {
			return "", fmt.Errorf("create URL: %w", err)
		}
	}

	encodedID := base62.StdEncoding.EncodeToString([]byte(fmt.Sprint(url.Id)))
	return encodedID, nil
}

func decodeURLId(urlID string) (int, error) {
	decodedBytes, err := base62.DecodeString(urlID)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidURLID, err)
	}

	actualId, err := strconv.Atoi(string(decodedBytes))
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidURLID, err)
	}

	return actualId, nil
}
