package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

func GetInfo(params operations.GetInfoParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewGetInfoInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	m := &Microfest{}

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(params.Host))
		latest := bucket.Get([]byte(params.Key))
		return json.Unmarshal(latest, m)
	}); err != nil {
		return operations.NewGetInfoNotFound().WithPayload(fmt.Sprintf("could not find a manifest with key %s for %s", params.Key, params.Host))
	}

	return operations.NewGetInfoOK().WithPayload(m)
}
