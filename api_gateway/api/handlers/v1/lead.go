package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/etc"
)

func (h *handlerV1) CreateLead(c *gin.Context) {
	title := c.Query("title")
	ctx, cancelFunc := etc.NewTimoutContext(c)
	defer cancelFunc()

}

func (h *handlerV1) GetLeadCommon(c *gin.Context) {

}

func (h *handlerV1) UpdateLead(c *gin.Context) {

}

func (h *handlerV1) DeleteLead(c *gin.Context) {

}

func (h *handlerV1) GetAllLead(c *gin.Context) {

}

func (h *handlerV1) GetLeadReports(context *gin.Context) {

}
