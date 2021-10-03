package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minhphong306/mindX/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: authen here
		ctx.Next()
	}
}
