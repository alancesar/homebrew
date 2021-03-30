package main

import (
	"github.com/alancesar/homebrew/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Handle(http.MethodPost, "/recipe", handler.RecipeHandler())
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
