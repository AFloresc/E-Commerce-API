package product

import (
	"errors"

	"github.com/google/uuid"
)

var products = map[string]Product{}

func ListProducts() []Product {
	list := []Product{}
	for _, p := range products {
		list = append(list, p)
	}
	return list
}

func GetProduct(id string) (Product, error) {
	p, ok := products[id]
	if !ok {
		return Product{}, errors.New("producto no encontrado")
	}
	return p, nil
}

func CreateProduct(p Product) Product {
	p.ID = uuid.New().String()
	products[p.ID] = p
	return p
}

func UpdateProduct(id string, updated Product) (Product, error) {
	p, ok := products[id]
	if !ok {
		return Product{}, errors.New("producto no encontrado")
	}
	p.Name = updated.Name
	p.Description = updated.Description
	p.Price = updated.Price
	p.Stock = updated.Stock
	products[id] = p
	return p, nil
}

func DeleteProduct(id string) error {
	if _, ok := products[id]; !ok {
		return errors.New("producto no encontrado")
	}
	delete(products, id)
	return nil
}
