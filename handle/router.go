package handle

import (
	_ "MoeBlog/docs"
	"MoeBlog/handle/controller"
	"MoeBlog/interceptor"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @host 127.0.0.1:8080
// @BasePath /api/v1
func InitRoutes(r *gin.Engine) {
	r.Use(interceptor.CORS())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controller.UserRegister)
		v1.POST("/login", controller.UserLogin)
	}
}
