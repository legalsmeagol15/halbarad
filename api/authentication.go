package main

import "golang.org/x/crypto/bcrypt"

type credentials struct {
	Username string
	Password string
}

func hashPswd(password string) (string, error) {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		return string(hashed), nil
	} else {
		return "", err
	}
}

func getClient(creds credentials) (*Client, error) {
	return nil, nil
}

func createAccount(credentials) error {
	return nil
}
