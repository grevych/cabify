package storage

/*
func TestCreate(t *testing.T) {
	tests := map[string]struct {
		storageType   string
		expectedError error
	}{
		"TestCreate": {
			storageType:   "memory",
			expectedError: nil,
		},

		"TestCreateUndefinedStorage": {
			storageType:   "undefined",
			expectedError: errors.New("Storage type undefined not found"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			storage, err := Create(testCase.storageType)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}

			if _, ok := storage.(Storage); !ok {
				t.Errorf("Expected storage %s to implement Storage interface", testCase.storageType)
			}
		})
	}
}
*/
