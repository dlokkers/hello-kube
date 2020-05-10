package routes

import (
	"errors"
	"net/http"

	pb "github.com/dlokkers/hello-kube/protobuf"
	"github.com/gin-gonic/gin"
)

// Init sets up the supported gin routes
func Init(router *gin.Engine, client pb.RecipesClient) {
	router.GET("/", showIndexPage(client))
	router.GET("/recipe/view/:recipe_title", getRecipe(client))
	router.GET("/recipe/add/", addForm())
	router.POST("/recipe/add/", addRecipe(client))
}

// showIndexPage is the route into the index page
func showIndexPage(ac pb.RecipesClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &pb.Empty{}
		recipes, err := ac.ListRecipes(c, req)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		titleList := recipes.GetTitle()

		type Title struct {
			Title string
		}
		var titles []Title
		for _, t := range titleList {
			titles = append(titles, Title{Title: t})
		}

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":   "Home Page",
				"payload": titles,
			},
		)
	}
}

func getRecipe(rc pb.RecipesClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		recipeTitle := c.Param("recipe_title")

		req := &pb.RecipeTitle{Title: recipeTitle}
		recipe, err := rc.RetrieveRecipe(c, req)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(
			http.StatusOK,
			"recipe.html",
			gin.H{
				"title":       recipe.Title,
				"ingredients": recipe.Ingredients,
				"process":     recipe.Process,
			},
		)
	}
}

func addForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"add.html",
			gin.H{},
		)
	}
}

func addRecipe(rc pb.RecipesClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		title, ok := c.GetPostForm("title")
		process, ok := c.GetPostForm("process")

		if !ok {
			c.AbortWithError(http.StatusNotFound, errors.New("not ok"))
		}

		var ingredients []*pb.Ingredient
		ingredient := &pb.Ingredient{
			Amount: "0",
			Type:   "lemons",
		}
		ingredients = append(ingredients, ingredient)

		req := &pb.Recipe{
			Title:       title,
			Ingredients: ingredients,
			Process:     process,
		}

		_, err := rc.AddRecipe(c, req)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{},
		)
	}
}
