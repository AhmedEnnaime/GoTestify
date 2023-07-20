package server

import "github.com/AhmedEnnaime/GoTestify/usecases"

type Usecases struct {
	TweetUsecase usecases.TweetUsecase
}

func SetupUsecases(repos *Repositories) *Usecases {
	tweetUsecase := usecases.InitializeTweetUsecase(*repos.TweetRepository)

	return &Usecases{
		TweetUsecase: tweetUsecase,
	}

}
