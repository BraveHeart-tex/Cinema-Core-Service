package cookies

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	sessionCookieName = "session"
	sessionCookieTTL  = 24 * time.Hour
)

func SetSessionCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(
		sessionCookieName,
		token,
		int(sessionCookieTTL.Seconds()),
		"/",
		"",
		os.Getenv("APP_ENV") == "production",
		true,
	)
}
