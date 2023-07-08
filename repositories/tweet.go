package repositories

import (
	"github.com/jmoiron/sqlx"
)

type tweetRepository struct {
	db *sqlx.DB
}

type TweetRepository interface {
	CreateTweet()
}
