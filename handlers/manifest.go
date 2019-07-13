package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

var BoltPath string

var ErrManifestNotFound = func(host string) error {
	return fmt.Errorf("manifest not found for %s", host)
}

func GetManifest(params operations.GetManifestParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewGetManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	m := map[string]interface{}{}

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		if bucket == nil {
			return ErrManifestNotFound(params.Host)
		}

		latest := bucket.Get([]byte("latest"))
		return json.Unmarshal(latest, &m)
	}); err != nil {
		switch err.Error() {
		case ErrManifestNotFound(params.Host).Error():
			return operations.NewGetManifestNotFound().WithPayload(err.Error())
		default:
			return operations.NewGetManifestInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewGetManifestOK().WithPayload(m).WithCacheControl("no-cache")
}

func PutManifest(params operations.PutManifestParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	m := map[string]interface{}{}

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		if bucket == nil {
			return ErrManifestNotFound(params.Host)
		}

		latest := bucket.Get([]byte("latest"))
		return json.Unmarshal(latest, &m)
	}); err != nil {
		switch err.Error() {
		case ErrManifestNotFound(params.Host).Error():
			return operations.NewPutManifestNotFound().WithPayload(err.Error())
		default:
			return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
		}
	}

	for microApp, bundle := range params.Manifest.(map[string]interface{}) {
		m[microApp] = bundle.(string)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))

		manifest, err := json.Marshal(m)
		if err != nil {
			return err
		}

		return bucket.Put([]byte("latest"), manifest)
	}); err != nil {
		return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
	}

	return operations.NewPutManifestCreated().WithPayload(fmt.Sprintf("created manifest for host %s", params.Host))
}

func PostManifest(params operations.PostManifestParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPostManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(params.Host))
		if err != nil {
			return err
		}

		manifest, err := json.Marshal(params.Manifest)
		if err != nil {
			return err
		}

		if err = bucket.Put([]byte("latest"), manifest); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return operations.NewPostManifestInternalServerError().WithPayload(err.Error())
	}

	return operations.NewPostManifestCreated().WithPayload(fmt.Sprintf("created manifest for host %s", params.Host))
}
