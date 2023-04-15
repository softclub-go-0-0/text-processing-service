package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/text-processing-service/pkg/handlers"
	"github.com/softclub-go-0-0/text-processing-service/pkg/middlewares"
)

func main() {
	fmt.Println("Todo Service API v0.0")

	port := flag.String("listenport", "4000", "Which port to listen")

	flag.Parse()
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware(""))

	router.GET("number-of-symbols", handlers.CalculateNumberOfSymbols)
	router.GET("number-of-words", handlers.CalculateNumberOfWords)
	router.GET("reading-time", handlers.CalculateReadingTime)

	router.Run(":" + *port)
}
