package url

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, controller *URLController) {
	urlGroup := router.Group("/url")
	{
		urlGroup.POST("/shorten", controller.ShortenURL)
		urlGroup.GET("/:short_url", controller.RedirectToLongURL)
	}
}
