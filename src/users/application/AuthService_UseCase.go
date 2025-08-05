package application

import (
	"AmethToledo/src/core/security"
	"AmethToledo/src/users/domain"
	"AmethToledo/src/users/domain/entities"
	"errors"
	"fmt"
	"strings"
)

type AuthService struct {
	clientRepo domain.IUser
}

func NewAuthService(clientRepo domain.IUser) *AuthService {
	return &AuthService{clientRepo: clientRepo}
}

func (as *AuthService) Login(email, password string) (map[string]interface{}, error) {
	email = strings.TrimSpace(email)
	fmt.Println("🔍 Buscando usuario con correo:", email)

	client, err := as.clientRepo.GetByCorreo(email)
	if err != nil {
		fmt.Println("❌ Error al obtener usuario:", err)
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}
	if client == nil {
		fmt.Println("⚠ Usuario no encontrado (nil)")
		return nil, errors.New("usuario no encontrado")
	}

	if !security.CheckPassword(client.PasswordHash, password) {
		fmt.Println("❌ Contraseña incorrecta")
		return nil, errors.New("contraseña incorrecta")
	}

	token, err := security.GenerateJWT(int(client.ID), client.Email)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":  token,
		"userId": client.ID,
		"name":   client.Nombres,
		"email":  client.Email,
	}, nil
}

func (as *AuthService) Register(client entities.User) (entities.User, error) {
	hashedPassword, err := security.HashPassword(client.PasswordHash)
	if err != nil {
		return entities.User{}, err
	}
	client.PasswordHash = hashedPassword

	savedUser, err := as.clientRepo.Save(client)
	if err != nil {
		return entities.User{}, err
	}

	return savedUser, nil
}
