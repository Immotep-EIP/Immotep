package router

import (
	"errors"
	"net/http"

	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/utils"
)

type TestUserVerifier struct{}

// Validates the username and password
func (*TestUserVerifier) ValidateUser(email, password, _scope string, _r *http.Request) error {
	pdb := services.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.Email.Equals(email)).Exec(pdb.Context)
	if err != nil {
		return errors.New("wrong user")
	}
	if utils.CheckPasswordHash(password, user.Password) {
		return nil
	}

	return errors.New("wrong user")
}

// Adds claims to the token
func (*TestUserVerifier) AddClaims(email, _tokenId, _tokenType, _scope string) (map[string]string, error) {
	pdb := services.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.Email.Equals(email)).Exec(pdb.Context)
	if err != nil {
		return nil, errors.New("wrong user")
	}

	claims := make(map[string]string)
	claims["role"] = string(user.Role)
	claims["id"] = user.ID
	// claims["customer_data"] = `{"order_date":"2016-12-14","order_id":"9999"}`

	return claims, nil
}

// Adds properties to the token
func (*TestUserVerifier) AddProperties(_email, _tokenId, _tokenType string, _scope string) (map[string]string, error) {
	props := make(map[string]string)
	// props["customer_name"] = "Gopher"
	return props, nil
}

// Unused methods ----------------------------------------------------------------------------------------
func (*TestUserVerifier) ValidateClient(_clientId, _clientSecret, _scope string, _r *http.Request) error {
	return errors.New("wrong client")
}

func (*TestUserVerifier) ValidateTokenId(_email, _tokenId, _refreshTokenId, _tokenType string) error {
	return nil
}

func (*TestUserVerifier) StoreTokenId(_email, _tokenId, _refreshTokenId, _tokenType string) error {
	return nil
}
