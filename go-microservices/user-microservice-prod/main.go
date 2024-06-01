

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
	"os"
	"io"
	"barqi.com/user/common"
	"barqi.com/user/database"
	"barqi.com/user/controllers"
	docs "barqi.com/user/docs"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	err = common.LoadConfig()

	err = database.Database.Init()
	if err != nil {
		return nil
	}

	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout,f)
		}else{
			gin.DefaultWriter = io.MultiWriter(f)
		}
	}

	m.router = gin.Default()

	return nil
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:9003
// @BasePath /api/v1
func main(){
	m := Main{}

	if m.initServer() != nil {
		return
	}

	defer database.Database.Close()

	c := controllers.User{}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth",c.Authenticate)
		}

		user := v1.Group("/users")

		user.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			user.POST("",c.AddUser)
			user.GET("/list",c.ListUsers)
			user.GET("detail/:id",c.GetUserByID)
			user.GET("/", c.GetUsersByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PATCH(":id", c.UpdateUser)
		}
	}

	m.router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	m.router.Run(common.Config.Port)
}