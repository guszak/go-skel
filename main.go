package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.com/guszak/test/handlers"
	"gitlab.com/guszak/test/middlewares"
)

func main() {
	logFile := "testlogfile"
	port := "3001"
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		logFile = "D:\\home\\site\\wwwroot\\testlogfile"
		port = os.Getenv("HTTP_PLATFORM_PORT")
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		http.ListenAndServe(":"+port, nil)
	} else {
		defer f.Close()
		log.SetOutput(f)
		log.Println("--->   UP @ " + port + "  <------")
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	g := gin.Default()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	v1 := g.Group("v1")
	{
		v1.Use(middlewares.IsAuthenticated())
		v1.POST("/products", handlers.CreateProduct)
		v1.GET("/products", handlers.QueryProducts)
		v1.GET("/products/:id", handlers.GetProduct)
		v1.PUT("/products/:id", handlers.UpdateProduct)
		v1.DELETE("/products/:id", handlers.DeleteProduct)
	}
	http.Handle("/", g)

	http.ListenAndServe(":"+port, nil)
}
