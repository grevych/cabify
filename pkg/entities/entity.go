package entities

import "errors"

type entity struct {
	Id string
}

type Entity interface {
	GetId() string
	SetId(string) error
}

var _ Entity = &entity{}

func NewEntity(id string) Entity {
	return &entity{id}
}

func (e *entity) GetId() string {
	return e.Id
}

func (e *entity) SetId(id string) error {
	if e.Id == "" {
		e.Id = id
		return nil
	}

	return errors.New("Entity id is immutable!")
}
