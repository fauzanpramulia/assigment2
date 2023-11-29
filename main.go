package main

import (
	"assigment2/config"
	"assigment2/controllers"
	"net/http"
	"github.com/labstack/echo/v4"
)


func main() {
	PORT := ":9000"
	config.ConnectGorm()

	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hallo ini Home"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/order", controllers.GetAllOrders)
	r.POST("/addorder", controllers.AddOrder)
	r.PUT("/updateorder/:id", controllers.UpdateOrder)
	r.DELETE("/deleteorder/:id", controllers.DeleteOrder)

	r.Logger.Fatal(r.Start(PORT))
}
