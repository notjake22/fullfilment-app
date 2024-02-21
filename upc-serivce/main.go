package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"main/cache"
	"main/handlers"
	"os"
)

func main() {
	r := gin.Default()

	var c cache.Cache = &cache.CRef{}
	go c.Init()

	r.GET("/upc-lookup", func(ctx *gin.Context) {
		handlers.LookUpUPC(ctx, c)
	})

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting HTTP endpoints")
	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}
