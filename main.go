package main

import (
	"Firstproject/pkg"
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoUri := "mongodb+srv://turabali57972:kKaYVPp0FCTptJj5@cluster0.sbl2ptv.mongodb.net/sample_restaurants?retryWrites=true&w=majority&appName=Cluster0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}
	srv := pkg.NewService(client)
	hdl := pkg.NewHandler(srv)

	e := echo.New()

	e.GET("/users", hdl.GetUsers)
	e.POST("/users", hdl.Createuser)

	e.Logger.Fatal(e.Start(":8000"))

}
