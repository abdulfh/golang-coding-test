package router

import (
	"fmt"
	"idn/controllers"
	"idn/services"
	"idn/utils/helper"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"idn/docs"
)

type GoRouter struct {
	GinFunc gin.HandlerFunc
	Router  *gin.Engine
}

// Server ..
func Server(listenAddress string) (err error) {
	debugMode := helper.GetEnv("APPS_DEBUG", "debug")
	gin.SetMode(debugMode)
	hrisRouter := GoRouter{}
	hrisRouter.Routers()

	err = http.ListenAndServe(listenAddress, hrisRouter.Router)

	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}
	fmt.Println("Routing successfully: ", listenAddress)

	return err
}

func (goRouter *GoRouter) Routers() {

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	generalServices := services.InitGeneralServices()
	generalController := controllers.InitGeneralController(generalServices)

	docs.SwaggerInfo.Title = "Golang Test Code Swagger"
	docs.SwaggerInfo.Description = "Documentation For Test Code"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", helper.GetEnv("SWAGGER_HOST", "localhost"), helper.GetEnv("SERVER_PORT", ":40001"))

	api := router.Group("/api/v1")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.GET("/fizzbuzz", generalController.FizzBuzz)
		api.GET("/multiple", generalController.Multiple)
		api.GET("/markpaid/:bill", generalController.MarkPaid)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	goRouter.Router = router
}
