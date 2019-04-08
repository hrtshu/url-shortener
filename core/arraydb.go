package core

import (
	"errors"
)

type UrlShortenerArrayDb struct {
	urls  []*ShortendUrl
	count int
}

func NewUrlShortenerArrayDb(size int) *UrlShortenerArrayDb {
	urls := make([]*ShortendUrl, size)
	return &UrlShortenerArrayDb{urls: urls}
}

func (d *UrlShortenerArrayDb) Register(shortened *ShortendUrl) error {
	if d.count >= len(d.urls) {
		return errors.New("reached limit of url registrations")
	}
	shortened_ := *shortened // copy
	d.urls[d.count] = &shortened_
	d.count++
	return nil
}

func (d *UrlShortenerArrayDb) Search(id string) (*ShortendUrl, error) {
	for i := 0; i < d.count; i++ {
		if d.urls[i].id == id {
			shortened := *(d.urls[i]) // copy
			return &shortened, nil
		}
	}
	return nil, nil
}

func (d *UrlShortenerArrayDb) SearchByUrl(original string) (*ShortendUrl, error) {
	for i := 0; i < d.count; i++ {
		if d.urls[i].original == original {
			shortened := *(d.urls[i]) // copy
			return &shortened, nil
		}
	}
	return nil, nil
}
