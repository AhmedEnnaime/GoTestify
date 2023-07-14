package handlers

import (
	"net/http"

	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/usecases"
	"github.com/gin-gonic/gin"
)

type tweetHandler struct {
	tweetUsecase usecases.TweetUsecase
}

type TweetHandler interface {
	CreateTweet(ctx *gin.Context) *entities.AppResult
}

func InitializeTweetHandler(usecase usecases.TweetUsecase) TweetHandler {
	return &tweetHandler{usecase}
}

func (handler *tweetHandler) CreateTweet(ctx *gin.Context) *entities.AppResult {
	var tweet entities.Tweet
	var result entities.AppResult

	if err := ctx.ShouldBindJSON(&tweet); err != nil {
		result.Err = err
		result.Message = "username and text cannot be empty"
		result.StatusCode = http.StatusBadRequest
		return &result
	}

	err := handler.tweetUsecase.CreateTweet(&tweet)

	if err == nil {
		result.Message = "Success to create tweet"
		result.StatusCode = http.StatusCreated
	} else {
		result.Err = err.(*entities.AppError).Err
		result.Message = err.(*entities.AppError).Error()
		result.StatusCode = err.(*entities.AppError).StatusCode
	}

	return &result

}
