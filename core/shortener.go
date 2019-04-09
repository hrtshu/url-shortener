package core

import (
	"errors"
)

type UrlShortener struct {
	db     UrlShortenerDb
	idSize int
}

func NewUrlShortener(db UrlShortenerDb, idSize int) *UrlShortener {
	return &UrlShortener{db: db, idSize: idSize}
}

func (s *UrlShortener) IdSize() int {
	return s.idSize
}

func (s *UrlShortener) Create(url string) (*ShortenedUrl, error) {
	// TODO lock

	// check if already created
	shortened, err := s.db.SearchByUrl(url)
	if err != nil {
		return nil, err // TODO
	}
	if shortened != nil {
		return shortened, nil
	}

	// check if url is invalid
	/* Note: The definition of URL correctness may change.
	   For this reason, validation is performed after DB search. */
	if !isValidUrl(url) {
		return nil, errors.New("invalid url")
	}

	// generate shortened url
	var id string
	for {
		id, err = generateId(s.idSize)
		if err != nil {
			return nil, err // TODO
		}
		shortened, err = s.db.Search(id)
		if err != nil {
			return nil, err // TODO
		}
		if shortened == nil {
			break
		}
	}
	shortened = NewShortenedUrl(url, id)

	// register shortened url to db
	err = s.db.Register(shortened)
	if err != nil {
		return nil, err // TODO
	}

	return shortened, nil
}

func (s *UrlShortener) Resolve(id string) (*ShortenedUrl, error) {
	shortened, err := s.db.Search(id)
	if err != nil {
		return nil, err // TODO
	}
	return shortened, err // shortened will be nil if not existing
}
