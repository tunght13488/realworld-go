package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/tunght13488/realworld-go/models"
	"net/http"
	"time"
)

// JwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	User Credentials `json:"user"`
}

func Login(c echo.Context) error {
	var payload LoginPayload
	if err := c.Bind(&payload); err != nil {
		return echo.ErrBadRequest
	}

	user, err := models.FindByEmail(payload.User.Email)
	if err != nil || !user.MatchHash(payload.User.Password) {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		user.Email,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user.WithToken(t),
	})
}
