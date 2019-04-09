package db

import (
	"errors"
	"github.com/hrtshu/url-shortener/core"
)

type UrlShortenerArrayDb struct {
	urls  []*core.ShortenedUrl
	count int
}

func NewUrlShortenerArrayDb(size int) *UrlShortenerArrayDb {
	urls := make([]*core.ShortenedUrl, size)
	return &UrlShortenerArrayDb{urls: urls}
}

func (d *UrlShortenerArrayDb) Register(shortened *core.ShortenedUrl) error {
	if d.count >= len(d.urls) {
		return errors.New("reached limit of url registrations")
	}
	shortened_ := *shortened // copy
	d.urls[d.count] = &shortened_
	d.count++
	return nil
}

func (d *UrlShortenerArrayDb) Search(id string) (*core.ShortenedUrl, error) {
	for i := 0; i < d.count; i++ {
		if d.urls[i].Id() == id {
			shortened := *(d.urls[i]) // copy
			return &shortened, nil
		}
	}
	return nil, nil
}

func (d *UrlShortenerArrayDb) SearchByUrl(original string) (*core.ShortenedUrl, error) {
	for i := 0; i < d.count; i++ {
		if d.urls[i].Original() == original {
			shortened := *(d.urls[i]) // copy
			return &shortened, nil
		}
	}
	return nil, nil
}
