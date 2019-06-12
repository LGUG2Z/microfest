package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

var BoltPath string

type Microfest struct {
	Release  *string           `json:"release"`
	Manifest map[string]string `json:"manifest"`
	Updated  []string          `json:"updated"`
}

func GetManifest(params operations.GetManifestParams) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewGetManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	m := &Microfest{}

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		latest := bucket.Get([]byte("latest"))
		return json.Unmarshal(latest, m)
	}); err != nil {
		return operations.NewGetManifestInternalServerError().WithPayload(err.Error())
	}

	return operations.NewGetManifestOK().WithPayload(m.Manifest)
}

func PutManifest(params operations.PutManifestParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	m := &Microfest{}

	key := time.Now().Format("20060102150405MST")

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		latest := bucket.Get([]byte("latest"))
		return json.Unmarshal(latest, m)
	}); err != nil {
		return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
	}

	m.Release = params.Microfest.Release
	m.Updated = params.Microfest.Updated

	for microApp, bundle := range params.Microfest.Manifest.(map[string]interface{}) {
		m.Manifest[microApp] = bundle.(string)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))

		manifest, err := json.Marshal(m)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(key), manifest); err != nil {
			return err
		}

		return bucket.Put([]byte("latest"), manifest)
	}); err != nil {
		return operations.NewPutManifestInternalServerError().WithPayload(err.Error())
	}

	return operations.NewPutManifestCreated().WithPayload(fmt.Sprintf("created manifest %s with key %s", *params.Microfest.Release, key))
}

func PostManifest(params operations.PostManifestParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPostManifestInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	key := time.Now().Format("20060102150405MST")

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(params.Host))
		if err != nil {
			return err
		}

		manifest, err := json.Marshal(params.Microfest)
		if err != nil {
			return err
		}

		if err = bucket.Put([]byte(key), manifest); err != nil {
			return err
		}

		if err = bucket.Put([]byte("latest"), manifest); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return operations.NewPostManifestInternalServerError().WithPayload(err.Error())
	}

	return operations.NewPostManifestCreated().WithPayload(fmt.Sprintf("created manifest %s with key %s", *params.Microfest.Release, key))
}
