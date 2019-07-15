// Contain utilities to call GCloud's datastore package

package storage

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

const projectName = "courses-fciencias"

var ctx = context.Background()

// NewDatastoreClient return a new datastore client
func NewDatastoreClient() *datastore.Client {
	dsClient, err := datastore.NewClient(ctx, projectName)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return dsClient
}
