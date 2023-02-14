package handlers

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/models"
	"google.golang.org/api/option"
	"os"
)

func VerifyingGoogleAuthProviderForFirebase(c *fiber.Ctx) error {
	if godotenv.Load() != nil {
		return c.Status(400).JSON(fiber.Map{"error": "GOOGLE CLIENT ID를 불러오지 못했습니다."})
	}

	credential := models.Credential{}

	if c.BodyParser(&credential) != nil {
		return c.Status(400).JSON(fiber.Map{"error": "토큰을 넘겨 받지 못했습니다."})
	}

	opt := option.WithAPIKey(os.Getenv("FIREBASE_API_KEY"))

	config := &firebase.Config{ProjectID: "pian-firebase-auth"}

	app, err := firebase.NewApp(context.Background(), config, opt)

	client, err := app.Auth(context.Background())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "auth client 생성에 실패했습니다."})
	}

	userData, err := client.VerifyIDToken(context.Background(), credential.CredentialString)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "토큰을 인증할 수 없습니다."})
	}

	claims := userData.Claims

	return c.Status(200).JSON(claims)
}
