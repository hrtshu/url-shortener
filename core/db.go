package core

type UrlShortenerDb interface {
	Register(shortened *ShortendUrl) error
	Search(id string) (*ShortendUrl, error)
	SearchByUrl(original string) (*ShortendUrl, error)
}
