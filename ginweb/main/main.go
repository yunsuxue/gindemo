package main

import (
	"fmt"
	"gindemo/ginweb/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.LoadBlog(r)
	routers.LoadShop(r)

	if err := r.Run("127.0.0.1:8100"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
