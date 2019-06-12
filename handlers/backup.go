package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func PostBackup(params operations.PostBackupParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return operations.NewPostBackupInternalServerError().WithPayload(err.Error())
	}
	defer client.Close()

	bucket := client.Bucket(params.Bucket)
	name := fmt.Sprintf("microfest-backup-%s.db", time.Now().Format("2006-01-02-15:04:05-MST"))
	objectWriter := bucket.Object(name).NewWriter(ctx)

	file, err := os.Open(BoltPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(objectWriter, file); err != nil {
		return operations.NewPostBackupInternalServerError()
	}

	if err := objectWriter.Close(); err != nil {
		return operations.NewPostBackupInternalServerError()
	}

	return operations.NewPostBackupCreated().WithPayload(fmt.Sprintf("gs://%s/%s", params.Bucket, name))
}
