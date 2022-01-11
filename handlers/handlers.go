package handlers

import (
	"github.com/gin-gonic/gin"
	"editt/database"
	"editt/models"
	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

/* Get All Posts */
func (a *APIEnv) GetPosts(c *gin.Context) {
	posts, err := database.GetPosts(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

/* Create Post */
func (a *APIEnv) CreatePost(c *gin.Context) {
	post := models.Post{}
	err := c.BindJSON(&post)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

/* Get Singe Post */
func (a *APIEnv) GetPost(c *gin.Context) {
	id := c.Params.ByName("id")
	post, exists, err := database.GetPostByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no book in db with this ID")
		return
	}

	c.JSON(http.StatusOK, post)
}


/* Delete Post */
func (a *APIEnv)DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetPostByID(id,a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "There is no any Posts with this ID")
		return
	}

	err = database.DeletePost(id,a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Post has been deleted successfully")
}

/* Update Post */
func (a *APIEnv) UpdatePost(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetPostByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "There is no any Posts with this ID")
		return
	}

	updatePost := models.Post{}
	err = c.BindJSON(&updatePost)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return	
	}

	if err := database.UpdatePost(a.DB, &updatePost); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetPost(c)
}
