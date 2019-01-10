package entities

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewEntity(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entityId       string
		expectedEntity *entity
	}{
		"TestNewEntity": {
			entityId:       entityId,
			expectedEntity: &entity{entityId},
		},

		"TestNewEntityWithEmptyId": {
			entityId:       "",
			expectedEntity: &entity{},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			entity := NewEntity(testCase.entityId)

			if !reflect.DeepEqual(testCase.expectedEntity, entity) {
				t.Errorf(
					"Expected entity %+v, got %+v",
					testCase.expectedEntity,
					entity,
				)
			}
		})
	}
}

func TestGetId(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entity           *entity
		expectedEntityId string
	}{
		"TestGetId": {
			entity:           &entity{entityId},
			expectedEntityId: entityId,
		},

		"TestGetEmptyId": {
			entity:           &entity{},
			expectedEntityId: "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			entityId := testCase.entity.GetId()

			if !reflect.DeepEqual(testCase.expectedEntityId, entityId) {
				t.Errorf(
					"Expected entity id %+v, got %+v",
					testCase.expectedEntityId,
					entityId,
				)
			}
		})
	}
}

func TestSetId(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entity        *entity
		entityId      string
		expectedError error
	}{
		"TestSetId": {
			entity:        &entity{""},
			entityId:      entityId,
			expectedError: nil,
		},

		"TestSetIdTwice": {
			entity:        &entity{entityId},
			entityId:      "new-entity-id",
			expectedError: errors.New("Entity id is immutable!"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.entity.SetId(testCase.entityId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}
		})
	}
}
