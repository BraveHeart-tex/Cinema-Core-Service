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
		int(sessionCookieTTL.Seconds()),      // maxAge in seconds (1 day)
		"/",                                  // path
		"",                                   // domain
		os.Getenv("APP_ENV") == "production", // secure
		true,                                 // httpOnly
	)
}
