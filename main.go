package main

import (
	"Firstproject/pkg"

	"github.com/labstack/echo/v4"
)

func main() {
	srv := pkg.NewService()
	hdl := pkg.NewHandler(srv)

	e := echo.New()

	e.GET("/users", hdl.GetUsers)
	e.POST("/users", hdl.Createuser)

	e.Logger.Fatal(e.Start(":8000"))

}
