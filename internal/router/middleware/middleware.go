package middleware

import (
	context2 "context"
	"github.com/cend-org/duval/internal/token"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.FormValue("Authentication")

		if len(strings.TrimSpace(token)) == 0 || !strings.Contains(token, "Bearer") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication failed",
			})
		}

		/*

		 inject a special header X-Request-Id to response headers that could be used to track incoming
		 requests for monitoring/debugging purposes.
		 Value of request id header is usually formatted as UUID V4.

		*/

		context.Writer.Header().Set("x-request-id", uuid.NewV4().String())

		context.Next()
	}
}

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("Authorization")

		if len(strings.TrimSpace(tokenStr)) != 0 && !strings.Contains(tokenStr, "Bearer") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication failed",
			})
		}

		if len(strings.TrimSpace(tokenStr)) > 0 {
			var tok *token.Token
			tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
			tok = token.Parse(tokenStr)

			if tok != nil {
				if tok.ExpirationDate.Value.IsZero() || tok.ExpirationDate.Value.Before(time.Now()) {
					context.AbortWithStatusJSON(http.StatusLocked, gin.H{
						"error": "token is expired",
					})
				}
				ctx := context2.WithValue(context.Request.Context(), "token", *tok)
				context.Request = context.Request.WithContext(ctx)
			}
		}

		/*

		 inject a special header X-Request-Id to response headers that could be used to track incoming
		 requests for monitoring/debugging purposes.
		 Value of request id header is usually formatted as UUID V4.

		*/

		context.Writer.Header().Set("x-request-id", uuid.NewV4().String())

		context.Next()
	}
}
