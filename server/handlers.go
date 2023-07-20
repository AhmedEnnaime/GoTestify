package server

import "github.com/AhmedEnnaime/GoTestify/handlers"

type Handlers struct {
	TweetHandler handlers.TweetHandler
}

func SetupHandlers(uscs *Usecases) *Handlers {
	tweetHandlers := handlers.InitializeTweetHandler(uscs.TweetUsecase)

	return &Handlers{
		TweetHandler: tweetHandlers,
	}

}
