package main

import (
	"echo/config"
	"echo/controller"
	echoMidd "echo/middleware"

	_ "echo/docs"

	"github.com/labstack/echo/v4"
	middEcho "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type M map[string]interface{} //JSON nya

// @title API Documentation Employee
// @version 1.0.0
// @description Coba coba API Swagger
// @contact.Name a
// @contact.Email a@a.com
// @host localhost:9000
// @BasePath /
func main() {

	config.Connect()
	// tmpl := template.Must(template.ParseGlob("template/*.html"))
	e := echo.New()

	e.Use(middEcho.CSRFWithConfig(middEcho.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	e.GET("/index", controller.Index)
	e.POST("/sayhello", controller.SayHello)

	//routes for middleware

	// e.Use(echoMidd.Authentication())

	item := e.Group("/item")
	item.Use(echoMidd.Authentication())
	item.POST("/", controller.CreateItem)

	emp := e.Group("/emp")
	emp.Use(echoMidd.Authentication())
	emp.GET("/", controller.HelloWorld)

	e.GET("/json", controller.JsonMap)
	e.GET("/page1", controller.Param)
	e.Any("/user", controller.User)
	e.POST("/create", controller.CreateUser)
	e.Any("/update", controller.UpdateEmployee)
	e.DELETE("/delete/:id", controller.DeleteEmployee)
	e.POST("/login", controller.UserLogin)

	v := e.Group("/swagger")

	//route for swagger
	v.GET("/*", echoSwagger.WrapHandler)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// e.GET("/json", func(c echo.Context) error {
	// 	data := M{
	// 		"message":    "Hello",
	// 		"counter":    2,
	// 		"statusCode": http.StatusOK,
	// 	}
	// 	return c.JSON(http.StatusOK, data)
	// })

	// e.GET("/page1", func(c echo.Context) error {
	// 	name := c.QueryParam("name")
	// 	data := "Hello " + name
	// 	result := fmt.Sprintf("%s", data)
	// 	return c.JSON(http.StatusOK, result)
	// })

	// e.GET("/page3", func(c echo.Context) error {
	// 	name := c.QueryParam("name")
	// 	data := "Hello " + name
	// 	result := fmt.Sprintf("%s", data)
	// 	return c.JSON(http.StatusOK, result)
	// })

	// e.Static(/static)

	e.Logger.Fatal(e.Start(":9000"))
}
