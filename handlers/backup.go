package handlers

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

func PostBackup(params operations.PostBackupParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return operations.NewPostBackupInternalServerError().WithPayload(err.Error())
	}
	defer client.Close()

	bucket := client.Bucket(params.Bucket)
	name := fmt.Sprintf("microfest-backup-%s.db", time.Now().Format("2006-01-02-150405-MST"))
	objectWriter := bucket.Object(name).NewWriter(ctx)

	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewGetInfoInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.View(func(tx *bolt.Tx) error {
		if _, err := tx.WriteTo(objectWriter); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return operations.NewPostBackupInternalServerError().WithPayload(err.Error())
	}

	if err := objectWriter.Close(); err != nil {
		return operations.NewPostBackupInternalServerError()
	}

	return operations.NewPostBackupCreated().WithPayload(fmt.Sprintf("gs://%s/%s", params.Bucket, name))
}
