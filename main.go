package main

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tunght13488/realworld-go/handlers"
	"github.com/tunght13488/realworld-go/models"
	"go.uber.org/zap"
	"net/http"
)

var user = models.User{
	Email:    "tunght13488@gmail.com",
	Username: "tunght13488",
	Bio:      "Bio",
	Image:    "",
	Token:    "",
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*handlers.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	logger, _ := zap.NewProduction()
	//defer logger.Sync()
	sugar := logger.Sugar()
	defer sugar.Sync()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			//logger.Info("request",
			//	zap.String("URI", v.URI),
			//	zap.Int("status", v.Status),
			//)
			sugar.Infow("request",
				"URI", v.URI,
				"status", v.Status,
			)

			return nil
		},
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Login route
	e.POST("/users/login", handlers.Login)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("", restricted)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
