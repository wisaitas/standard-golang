package user_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	mock_repositories "github.com/wisaitas/standard-golang/internal/mocks/repositories"
	mock_utils "github.com/wisaitas/standard-golang/internal/mocks/utils"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/services/user"

	"testing"
)

type createUserTestSuite struct {
	suite.Suite
	mockRepo   *mock_repositories.MockUserRepository
	mockBcrypt *mock_utils.MockBcrypt
	mockRedis  *mock_utils.MockRedis
	service    user.Create
}

func (s *createUserTestSuite) SetupTest() {
	s.mockRepo = new(mock_repositories.MockUserRepository)
	s.mockBcrypt = new(mock_utils.MockBcrypt)
	s.mockRedis = new(mock_utils.MockRedis)
	s.service = user.NewCreate(s.mockRepo, s.mockRedis)
}

func (s *createUserTestSuite) TestCreateUserSuccess() {

	s.mockRepo.On("Create", mock.MatchedBy(func(u *models.User) bool {
		return u.Username == "testuser" && u.Email == "test@example.com"
	})).Return(nil)

	_, status, err := s.service.CreateUser(requests.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, status)
}

func (s *createUserTestSuite) TestCreateUserUsernameExists() {
	s.mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(errors.New("ERROR: duplicate key value violates unique constraint"))

	_, status, err := s.service.CreateUser(requests.CreateUserRequest{
		Username:        "existinguser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusBadRequest, status)
}

func (s *createUserTestSuite) TestCreateUserInternalServerError() {
	s.mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(errors.New("database error"))

	_, status, err := s.service.CreateUser(requests.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, status)
}

func (s *createUserTestSuite) TestCreateUserHashError() {
	s.mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	longPassword := string(make([]byte, 73))

	_, status, err := s.service.CreateUser(requests.CreateUserRequest{
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
