package user_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	mockRepository "github.com/wisaitas/standard-golang/internal/standard-service/mocks/repository"
	mockUtil "github.com/wisaitas/standard-golang/internal/standard-service/mocks/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/service/user"
)

type createUserTestSuite struct {
	suite.Suite
	mockRepo  *mockRepository.MockUserRepository
	mockRedis *mockUtil.MockRedis
	service   user.Post
}

func (s *createUserTestSuite) SetupTest() {
	s.mockRepo = new(mockRepository.MockUserRepository)
	s.mockRedis = new(mockUtil.MockRedis)
	s.service = user.NewPost(s.mockRepo, s.mockRedis)
}

func (s *createUserTestSuite) TestCreateUserSuccess() {
	s.mockRepo.On("Create", mock.MatchedBy(func(u interface{}) bool {
		return true
	})).Return(nil)

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, status)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *createUserTestSuite) TestCreateUserUsernameExists() {
	s.mockRepo.On("Create", mock.Anything).Return(errors.New("ERROR: duplicate key value violates unique constraint"))

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "existinguser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusBadRequest, status)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *createUserTestSuite) TestCreateUserInternalServerError() {
	s.mockRepo.On("Create", mock.Anything).Return(errors.New("database error"))

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, status)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *createUserTestSuite) TestCreateUserHashError() {
	longPassword := string(make([]byte, 73))

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        longPassword,
		ConfirmPassword: longPassword,
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, status)
	s.mockRepo.AssertNotCalled(s.T(), "Create")
}

func TestCreateUser(t *testing.T) {
	suite.Run(t, new(createUserTestSuite))
}
