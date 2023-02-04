package handlers

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/models"
	"google.golang.org/api/idtoken"
)

func VerifyingCredential(c *fiber.Ctx) error {
	godotenv.Load()

	session, err := database.SessionStore.Get(c)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "GOOGLE CLIENT ID를 불러오지 못했습니다."})
	}

	credential := models.Credential{}

	if c.BodyParser(&credential) != nil {
		return c.Status(400).JSON(fiber.Map{"error": "토큰을 넘겨 받지 못했습니다."})
	}

	payload, validateErr := idtoken.Validate(context.Background(), credential.CredentialString, os.Getenv("GOOGLE_OAUTH_CLIENT_ID"))

	if validateErr != nil {
		return c.Status(400).JSON(fiber.Map{"error": "토큰을 인증할 수 없습니다."})
	}

	claims := payload.Claims
	user := models.User{}
	result := database.DBConn.First(&user, "email = ?", claims["email"])

	if result.Error != nil {
		user.Iss = claims["iss"].(string)
		user.Email = claims["email"].(string)
		user.Name = claims["name"].(string)

		database.DBConn.Save(&user)
	}

	session.Set("pian-login", true)
	session.Save()

	return c.Status(200).JSON(payload.Claims)
}
