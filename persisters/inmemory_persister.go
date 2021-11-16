package persisters

import (
	"fmt"
	"sync"

	"github.com/djdongjin/app-metadata-server/common"
)

// InMemoryPersister holds a map to store metadata where the key is title
// and value is corresponding Metadata.
type InMemoryPersister struct {
	mu   sync.Mutex
	data map[string]common.Metadata
}

func (p *InMemoryPersister) Persist(metadata common.Metadata) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key := metadata.Title
	p.data[key] = metadata

	return nil
}

func (p *InMemoryPersister) Retrieve(subqueries []common.SubQuery) ([]common.Metadata, error) {
	res := make([]common.Metadata, 0)

	for _, v := range p.data {
		if common.ExecuteQuery(v, subqueries) {
			res = append(res, v)
		}
	}

	return res, nil
}

func (p *InMemoryPersister) Get(title string) (common.Metadata, bool) {
	metadata, ok := p.data[title]

	return metadata, ok
}

func (p *InMemoryPersister) Delete(title string) (common.Metadata, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if metadata, ok := p.data[title]; ok {
		fmt.Println(len(p.data))
		delete(p.data, title)
		fmt.Println(len(p.data))
		return metadata, ok
	} else {
		return common.Metadata{}, false
	}
}

func NewInMemoryPersister() (*InMemoryPersister, error) {
	return &InMemoryPersister{data: make(map[string]common.Metadata)}, nil
}
