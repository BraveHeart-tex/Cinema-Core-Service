package middleware

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/cookies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

const SessionContextKey = "session"

func SessionAuthMiddleware(sessionService *services.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookies.SessionCookieName)
		if err != nil || token == "" {
			responses.Error(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		session, err := sessionService.ValidateSessionToken(token)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		if session == nil {
			responses.Error(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Set(SessionContextKey, session)
		c.Next()
	}
}
