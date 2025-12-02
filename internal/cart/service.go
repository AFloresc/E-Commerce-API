package cart

import (
	"errors"
)

var carts = map[string]Cart{}

func GetCart(userID string) Cart {
	c, ok := carts[userID]
	if !ok {
		c = Cart{UserID: userID, Items: []CartItem{}}
		carts[userID] = c
	}
	return c
}

func AddToCart(userID, productID string, qty int) Cart {
	c := GetCart(userID)
	// Buscar si ya existe el producto
	found := false
	for i, item := range c.Items {
		if item.ProductID == productID {
			c.Items[i].Quantity += qty
			found = true
			break
		}
	}
	if !found {
		c.Items = append(c.Items, CartItem{ProductID: productID, Quantity: qty})
	}
	carts[userID] = c
	return c
}

func RemoveFromCart(userID, productID string) (Cart, error) {
	c := GetCart(userID)
	newItems := []CartItem{}
	removed := false
	for _, item := range c.Items {
		if item.ProductID == productID {
			removed = true
			continue
		}
		newItems = append(newItems, item)
	}
	if !removed {
		return c, errors.New("producto no encontrado en el carrito")
	}
	c.Items = newItems
	carts[userID] = c
	return c, nil
}
