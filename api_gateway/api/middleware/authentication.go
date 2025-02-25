package middleware

import (
	v1 "github.com/golanguzb71/livesphere-api-gateway/api/handlers/v1"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golanguzb71/livesphere-api-gateway/api/models"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/jwt"
)

func Authentication(handlerOptions *v1.HandlerV1Options, cfg *config.Config) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		isOpen := config.OpenApis[ctx.FullPath()]
		if isOpen {
			ctx.Next()
			return
		}

		token := ctx.GetHeader("Authorization")
		if token != "" {
			if len(strings.Split(token, " ")) > 1 {
				token = strings.Split(token, " ")[1]
			}

			tokenInfo, err := jwt.JWTExtract(token, cfg.JWTSigningKey)
			if err != nil {
				ctx.JSON(401, models.Response{
					Status:  config.ErrorCodeUnauthorized,
					Message: err.Error(),
				})
				ctx.Abort()
				return
			}

			for key, value := range tokenInfo {
				ctx.Request.Header.Set(key, value)
				key1 := key
				value1 := value
				ctx.Set(key1, value1)
			}
		}

		ctx.Next()
	}
}
