package v1

import (
	"github.com/gin-gonic/gin"
	rediscache "github.com/golanguzb70/redis-cache"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/services"
	"strconv"
)

type handlerV1 struct {
	log      logger.Logger
	cfg      *config.Config
	services services.ServiceManager
	cache    rediscache.RedisCache
}

type HandlerV1Options struct {
	Log      logger.Logger
	Cfg      *config.Config
	Services services.ServiceManager
	Cache    rediscache.RedisCache
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:      options.Log,
		cfg:      options.Cfg,
		services: options.Services,
		cache:    options.Cache,
	}
}

func (h *handlerV1) Services() services.ServiceManager {
	return h.services
}

func (h *handlerV1) Log() logger.Logger {
	return h.log
}

func (h *handlerV1) Config() *config.Config {
	return h.cfg
}

func (h *handlerV1) Cache() rediscache.RedisCache {
	return h.cache
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 32)
	if err != nil {
		return 0, err
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}
