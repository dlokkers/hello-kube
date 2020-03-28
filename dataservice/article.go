package routes

import (
	"net/http"
	"strconv"

	"github.com/dlokkers/hello-kube/model"
	"github.com/gin-gonic/gin"
)

// showIndexPage is the route into the index page
func showIndexPage(c *gin.Context) {
	articles := model.GetAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	article, err := model.GetArticle(articleID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   article.Title,
			"ID":      article.ID,
			"content": article.Content,
		},
	)
}
