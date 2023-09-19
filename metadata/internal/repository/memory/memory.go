package memory

import (
	"context"
	"sync"

	"github.com/bobhonores/somello/metadata/internal/repository"
	"github.com/bobhonores/somello/metadata/pkg/model"
)

// Repository defines a memory song metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory repository
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves song metadata by id
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.Unlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds song metadata per specific song id
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.RLock()
	defer r.Unlock()

	r.data[id] = metadata
	return nil
}
