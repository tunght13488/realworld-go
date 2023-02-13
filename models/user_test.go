package models

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-faker/faker/v4"
	"reflect"
	"testing"
)

func TestUser_MatchHash(t *testing.T) {
	password := faker.Password()
	user := User{
		Email:    faker.Email(),
		Username: faker.Username(),
		Bio:      faker.Sentence(),
		Image:    faker.URL(),
		Token:    "",
		Hash:     fmt.Sprintf("%s", sha256.Sum256([]byte(password))),
	}
	expected := true
	t.Run("password", func(t *testing.T) {
		if got := user.MatchHash(password); got != expected {
			t.Errorf("MatchHash() = %v, want %v", got, expected)
		}
	})
}

func TestUser_WithToken(t *testing.T) {
	user := User{
		Email:    faker.Email(),
		Username: faker.Username(),
		Bio:      faker.Sentence(),
		Image:    faker.URL(),
		Token:    "",
		Hash:     "",
	}
	token := faker.UUIDHyphenated()
	expected := User{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
		Hash:     user.Hash,
	}
	t.Run("token", func(t *testing.T) {
		if got := user.WithToken(token); !reflect.DeepEqual(got, expected) {
			t.Errorf("WithToken() = %v, want %v", got, expected)
		}
	})
}
