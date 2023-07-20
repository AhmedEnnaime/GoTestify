package server

import (
	"github.com/AhmedEnnaime/GoTestify/repositories"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	TweetRepository *repositories.TweetRepository
}

func SetupRepositories(db *sqlx.DB) *Repositories {
	tweetRepository := repositories.InitializeTweetRepository(db)

	return &Repositories{
		TweetRepository: &tweetRepository,
	}

}
