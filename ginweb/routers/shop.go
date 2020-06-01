package routers

import (
	"gindemo/ginweb/controllers"
	"github.com/gin-gonic/gin"
)

func LoadShop(e *gin.Engine) {
	e.GET("/hello", controllers.HelloHandler)
	e.GET("/goods", controllers.GoodsHandler)
	e.GET("/checkout", controllers.CheckoutHandler)
}
