package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func JwtAuthMiddleware(jwtMaker token.JwtMakerInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHandler := ctx.Request.Header.Get(authorizationHeaderKey)

		if len(authHandler) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		fields := strings.Fields(authHandler)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		accessToken := fields[0]

		payload, err := jwtMaker.VerifyJwtToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if payload.Validate() != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)

		ctx.Next()
	}
}
