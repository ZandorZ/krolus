package feed

import (
	"krolus/models"
	"net/http"
)

// Parser ...
type Parser interface {
	Parse(*models.SubscriptionModel) (models.ItemCollection, error)
}

// Requester ..
type Requester interface {
	Request(url string) (*http.Response, error)
}
