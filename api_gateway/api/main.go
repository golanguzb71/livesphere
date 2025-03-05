package api

import (
	"github.com/gin-gonic/gin"
	rediscache "github.com/golanguzb70/redis-cache"
	v1 "github.com/golanguzb71/livesphere-api-gateway/api/handlers/v1"
	"github.com/golanguzb71/livesphere-api-gateway/api/middleware"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/services"

	_ "github.com/golanguzb71/livesphere-api-gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      *config.Config
	Services services.ServiceManager
	Cache    rediscache.RedisCache
}

// @title Eld API Gateway v1
// @version 1.0
// @description This is a Eld API Gateway server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	options := &v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
		Cache:    opt.Cache,
	}

	handlerV1 := v1.New(options)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	v1Group := router.Group("/v1")
	v1Group.Use(middleware.Authentication(options, opt.Cfg))
	lead := v1Group.Group("/lead")
	{
		lead.POST("/create", handlerV1.CreateLead)
		lead.POST("/get-lead-common", handlerV1.GetLeadCommon)
		lead.PUT("/update/:id", handlerV1.UpdateLead)
		lead.DELETE("/delete/:id", handlerV1.DeleteLead)
		lead.GET("/get-all", handlerV1.GetAllLead)
		lead.GET("/get-lead-reports", handlerV1.GetLeadReports)
	}
	expectation := v1Group.Group("/expectation")
	{
		expectation.POST("/create", handlerV1.CreateExpectation)
		expectation.PUT("/update/:id", handlerV1.UpdateExpectation)
		expectation.DELETE("/delete/:id", handlerV1.DeleteExpectation)
	}
	set := v1Group.Group("/set")
	{
		set.POST("/create", handlerV1.CreateSet)
		set.PUT("/update", handlerV1.UpdateSet)
		set.DELETE("/delete/:id", handlerV1.DeleteSet)
		set.PATCH("/change-to-group", handlerV1.ChangeToSet)
		set.GET("/get-by-id/:id", handlerV1.GetByIdSet)
	}
	leadData := v1Group.Group("leadData")
	{
		leadData.POST("/create", handlerV1.CreateLeadData)
		leadData.PUT("/update", handlerV1.UpdateLeadData)
		leadData.DELETE("/delete/:id", handlerV1.DeleteLeadData)
		leadData.PATCH("/change-lead-data", handlerV1.ChangeLeadData)
	}

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, HEAD, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
