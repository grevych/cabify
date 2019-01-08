package memory

import (
	"github.com/grevych/cabify/pkg/entities"
)

type store struct {
	Items map[string]entities.Entity
}

func NewStore() *store {
	return &store{
		Items: map[string]entities.Entity{},
	}
}

func (s *store) FindEntityById(itemId string) (entities.Entity, error) {
	return nil, nil
}

func (s *store) Save(item entities.Entity) (string, error) {
	if item.GetId() == "" {
		_ = item.SetId("UUID")
	}
	return "", nil
}

func (s *store) Update(item entities.Entity) error {
	return nil
}

func (s *store) Delete(itemId string) error {
	return nil
}
