package repositories

import (
	"errors"

	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/jmoiron/sqlx"
)

type tweetRepository struct {
	db *sqlx.DB
}

type TweetRepository interface {
	// GetAllTweets() (*[]entities.Tweet, error)
	CreateTweet(tweet *entities.Tweet) error
	// GetTweetById(id int) (*entities.Tweet, error)
	// SearchTweetByText(text string) (*[]entities.Tweet, error)
	// CreateTweet(tweet *entities.Tweet) error
	// UpdateTweet(tweet *entities.Tweet) error
	// DeleteTweet(id int) error
}

func InitializeTweetRepository(db *sqlx.DB) TweetRepository {
	return &tweetRepository{db}
}

// func (repository *tweetRepository) GetAllTweets() (*[]entities.Tweet, error) {
// 	var result []entities.Tweet

// 	rows, err := repository.db.Queryx(`SELECT id, username, text, created_at, modified_at FROM tweets`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var tweet entities.Tweet
// 		err = rows.StructScan(&tweet)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, tweet)
// 	}
// 	return &result, err
// }

// func (repository *tweetRepository) GetTweetById(id int) (*entities.Tweet, error) {
// 	// Implementation for GetTweetById
// }

// func (repository *tweetRepository) SearchTweetByText(text string) (*[]entities.Tweet, error) {
// 	// Implementation for SearchTweetByText
// }

func (repository *tweetRepository) CreateTweet(tweet *entities.Tweet) error {
	var err error

	if tweet == nil {
		return errors.New("tweet can not be nil")
	}
	tx, err := repository.db.Beginx()
	if err != nil {
		return err
	} else {
		err = insertTweet(tx, tweet)
		if err != nil {
			return err
		}
	}

	if err == nil {
		tx.Commit()

	} else {
		tx.Rollback()
	}

	return err

}

func insertTweet(tx *sqlx.Tx, tweet *entities.Tweet) error {
	_, err := tx.NamedExec(`INSERT INTO tweets(username, text)
	VALUES (:username, :text);`, tweet)
	return err

}

// func (repository *tweetRepository) UpdateTweet(tweet *entities.Tweet) error {
// 	// Implementation for UpdateTweet
// }

// func (repository *tweetRepository) DeleteTweet(id int) error {
// 	// Implementation for DeleteTweet
// }
