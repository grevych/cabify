package entities

import (
  "errors"
)

type entity struct {
  id string
}

type Entity interface {
	GetId() string
	SetId(string) error
}

var _ Entity = &entity{}

func(e *entity) GetId() string {
  return e.id
}

func(e *entity) SetId(id string) error {
  if e.id == "" {
	e.id = id
	return nil
  }

  return errors.New("Entity id is immutable!")
}
