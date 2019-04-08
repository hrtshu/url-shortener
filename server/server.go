package server

import (
	"github.com/hrtshu/url-shortener/core"
	"net/http"
)

const (
	ENTRY_POINT_CREATE = "/v1/create"
)

type UrlShortenerServer struct {
	baseHostname string
	apiHostname  string
	scheme       string
	shortener    *core.UrlShortener
}

func NewUrlShortenerServer(baseHostname string, apiHostname string, scheme string, shortener *core.UrlShortener) *UrlShortenerServer {
	return &UrlShortenerServer{baseHostname: baseHostname, apiHostname: apiHostname, scheme: scheme, shortener: shortener}
}

func (s *UrlShortenerServer) BaseHostname() string {
	return s.baseHostname
}

func (s *UrlShortenerServer) ApiHostname() string {
	return s.apiHostname
}

func (s *UrlShortenerServer) Shortener() *core.UrlShortener {
	return s.shortener
}

func (s *UrlShortenerServer) Start(addr string) error {
	http.Handle("/", rootHandler(s))
	http.Handle(ENTRY_POINT_CREATE, createHandler(s))
	return http.ListenAndServe(addr, nil)
}

func rootHandler(s *UrlShortenerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := extractIdFromHostname(r.Host, s.baseHostname, s.shortener.IdSize())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		shortened, err := s.shortener.Resolve(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if shortened == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		redirect(w, shortened.Original())
	})
}

func createHandler(s *UrlShortenerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Host != s.apiHostname {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		query := r.URL.Query()
		urls, ok := query["url"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest) // TODO
			respondJson(w, newErrorResponse("missing parameter: url"))
			return
		}
		url := urls[0] // use the first one if multiple urls are given
		shortened, err := s.shortener.Create(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)                                     // TODO
			respondJson(w, newErrorResponse("invalid url or internal server error")) // TODO
			return
		}
		fullShortenedUrl := shortened.GetFull(s.scheme, s.baseHostname)
		respondJson(w, newSuccessResponse(fullShortenedUrl, shortened.Original()))
	})
}
