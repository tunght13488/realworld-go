package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var loginJSON = `{"user":{"email":"tunght13488@gmail.com","password":"xxxxxxxx"}}`

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var payload UserPayload
		err := json.NewDecoder(rec.Body).Decode(&payload)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "tunght13488@gmail.com", payload.User.Email)
		assert.Equal(t, "tunght13488", payload.User.Username)
		assert.Equal(t, "Bio", payload.User.Bio)
		assert.Equal(t, "", payload.User.Image)
		assert.NotEmpty(t, payload.User.Token)
	}
}
