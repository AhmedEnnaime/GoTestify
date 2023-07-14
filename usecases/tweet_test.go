package usecases

import (
	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/mocks"
	"github.com/stretchr/testify/suite"
)

type tweetUsecaseSuite struct {
	suite.Suite
	repository *mocks.TweetRepository
	usecase    TweetUsecase
}

func (suite *tweetUsecaseSuite) SetupSuite() {
	repository := new(mocks.TweetRepository)
	usecase := InitializeTweetUsecase(repository)

	suite.repository = repository
	suite.usecase = usecase

}

func (suite *tweetUsecaseSuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text:     "text",
	}

	suite.repository.On("CreateTweet", &tweet).Return(nil)
	err := suite.usecase.CreateTweet(&tweet)

	suite.Nil(err, "err is a nil pointer so no error in this process")
	suite.repository.AssertExpectations(suite.T())

}
