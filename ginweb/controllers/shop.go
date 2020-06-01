package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "shop %s", "hello")
}

func GoodsHandler(c *gin.Context) {
	c.String(http.StatusOK, "shop %s", "goods")
}

func CheckoutHandler(c *gin.Context) {
	c.String(http.StatusOK, "shop %s", "checkout")
}
