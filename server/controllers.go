package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func CreateTask(c *gin.Context) {
	var newBlogDetail BlogDetailJSON
	if err := c.BindJSON(&newBlogDetail); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		err = CallCreateTask(newBlogDetail.Content)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
		}
	}
}
