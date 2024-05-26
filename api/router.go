package api

import (
	_ "github/test-ci-cd-github/api/docs" // swag
	"github/test-ci-cd-github/api/handlers/token"
	v1 "github/test-ci-cd-github/api/handlers/v1"
	"github/test-ci-cd-github/config"
	"github.com/gin-contrib/cors"
	"github/test-ci-cd-github/pkg/logger"
	"github/test-ci-cd-github/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @Title Microservices architecture example
// @Security Auth
// @Version 1.0
// @Description This is a example of Social Network
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SigningKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		JWTHandler:     jwtHandler,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, option.Conf))

	apiV1 := router.Group("/v1")

	// Users...
	apiV1.GET("/ping/:id", handlerV1.GetUser)
	apiV1.GET("/pong/:id", handlerV1.TestPong)
	apiV1.GET("/github/:id", handlerV1.TestGithub)


	url := ginSwagger.URL("swagger/doc.json")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
