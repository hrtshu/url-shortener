package db

import (
	"context"
	"github.com/hrtshu/url-shortener/core"
	"go.mercari.io/datastore"
	"google.golang.org/api/iterator"
)

const SHORTENED_URL_ENTITY = "ShortenedUrl"

type DatastoreDb struct {
	ctx    context.Context
	client *datastore.Client
}

func NewDatastoreDb(ctx context.Context, client *datastore.Client) *DatastoreDb {
	return &DatastoreDb{ctx: ctx, client: client}
}

func (d *DatastoreDb) Register(shortened *core.ShortenedUrl) error {
	key := (*d.client).IncompleteKey(SHORTENED_URL_ENTITY, nil)
	entity := newShortenedUrlEntity(shortened)
	key, err := (*d.client).Put(d.ctx, key, entity)
	if err != nil {
		return err // TODO
	}
	return nil
}

func (d *DatastoreDb) searchByQuery(query datastore.Query) (*core.ShortenedUrl, error) {
	it := (*d.client).Run(d.ctx, query.Limit(1))
	var entity *ShortenedUrlEntity
	entity = &ShortenedUrlEntity{}
	_, err := it.Next(entity)
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err // TODO
	}
	shortened := entity.toShortenedUrl()
	return shortened, nil
}

func (d *DatastoreDb) Search(id string) (*core.ShortenedUrl, error) {
	query := (*d.client).NewQuery(SHORTENED_URL_ENTITY).Filter("Id =", id)
	return d.searchByQuery(query)
}

func (d *DatastoreDb) SearchByUrl(original string) (*core.ShortenedUrl, error) {
	query := (*d.client).NewQuery(SHORTENED_URL_ENTITY).Filter("Original =", original)
	return d.searchByQuery(query)
}

type ShortenedUrlEntity struct {
	Original string
	Id       string
}

func newShortenedUrlEntity(shortened *core.ShortenedUrl) *ShortenedUrlEntity {
	return &ShortenedUrlEntity{Original: shortened.Original(), Id: shortened.Id()}
}

func (e *ShortenedUrlEntity) toShortenedUrl() *core.ShortenedUrl {
	return core.NewShortenedUrl(e.Original, e.Id)
}
