package main

import (
	"assigment2/config"
	"assigment2/controllers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var PORT = os.Getenv("PORTAPP")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	PORT = ":" + os.Getenv("PORTAPP")
	config.ConnectGorm()

	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hallo ini cuma main"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/order", controllers.GetAllOrders)
	r.POST("/addorder", controllers.AddOrder)
	r.PUT("/updateorder/:id", controllers.UpdateOrder)
	r.DELETE("/deleteorder/:id", controllers.DeleteOrder)

	r.Logger.Fatal(r.Start(PORT))
}
