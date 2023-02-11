package handlers

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/models"
	"github.com/tyange/pian-fiber/store"
	"google.golang.org/api/option"
	"os"
)

func AuthMiddleware() fiber.Handler {
	return LoginMiddleware
}

func LoginMiddleware(c *fiber.Ctx) error {
	sess, err := store.Store.Get(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "세션을 불러오지 못했습니다."})

	}

	text := sess.Get("pian_login")

	fmt.Println(text)

	if sess.Get("pian_login") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}

func VerifyingGoogleAuthProviderForFirebase(c *fiber.Ctx) error {
	if godotenv.Load() != nil {
		return c.Status(400).JSON(fiber.Map{"error": "environment를 불러오지 못했습니다."})
	}

	session, err := store.Store.Get(c)

	if err != nil {
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

	userEmail := userData.Claims["email"].(string)

	user := models.User{}
	result := database.DBConn.First(&user, "email = ?", userEmail)

	if result.Error != nil {
		user.Iss = userData.Firebase.SignInProvider
		user.Email = userEmail
		user.UID = userData.UID

		database.DBConn.Save(&user)
	}

	session.Set("pian_login", "true")

	if session.Save() != nil {
		return c.Status(400).JSON(fiber.Map{"error": "세션 정보 저장에 실패했습니다."})
	}

	return c.Status(200).JSON(userData)
}
