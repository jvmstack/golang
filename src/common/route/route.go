package route

import (
	"gin_template/src/controller"
	"github.com/gin-gonic/gin"
)

func PathRoute(r *gin.Engine) *gin.Engine {

	rootPath := r.Group("/gin")
	{
		userPath := rootPath.Group("/user")
		{
			controller.UserRegister(userPath)
		}
	}

	return r
}
