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
	PayloadKey string = "token-payload"
)

func AuthCheck(tokenParser token.Parser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader(AuthHeader)
		token, ok := strings.CutPrefix(tokenHeader, "Bearer ")
		if !ok {
			ctx.AbortWithError(403, ErrAuthHeader)
			return
		}

		payload, err := tokenParser.Unpack(token)
		if err != nil {
			ctx.AbortWithError(403, ErrTokenParse)
			return
		}

		ctx.Set(PayloadKey, payload)

		ctx.Next()
	}
}
