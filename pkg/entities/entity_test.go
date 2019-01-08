package entities

import (
	"errors"
	"testing"
	"reflect"
)


func TestGetId(t *testing.T) {
	tests := map[string]struct{
		entity *entity
		expectedId string
	}{
		"TestGetId": {
			entity: &entity{"entity-id"},
			expectedId: "entity-id",
		},

		"TestGetEmptyId": {
			entity: &entity{},
			expectedId: "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			id := testCase.entity.GetId()

			if !reflect.DeepEqual(testCase.expectedId, id) {
				t.Errorf("Expected id %+v, got %+v", testCase.expectedId, id)
			}
		})
	}
}

func TestSetId(t *testing.T) {
	tests := map[string]struct{
		entity *entity
		entityId string
		expectedError error
	}{
		"TestSetId": {
			entity: &entity{""},
			entityId: "entity-id",
			expectedError: nil,
		},

		"TestSetIdTwice": {
			entity: &entity{"entity-id"},
			entityId: "entity-id",
			expectedError: errors.New("Entity id is immutable!"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.entity.SetId(testCase.entityId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}
		})
	}
}
