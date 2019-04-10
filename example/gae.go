package example

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/hrtshu/url-shortener/core"
	"github.com/hrtshu/url-shortener/db"
	"github.com/hrtshu/url-shortener/server"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

func RunGAE() {
	const idSize = 6
	const addr = ":8080"

	baseHostname := os.Getenv("BASE_HOSTNAME")
	if baseHostname == "" {
		panic(errors.New("BASE_HOSTNAME is not set"))
	}
	apiHostname := os.Getenv("API_HOSTNAME")
	if apiHostname == "" {
		panic(errors.New("API_HOSTNAME is not set"))
	}
	scheme := os.Getenv("SCHEME")
	if scheme == "" {
		panic(errors.New("SCHEME is not set"))
	}
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	if projectID == "" {
		panic(errors.New("GCLOUD_DATASET_ID is not set"))
	}

	ctx := context.Background()
	client, err := clouddatastore.FromContext(
		ctx,
		datastore.WithProjectID(projectID),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	db := db.NewDatastoreDb(ctx, &client)
	shortener := core.NewUrlShortener(db, idSize)
	shortenerServer := server.NewUrlShortenerServer(baseHostname, apiHostname, scheme, shortener)
	log.Fatal(shortenerServer.Start(addr))
}
