package routes

import (
	"go-url-shortener/middlewares"
	"go-url-shortener/modules/url"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Exclude Swagger from rate limiting
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply rate limiting middleware only to API routes
	apiV1 := router.Group("/api/v1")
	apiV1.Use(middlewares.RateLimiterMiddleware()) // ✅ Apply middleware here, not globally

	// Initialize URL Module under api/v1
	urlRepo := url.NewURLRepository()
	urlService := url.NewURLService(urlRepo)
	urlController := url.NewURLController(urlService)
	url.RegisterRoutes(apiV1, urlController)

	return router
}
