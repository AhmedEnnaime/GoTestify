package server

import (
	"fmt"
	"net/http"

	"github.com/AhmedEnnaime/GoTestify/config"
	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/utils"
	"github.com/gin-gonic/gin"
)

func rootHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, entities.Response{
			Success: true,
			Message: "Hello world",
			Data:    struct{}{},
		})
	}
}

func registerRoutes(router *gin.Engine, handlers *Handlers) {
	serveHttp := utils.ServeHTTP

	router.GET("/", rootHandler())
	router.POST("/tweets", serveHttp(handlers.TweetHandler.CreateTweet))

}

func SetupServer() {
	fmt.Println("Setting up server")

	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	panic("all good")

	repos := SetupRepositories(db)
	uscs := SetupUsecases(repos)
	handlers := SetupHandlers(uscs)

	router := gin.Default()
	registerRoutes(router, handlers)

	router.Run(":9090")

}
