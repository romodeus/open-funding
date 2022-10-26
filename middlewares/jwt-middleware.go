package middlewares

import (
	"fmt"
	"net/http"
	"open-funding/config"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token, err := GetToken(c)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(err.Error())
	}

	return nil
}

func CreateToken(userID uint) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT_SECRET))
}

func GetToken(c *fiber.Ctx) (jwt.Token, error) {
	cfg := config.GetConfig()

	authorizeHeader := c.GetRespHeader("Authorization")
	if !strings.Contains(authorizeHeader, "Bearer") {
		return jwt.Token{}, c.Status(http.StatusUnauthorized).JSON("failed authorization")
	}

	tokenString := strings.Replace(authorizeHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return cfg.JWT_SECRET, nil
	})

	if err != nil {
		return jwt.Token{}, c.Status(http.StatusUnauthorized).JSON(err.Error())
	}

	return *token, nil
}

func ExtractToken(c *fiber.Ctx) (int, error) {
	token, err := GetToken(c)
	if err != nil {
		return 0, err
	}

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		uid := claim["userID"].(float64)
		return int(uid), nil
	}

	return 0, fmt.Errorf("token invalid")
}
