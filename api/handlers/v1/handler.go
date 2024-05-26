package v1

import (
	"github/test-ci-cd-github/api/handlers/token"
	"github/test-ci-cd-github/config"
	"github/test-ci-cd-github/pkg/logger"
	"github/test-ci-cd-github/services"

)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     token.JWTHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     token.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
	}
}
