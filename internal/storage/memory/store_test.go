package memory

import (
	"errors"
	"reflect"
	"testing"

	"github.com/grevych/cabify/pkg/entities"
)

func createStore(items ...entities.Entity) *store {
	store := NewStore()

	for _, item := range items {
		store.items[item.GetId()] = item
	}

	return store
}

func createBasketStore(baskets ...entities.Entity) *BasketStore {
	basketStore := NewBasketStore()

	for _, basket := range baskets {
		basketStore.items[basket.GetId()] = basket
	}

	return basketStore
}

func createProductStore(products ...entities.Entity) *ProductStore {
	productStore := NewProductStore()

	for _, product := range products {
		productStore.items[product.GetId()] = product
	}

	return productStore
}

func createBasket(basketId string, products ...*entities.Product) *entities.Basket {
	basket, _ := entities.NewBasket(basketId, products)

	return basket
}

func createProduct(productId, productCode, productName string, productPrice int64) *entities.Product {
	product, _ := entities.NewProduct(
		productId, productCode, productName, productPrice,
	)

	return product
}

func TestNewStore(t *testing.T) {
	tests := map[string]struct {
		expectedStore *store
	}{
		"TestStore": {
			expectedStore: &store{
				items: map[string]entities.Entity{},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			store := NewStore()

			if !reflect.DeepEqual(testCase.expectedStore, store) {
				t.Errorf(
					"Expected store %+v, got %+v",
					testCase.expectedStore,
					store,
				)
			}
		})
	}
}

func TestAll(t *testing.T) {
	entityA := entities.NewEntity("entity-A-id")
	entityB := entities.NewEntity("entity-B-id")

	tests := map[string]struct {
		store            *store
		expectedEntities []entities.Entity
	}{
		"TestAll": {
			store:            createStore(entityA, entityB),
			expectedEntities: []entities.Entity{entityA, entityB},
		},

		"TestAllWithEmptyStore": {
			store:            createStore(),
			expectedEntities: []entities.Entity{},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			entities := testCase.store.All()

			if !reflect.DeepEqual(testCase.expectedEntities, entities) {
				t.Errorf(
					"Expected entities %+v, got %+v",
					testCase.expectedEntities,
					entities,
				)
			}
		})
	}
}

func TestFindById(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entityId       string
		store          *store
		expectedEntity entities.Entity
		expectedError  error
	}{
		"TestFindById": {
			entityId:       entityId,
			store:          createStore(entities.NewEntity(entityId)),
			expectedEntity: entities.NewEntity(entityId),
			expectedError:  nil,
		},

		"TestFindByIdNonExistentEntity": {
			entityId:       entityId,
			store:          NewStore(),
			expectedEntity: nil,
			expectedError:  errors.New("Entity entity-id not found"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			entity, err := testCase.store.FindById(testCase.entityId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}

			if !reflect.DeepEqual(testCase.expectedEntity, entity) {
				t.Errorf("Expected entity %+v, got %+v", testCase.expectedEntity, entity)
			}
		})
	}
}

func TestSave(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		store           *store
		entity          entities.Entity
		expectingEntity bool
		expectedError   error
	}{
		"TestSave": {
			entity:          entities.NewEntity(""),
			store:           NewStore(),
			expectingEntity: true,
			expectedError:   nil,
		},

		"TestSaveExistentEntity": {
			entity:          entities.NewEntity(entityId),
			store:           createStore(entities.NewEntity(entityId)),
			expectingEntity: false,
			expectedError:   errors.New("Entity entity-id exists"),
		},

		"TestSaveNonValidEntity": {
			entity:          entities.NewEntity(entityId),
			store:           NewStore(),
			expectingEntity: false,
			expectedError:   errors.New("Entity entity-id not valid"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			storedEntityId, err := testCase.store.Save(testCase.entity)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if testCase.expectingEntity && storedEntityId == nil {
				t.Errorf("Expected entity id, got nil")
			}

			if testCase.expectingEntity {
				if _, ok := testCase.store.items[*storedEntityId]; !ok {
					t.Errorf("Expected entity in store, got entity not in store")
				}
			}

			if !testCase.expectingEntity && storedEntityId != nil {
				t.Errorf("Expected nil, got entity id")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entity        entities.Entity
		store         *store
		expectedError error
	}{
		"TestUpdate": {
			entity:        entities.NewEntity(entityId),
			store:         createStore(entities.NewEntity(entityId)),
			expectedError: nil,
		},

		"TestUpdateNewEntity": {
			entity:        entities.NewEntity(entityId),
			store:         NewStore(),
			expectedError: errors.New("Entity entity-id not found"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.store.Update(testCase.entity)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			currentEntityId := testCase.entity.GetId()

			if err != nil {
				if _, ok := testCase.store.items[currentEntityId]; ok {
					t.Errorf("Expected entity not in store, got in store")
				}
			} else {
				storedEntity, _ := testCase.store.items[currentEntityId]
				if testCase.entity != storedEntity {
					t.Errorf(
						"Expected updated entity %+v, got %+v",
						testCase.entity, storedEntity,
					)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	entityId := "entity-id"

	tests := map[string]struct {
		entityId      string
		store         *store
		expectedError error
	}{
		"TestDelete": {
			entityId:      entityId,
			store:         createStore(entities.NewEntity(entityId)),
			expectedError: nil,
		},

		"TestDeleteNewEntity": {
			entityId:      entityId,
			store:         NewStore(),
			expectedError: errors.New("Entity entity-id not found"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.store.Delete(testCase.entityId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if _, ok := testCase.store.items[testCase.entityId]; ok {
				t.Errorf("Expected entity not in store, got in store")
			}
		})
	}
}
