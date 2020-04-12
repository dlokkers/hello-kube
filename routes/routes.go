package routes

import (
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/dlokkers/hello-kube/protobuf"
	"github.com/gin-gonic/gin"
)

// Init sets up the supported gin routes
func Init(router *gin.Engine, client pb.ArticlesClient) {
	router.GET("/", showIndexPage(client))
	router.GET("/article/view/:article_id", getArticle(client))
}

// showIndexPage is the route into the index page
func showIndexPage(ac pb.ArticlesClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &pb.Empty{}
		articles, err := ac.ListArticles(c, req)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":   "Home Page",
				"payload": articles.GetArticles(),
			},
		)
	}
}

func getArticle(ac pb.ArticlesClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		articleID, err := strconv.Atoi(c.Param("article_id"))
		fmt.Println(articleID)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		req := &pb.ArticleID{Id: uint32(articleID)}
		article, err := ac.RetrieveArticle(c, req)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(
			http.StatusOK,
			"article.html",
			gin.H{
				"title":   article.Title,
				"ID":      article.Id,
				"content": article.Content,
			},
		)
	}
}
