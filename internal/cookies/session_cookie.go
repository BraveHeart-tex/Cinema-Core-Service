package cookies

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName = "session"
	sessionCookieTTL  = 24 * time.Hour
)

func SetSessionCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(
		SessionCookieName,
		token,
		int(sessionCookieTTL.Seconds()),
		"/",
		"",
		os.Getenv("APP_ENV") == "production",
		true,
	)
}
