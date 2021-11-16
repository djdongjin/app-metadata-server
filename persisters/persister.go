package persisters

import (
	"fmt"

	"github.com/djdongjin/app-metadata-server/common"
)

// Persister is the interface for storing/retrieving Metadata.
type Persister interface {
	Persist(common.Metadata) error
	Retrieve([]common.SubQuery) ([]common.Metadata, error)
	Get(title string) (common.Metadata, bool)
	Delete(title string) (common.Metadata, bool)
}

// New is the factory function that returns a Persister implementation.
func New(tp string) (Persister, error) {
	switch tp {
	case "InMemoryPersister":
		return NewInMemoryPersister()
	default:
		return nil, fmt.Errorf("%s type doesn't exist", tp)
	}
}
