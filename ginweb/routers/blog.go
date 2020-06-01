package routers

import (
	"gindemo/ginweb/controllers"
	"github.com/gin-gonic/gin"
)

func LoadBlog(e *gin.Engine) {
	e.GET("/post", controllers.PostHandler)
	e.GET("/comment", controllers.CommentHandler)
}
