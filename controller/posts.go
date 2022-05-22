package controller

import (
	"fmt"
	"instago/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostsController struct {
	service *service.PostsService
}

func NewPostController() *PostsController {
	return &PostsController{
		service: service.NewPostsService(),
	}
}

func (p *PostsController) NavigateAddPage(c *gin.Context) {
	data := gin.H{
		"title":  "Add Post",
		"userId": 1,
	}
	c.HTML(200, "add.html", data)
}

func (p *PostsController) AddPost(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "no file uploaded",
		})
	}
	description := c.PostForm("description")
	userId, _ := strconv.Atoi(c.PostForm("userid"))

	err = p.service.AddPost(userId, file.Filename, description)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		errUpload := c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
		if errUpload != nil {
			c.AbortWithError(http.StatusInternalServerError, errUpload)
		} else {
			c.Redirect(http.StatusFound, "/")
		}
	}
}

func (p *PostsController) ShowUpdatePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	post, err := p.service.GetPostDetailsById(id, user_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		data := gin.H{
			"status": http.StatusOK,
			"post":   post,
		}
		c.HTML(http.StatusOK, "edit.html", data)
	}
}

func (p *PostsController) UpdateDescription(c *gin.Context) {
	description, id := c.PostForm("description"), c.PostForm("id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	err = p.service.UpdateDescription(description, post_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func (p *PostsController) GetAllPosts(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	if err != nil {
		idx = 1
	}
	fmt.Println(idx)
	allPosts, totalPostCount, pageCount, err := p.service.GetAllPosts(idx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		pageCount := pageCount
		data := gin.H{
			"title":     "Hello world",
			"posts":     allPosts,
			"count":     totalPostCount,
			"pageCount": pageCount,
		}
		c.HTML(http.StatusOK, "home.html", data)
	}

}

func (p *PostsController) DeletePostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	err = p.service.DeletePostById(id, user_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Redirect(http.StatusOK, "/")
	}
}

func (p *PostsController) GetPostDetailsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	selectedPost, err := p.service.GetPostDetailsById(id, user_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		data := gin.H{
			"title": "Selected Posts",
			"post":  selectedPost,
		}
		c.HTML(200, "details.html", data)
	}
}
