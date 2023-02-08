package models

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
	Hash     string `json:"-"`
}

func (u User) WithToken(t string) User {
	u.Token = t
	return u
}

func (u User) MatchHash(h string) bool {
	return fmt.Sprintf("%s", sha256.Sum256([]byte(h))) == u.Hash
}
