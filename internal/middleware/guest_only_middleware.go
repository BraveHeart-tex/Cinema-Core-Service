package middleware

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/cookies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	services "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	"github.com/gin-gonic/gin"
)

func GuestOnlyMiddleware(sessionService *services.SessionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(cookies.SessionCookieName)
		if err == nil && token != "" {
			session, err := sessionService.ValidateSessionToken(ctx, token)
			if err == nil && session != nil {
				responses.Error(ctx, http.StatusForbidden, "forbidden")
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
