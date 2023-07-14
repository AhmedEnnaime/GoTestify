package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/AhmedEnnaime/GoTestify/mocks"
	"github.com/AhmedEnnaime/GoTestify/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type tweetHandlerSuite struct {
	suite.Suite
	usecase       *mocks.TweetUsecase
	handler       TweetHandler
	testingServer *httptest.Server
}

func (suite *tweetHandlerSuite) SetupSuite() {
	usecase := new(mocks.TweetUsecase)
	handler := InitializeTweetHandler(usecase)

	router := gin.Default()
	router.POST("/tweet", utils.ServeHTTP(handler.CreateTweet))

	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer
	suite.usecase = usecase
	suite.handler = handler

}

func (suite *tweetHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *tweetHandlerSuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text:     "text",
	}

	suite.usecase.On("CreateTweet", &tweet).Return(nil)

	requestBody, err := json.Marshal(&tweet)
	suite.NoError(err, "cannot marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/tweet", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := entities.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal(responseBody.Message, "Success to create tweet")
	suite.usecase.AssertExpectations(suite.T())

}
