package persisters

import (
	"fmt"

	"github.com/djdongjin/app-metadata-server/common"
)

type Persister interface {
	Persist(common.Metadata) error
	Retrieve(query string) []common.Metadata
	Get(title string) (common.Metadata, bool)
	Delete(title string) (common.Metadata, bool)
}

type InMemoryPersister struct {
	data map[string]common.Metadata
}

func (p InMemoryPersister) Persist(metadata common.Metadata) error {
	key := metadata.Title
	p.data[key] = metadata

	return nil
}

func (p InMemoryPersister) Retrieve(query string) []common.Metadata {
	return make([]common.Metadata, 0)
}

func (p InMemoryPersister) Get(title string) (common.Metadata, bool) {
	metadata, ok := p.data[title]

	return metadata, ok
}

func (p InMemoryPersister) Delete(title string) (common.Metadata, bool) {
	if metadata, ok := p.data[title]; ok {
		delete(p.data, title)
		return metadata, ok
	} else {
		return common.Metadata{}, false
	}
}

func New(tp string) (Persister, error) {
	switch tp {
	case "InMemoryPersister":
		return InMemoryPersister{}, nil
	default:
		return nil, fmt.Errorf("%s type doesn't exist", tp)
	}
}
