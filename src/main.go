package main

import (
	"gin_template/src/common/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r = route.PathRoute(r)

	r.Run()
}
