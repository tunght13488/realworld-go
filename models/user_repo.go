package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

var users = []User{
	{
		Email:    "tunght13488@gmail.com",
		Username: "tunght13488",
		Bio:      "Bio",
		Image:    "",
		Token:    "",
		Hash:     fmt.Sprintf("%s", sha256.Sum256([]byte("xxxxxxxx"))),
	},
}

func FindByEmail(email string) (User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New(fmt.Sprintf("User not found: email=%s", email))
}
