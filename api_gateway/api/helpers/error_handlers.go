package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb71/livesphere-api-gateway/api/models"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  config.ErrorCodeInternal,
			Message: st.Message(),
		})
		l.Error(message, logger.Error(err))
		return true
	}

	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, models.Response{
			Status:  config.ErrorCodeNotFound,
			Message: st.Message(),
		})
		l.Error(message+", not found", logger.Error(err))
		return true
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  config.ErrorCodeInternal,
			Message: "Internal Server Error",
		})
		l.Error(message+", service unavailable", logger.Error(err))
		return true
	} else if st.Code() == codes.AlreadyExists {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  config.ErrorCodeAlreadyExists,
			Message: fmt.Sprintf("%s is alredy exists", strings.ToTitle(strings.ReplaceAll(st.Message(), "_key", ""))),
		})

		l.Error(message+", already exists", logger.Error(err))
		return true
	} else if st.Code() == codes.InvalidArgument {
		if st.Message() == config.ErrorCodeSessionLimit {
			c.JSON(http.StatusBadRequest, models.Response{
				Status:  config.ErrorCodeSessionLimit,
				Message: st.Message(),
			})
			l.Error(message+", session limit", logger.Error(err))
			return true
		} else if st.Message() == config.ErrorCodeInvalidCode {
			c.JSON(http.StatusBadRequest, models.Response{
				Status:  config.ErrorCodeInvalidCode,
				Message: st.Message(),
			})
			l.Error(message+", Invalid OTP code", logger.Error(err))
			return true
		}

		c.JSON(http.StatusBadRequest, models.Response{
			Status:  config.ErrorBadRequest,
			Message: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Code() == codes.Code(20) {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  config.ErrorBadRequest,
			Message: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Err() != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  config.ErrorBadRequest,
			Message: st.Message(),
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	}
	return false
}

func HandleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  config.ErrorCodeInternal,
			Message: "Internal Server Error",
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func HandleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  config.ErrorCodeInvalidJSON,
			Message: "Invalid Json",
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}

func HandleUserStatusErrWithMessage(c *gin.Context, l logger.Logger, message string) bool {
	c.JSON(http.StatusBadRequest, models.Response{
		Status:  config.ErrorCodeInvalidJSON,
		Message: "in_verify",
	})
	return true
}
