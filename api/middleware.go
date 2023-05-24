package api

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/julysNICK/stock_system/token"
)

const (
	authorizationHeaderToken     = "authorization"
	authorizationTypeToken       = "Bearer"
	AuthorizationPayloadKeyToken = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderToken)

		if len(authorizationHeader) == 0 {

			ctx.AbortWithStatusJSON(401, gin.H{"error": "authorization header is not provided"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		fmt.Printf("fields: %v", strings.ToLower(fields[0]))

		if len(fields) < 2 || strings.ToLower(fields[0]) != strings.ToLower(authorizationTypeToken) {

			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header"})
			return
		}

		authorizationTypeProvider := strings.ToLower(fields[0])

		if authorizationTypeProvider != strings.ToLower(authorizationTypeToken) {

			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization type"})
			return
		}

		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		ctx.Set(AuthorizationPayloadKeyToken, payload)
		ctx.Next()
	}

}
