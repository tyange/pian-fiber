package store

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

var Store *session.Store

func SetSession() {
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 5,
	})
}
