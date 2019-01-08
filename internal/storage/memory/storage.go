package memory

type Storage struct {
	basketStore  *BasketStore
	productStore *ProductStore
}

func NewStorage(bs *BasketStore, ps *ProductStore) *Storage {
	return &Storage{
		basketStore:  bs,
		productStore: ps,
	}
}

func (s *Storage) GetBasketStore() *BasketStore {
	return s.basketStore
}

func (s *Storage) GetProductStore() *ProductStore {
	return s.productStore
}
