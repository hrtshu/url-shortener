package example

import (
	"github.com/hrtshu/url-shortener/core"
	"github.com/hrtshu/url-shortener/server"
	"log"
)

func RunLocalhost() {
	const maxUrls = 100
	const idSize = 6
	const baseHostname = "localhost:8080"
	const apiHostname = "api.localhost:8080"
	const scheme = "http"
	const addr = ":8080"

	db := core.NewUrlShortenerArrayDb(maxUrls)
	shortener := core.NewUrlShortener(db, idSize)
	shortenerServer := server.NewUrlShortenerServer(baseHostname, apiHostname, scheme, shortener)
	log.Fatal(shortenerServer.Start(addr))
}
