package memory

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/grevych/cabify/pkg/entities"
)

type store struct {
	items map[string]entities.Entity
}

func NewStore() *store {
	return &store{
		items: map[string]entities.Entity{},
	}
}

func (s *store) FindById(itemId string) (entities.Entity, error) {
	item, ok := s.items[itemId]
	if !ok {
		return nil, fmt.Errorf("Entity %s not found", itemId)
	}

	return item, nil
}

func (s *store) Save(item entities.Entity) (*string, error) {
	itemId := item.GetId()

	_, ok := s.items[itemId]
	if ok {
		return nil, fmt.Errorf("Entity %s exists", itemId)
	}

	uuidv4, _ := uuid.NewV4()
	newItemId := uuidv4.String()

	if err := item.SetId(newItemId); err != nil {
		return nil, fmt.Errorf("Entity %s not valid", itemId)
	}

	s.items[newItemId] = item

	return &newItemId, nil
}

func (s *store) Update(item entities.Entity) error {
	itemId := item.GetId()

	_, ok := s.items[itemId]
	if !ok {
		return fmt.Errorf("Entity %s not found", itemId)
	}

	s.items[itemId] = item

	return nil
}

func (s *store) Delete(itemId string) error {
	_, ok := s.items[itemId]
	if !ok {
		return fmt.Errorf("Entity %s not found", itemId)
	}

	delete(s.items, itemId)

	return nil
}
