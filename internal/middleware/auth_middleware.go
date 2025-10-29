package middleware

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/cookies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	sessionServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	userServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/user"
	"github.com/gin-gonic/gin"
)

const SessionContextKey = "session"

func SessionAuthMiddleware(sessionService *sessionServices.SessionService, userService *userServices.UserService) gin.HandlerFunc {
	// TODO: Investigate trx behavior here
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(cookies.SessionCookieName)
		if err != nil || token == "" {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		session, err := sessionService.ValidateSessionToken(ctx, token)
		if err != nil {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}
		if session == nil {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		user, err := userService.FindById(ctx, session.UserID)
		if err != nil {
			responses.Error(ctx, http.StatusInternalServerError, "internal error")
			ctx.Abort()
			return
		}
		if user == nil {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		ctx.Set(SessionContextKey, map[string]any{
			"session": session,
			"user":    user,
		})
		ctx.Next()
	}
}
