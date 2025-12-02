package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     string
}

// Simulación de base de datos en memoria
var users = map[string]User{}

func CreateUser(email, password string) (User, error) {
	if _, exists := users[email]; exists {
		return User{}, errors.New("usuario ya existe")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{
		ID:       email,
		Email:    email,
		Password: string(hashed),
		Role:     "user",
	}
	users[email] = user
	return user, nil
}

func ValidateUser(email, password string) (User, error) {
	user, exists := users[email]
	if !exists {
		return User{}, errors.New("usuario no encontrado")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("contraseña incorrecta")
	}
	return user, nil
}
