package middlewares

import (
	"fmt"
	"onshops/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthGuaard(c *fiber.Ctx) error {
	token, err := GetTokenFromHeader(c)
	if err != nil {
		return pkg.ErrRosponse(c, 401, "unauthorized", err.Error())
	}
	id, err := pkg.JwtVerify(token)
	if err != nil {
		return pkg.ErrRosponse(c, 401, "unauthorized", err.Error())
	}
	c.Locals("customer_id", id)
	return c.Next()
}
func GetTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("token is empty")
	}
	token := strings.Split(header, " ")
	return token[1], nil
}
