package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golanguzb71/livesphere-api-gateway/api/helpers"
	coreService "github.com/golanguzb71/livesphere-api-gateway/genproto/core_service"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/etc"
)

func (h *handlerV1) CreateLead(c *gin.Context) {
	title := c.Query("title")
	ctx, cancelFunc := etc.NewTimoutContext(c)
	defer cancelFunc()
	response, err := h.services.CoreService().LeadService(ctx).Create(ctx, &coreService.Lead{
		Title: title,
	})
	if helpers.HandleGrpcErrWithMessage(c, h.log, err, "Error while create lead") {
		return
	}
	c.JSON(200, response)
}

func (h *handlerV1) GetLeadCommon(c *gin.Context) {
	var (
		body *coreService.GetListFilter
	)
	err := c.ShouldBindJSON(&body)
	if helpers.HandleBadRequestErrWithMessage(c, h.log, err, "Invalid request body") {
		return
	}
	ctx, cancelFunc := etc.NewTimoutContext(c)
	defer cancelFunc()
	response, err := h.Services().CoreService().LeadService(ctx).GetList(ctx, body)
	if helpers.HandleGrpcErrWithMessage(c, h.log, err, "Error while get lead list") {
		return
	}
	c.JSON(200, response)
}

func (h *handlerV1) UpdateLead(c *gin.Context) {
	var (
		body *coreService.Lead
	)

}

func (h *handlerV1) DeleteLead(c *gin.Context) {

}

func (h *handlerV1) GetAllLead(c *gin.Context) {

}

func (h *handlerV1) GetLeadReports(context *gin.Context) {

}
func (h *handlerV1) CreateExpectation(context *gin.Context) {

}

func (h *handlerV1) UpdateExpectation(context *gin.Context) {

}

func (h *handlerV1) DeleteExpectation(context *gin.Context) {

}

func (h *handlerV1) CreateSet(context *gin.Context) {

}

func (h *handlerV1) UpdateSet(context *gin.Context) {

}

func (h *handlerV1) DeleteSet(context *gin.Context) {

}

func (h *handlerV1) ChangeToSet(context *gin.Context) {

}

func (h *handlerV1) GetByIdSet(context *gin.Context) {

}

func (h *handlerV1) CreateLeadData(context *gin.Context) {

}

func (h *handlerV1) UpdateLeadData(context *gin.Context) {

}

func (h *handlerV1) DeleteLeadData(context *gin.Context) {

}

func (h *handlerV1) ChangeLeadData(context *gin.Context) {

}
