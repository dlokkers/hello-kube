package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes handles the gin routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
