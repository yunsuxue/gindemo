package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostHandler(c *gin.Context) {
	c.String(http.StatusOK, "blog %s", "post")
}

func CommentHandler(c *gin.Context) {
	c.String(http.StatusOK, "blog %s", "comment")
}
