package persistence

import "krolus/treex/models"

// Persister ...
type Persister interface {
	Save(models.Node) error
	Load(*models.Node) error
}
