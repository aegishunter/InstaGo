package main

import (
	"instago/controller"
	"instago/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := routerSetup()

	postController := controller.NewPostController()
	router.GET("/", postController.GetAllPosts)
	router.GET("/:user_id/posts/:id", postController.GetPostDetailsById)
	router.DELETE("/:user_id/posts/:id", postController.DeletePostById)
	router.GET("/:user_id/edit/:id", postController.ShowUpdatePage)
	router.POST("/edit/description", postController.UpdateDescription)
	router.POST("/add", postController.AddPost)
	router.GET("/add", postController.NavigateAddPage)
	router.Run(":8080")
}

func routerSetup() *gin.Engine {
	router := gin.New()
	router.Use(middleware.GetLogger())
	router.Use(middleware.GetRecoverHandler())

	router.Static("/css", "./view/css")
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*.html")
	return router
}
