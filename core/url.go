package core

import (
	"fmt"
)

type ShortenedUrl struct {
	original string
	id       string
}

func newShortenedUrl(original string, id string) *ShortenedUrl {
	return &ShortenedUrl{original: original, id: id}
}

func (s *ShortenedUrl) Original() string {
	return s.original
}

func (s *ShortenedUrl) Id() string {
	return s.id
}

func (s *ShortenedUrl) GetFull(scheme string, baseHostname string) string {
	return fmt.Sprintf("%s://%s.%s", scheme, s.id, baseHostname)
}
