package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

var ErrConfigurationNotFound = func(host string) error {
	return fmt.Errorf("configuration not found for %s", host)
}

func GetConfiguration(params operations.GetConfigurationParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewGetConfigurationInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	c := &map[string]interface{}{}

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		if bucket == nil {
			return ErrConfigurationNotFound(params.Host)
		}

		latest := bucket.Get([]byte("configuration"))
		return json.Unmarshal(latest, c)
	}); err != nil {
		switch err.Error() {
		case ErrConfigurationNotFound(params.Host).Error():
			return operations.NewGetConfigurationNotFound().WithPayload(err.Error())
		default:
			return operations.NewGetConfigurationInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewGetConfigurationOK().WithPayload(c).WithCacheControl("no-cache")
}

func PostConfiguration(params operations.PostConfigurationParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPostConfigurationInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(params.Host))
		if err != nil {
			return err
		}

		configuration, err := json.Marshal(params.Configuration)
		if err != nil {
			return err
		}

		if err = bucket.Put([]byte("configuration"), configuration); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return operations.NewPostConfigurationInternalServerError().WithPayload(err.Error())
	}

	return operations.NewPostConfigurationCreated().WithPayload(fmt.Sprintf("created configuration for host %s", params.Host))
}
