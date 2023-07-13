package repositories

import (
	"github.com/AhmedEnnaime/GoTestify/config"
	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/utils"
	"github.com/stretchr/testify/suite"
)

type tweetRepositorySuite struct {
	suite.Suite
	repository      TweetRepository
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *tweetRepositorySuite) SetupSuite() {
	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	repository := InitializeTweetRepository(db)

	suite.repository = repository
	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)

}

func (suite *tweetRepositorySuite) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable([]string{"tweets"})
}

func (suite *tweetRepositorySuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text:     "text",
	}

	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when creating tweet with valid input")
}
