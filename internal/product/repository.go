package product

import (
	"errors"

	"github.com/google/uuid"
)

type ProductRepo struct {
	store map[string]Product
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{store: make(map[string]Product)}
}

func (r *ProductRepo) List() []Product {
	list := []Product{}
	for _, p := range r.store {
		list = append(list, p)
	}
	return list
}

func (r *ProductRepo) Get(id string) (Product, error) {
	p, ok := r.store[id]
	if !ok {
		return Product{}, errors.New("producto no encontrado")
	}
	return p, nil
}

func (r *ProductRepo) Create(p Product) Product {
	p.ID = uuid.New().String()
	r.store[p.ID] = p
	return p
}

func (r *ProductRepo) Update(id string, updated Product) (Product, error) {
	p, ok := r.store[id]
	if !ok {
		return Product{}, errors.New("producto no encontrado")
	}
	p.Name = updated.Name
	p.Description = updated.Description
	p.Price = updated.Price
	p.Stock = updated.Stock
	r.store[id] = p
	return p, nil
}

func (r *ProductRepo) Delete(id string) error {
	if _, ok := r.store[id]; !ok {
		return errors.New("producto no encontrado")
	}
	delete(r.store, id)
	return nil
}

func (r *ProductRepo) ReduceStock(id string, qty int) error {
	p, ok := r.store[id]
	if !ok {
		return errors.New("producto no encontrado")
	}
	if p.Stock < qty {
		return errors.New("stock insuficiente")
	}
	p.Stock -= qty
	r.store[id] = p
	return nil
}
