package core

import (
	"github.com/jmcvetta/randutil"
)

type UrlShortener struct {
	db     UrlShortenerDb
	idSize int
}

func NewUrlShortener(db UrlShortenerDb, idSize int) *UrlShortener {
	return &UrlShortener{db: db, idSize: idSize}
}

func (s *UrlShortener) Create(url string) (*ShortendUrl, error) {
	// TODO lock

	// check if already created
	shortened, err := s.db.SearchByUrl(url)
	if err != nil {
		return nil, err // TODO
	}
	if shortened != nil {
		return shortened, nil
	}

	// generate shortened url
	var id string
	for {
		id, err = randutil.AlphaString(s.idSize)
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
	shortened = &ShortendUrl{original: url, id: id}

	// register shortened url to db
	err = s.db.Register(shortened)
	if err != nil {
		return nil, err // TODO
	}

	return shortened, nil
}

func (s *UrlShortener) Resolve(id string) (*ShortendUrl, error) {
	shortened, err := s.db.Search(id)
	if err != nil {
		return nil, err // TODO
	}
	return shortened, err // shortened will be nil if not existing
}
