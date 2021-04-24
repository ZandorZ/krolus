package feed

import (
	"krolus/models"
	"net/http"
)

type Requester interface {
	Request(url string) (*http.Response, error)
}

type Checker interface {
	Requester
	Check(*models.SubscriptionModel) (models.ItemCollection, error)
}
