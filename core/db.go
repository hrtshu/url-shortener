package core

type UrlShortenerDb interface {
	Register(shortened *ShortenedUrl) error
	Search(id string) (*ShortenedUrl, error)
	SearchByUrl(original string) (*ShortenedUrl, error)
}
