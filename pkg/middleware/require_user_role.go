package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

var (
	ErrNoToken        = fmt.Errorf("no token")
	ErrRoleNotAllowed = fmt.Errorf("insufficient permissions")
)

func getTokenPayload(ctx *gin.Context) (token.Payload, bool) {
	untyped, ok1 := ctx.Get(payloadKey)
	payload, ok2 := untyped.(token.Payload)
	return payload, ok1 && ok2
}

func RequireUserRole(userRole entity.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, ok := getTokenPayload(ctx)
		if !ok {
			ctx.AbortWithError(403, ErrNoToken)
			return
		}

		if payload.UserRole != userRole {
			ctx.AbortWithError(403, ErrRoleNotAllowed)
			return
		}

		ctx.Next()
	}
}
