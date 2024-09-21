package router

import (
	"errors"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"net/http"
)

type TestUserVerifier struct {
}

// Validates the username and password
func (*TestUserVerifier) ValidateUser(email, password, scope string, r *http.Request) error {
	pdb := database.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.Email.Equals(email)).Exec(pdb.Context)
	if err != nil {
		return errors.New("wrong user")
	}
	if user.Password == password {
		return nil
	}

	return errors.New("wrong user")
}

// Adds claims to the token
func (*TestUserVerifier) AddClaims(email, tokenID, tokenType, scope string) (map[string]string, error) {
	pdb := database.DBclient
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
func (*TestUserVerifier) AddProperties(email, tokenId, tokenType string, scope string) (map[string]string, error) {
	props := make(map[string]string)
	// props["customer_name"] = "Gopher"
	return props, nil
}



// Unused methods ----------------------------------------------------------------------------------------
func (*TestUserVerifier) ValidateClient(clientID, clientSecret, scope string, r *http.Request) error {
	return errors.New("wrong client")
}

func (*TestUserVerifier) ValidateTokenId(email, tokenId, refreshTokenID, tokenType string) error {
	return nil
}

func (*TestUserVerifier) StoreTokenId(email, tokenId, refreshTokenID, tokenType string) error {
	return nil
}
