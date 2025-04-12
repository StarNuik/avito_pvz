package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/token"
)

var (
	ErrAuthHeader = fmt.Errorf("incorrect auth header")
	ErrTokenParse = fmt.Errorf("couldnt parse token")
)

const (
	AuthHeader string = "Authorization"
	payloadKey string = "token-payload"
)

func UnpackBearerToken(tokenParser token.Parser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader(AuthHeader)
		token, found := strings.CutPrefix(tokenHeader, "Bearer ")
		if !found {
			ctx.AbortWithError(403, ErrAuthHeader)
			return
		}

		payload, err := tokenParser.Unpack(token)
		if err != nil {
			ctx.AbortWithError(403, ErrTokenParse)
			return
		}

		ctx.Set(payloadKey, payload)

		ctx.Next()
	}
}
