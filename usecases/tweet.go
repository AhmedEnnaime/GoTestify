package usecases

import (
	"errors"
	"net/http"

	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/repositories"
)

type tweetUsecase struct {
	tweetRepository repositories.TweetRepository
}

type TweetUsecase interface {
	CreateTweet(tweet *entities.Tweet) error
}

func InitializeTweetUsecase(repository repositories.TweetRepository) TweetUsecase {
	return &tweetUsecase{repository}
}

func (usecase *tweetUsecase) CreateTweet(tweet *entities.Tweet) error {
	if tweet == nil {
		return &entities.AppError{
			Err:        errors.New("tweet is nil pointer"),
			StatusCode: http.StatusInternalServerError,
		}

	}

	if !tweet.IsValid() {
		return &entities.AppError{
			Err:        errors.New("username and text cannot be empty"),
			StatusCode: http.StatusBadRequest,
		}
	}

	err := usecase.tweetRepository.CreateTweet(tweet)

	if err != nil {
		return &entities.AppError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil

}
