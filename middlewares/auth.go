package middlewares

import (
	"github.com/IKHINtech/go-fiber-rest-boilerplate/config"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	jwtware "github.com/gofiber/contrib/jwt"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(config.AppConfig.SECRET)},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	h := utils.ResponseHandler{}
	if err.Error() == "Missing or malformed JWT" {
		return h.BadRequest(c, []string{"Missing or malformed JWT"})
	}
	return h.Unauthorized(c, []string{"Invalid or expired JWT"})
}

func jwtSuccess(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}
	// Ambil token yang sudah didekode
	token := c.Locals("user").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return h.Unauthorized(c, []string{"Invalid JWT claims"})
	}

	// Ambil informasi user dari claims dan set ke context
	userID := claims["user_id"].(string)
	username := claims["username"].(string)

	// Menyimpan informasi user ke dalam context Fiber
	c.Locals("user_id", userID)
	c.Locals("username", username)
	return c.Next()
}
