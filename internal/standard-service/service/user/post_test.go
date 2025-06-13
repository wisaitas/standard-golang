package user_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	mockPkg "github.com/wisaitas/standard-golang/internal/standard-service/mock/pkg"
	mockRepo "github.com/wisaitas/standard-golang/internal/standard-service/mock/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/service/user"

	"testing"
)

type createUserTestSuite struct {
	suite.Suite
	mockRepo   *mockRepo.MockUserRepository
	mockBcrypt *mockPkg.MockBcrypt
	mockRedis  *mockPkg.MockRedis
	service    user.Post
}

func (s *createUserTestSuite) SetupTest() {
	s.mockRepo = new(mockRepo.MockUserRepository)
	s.mockBcrypt = new(mockPkg.MockBcrypt)
	s.mockRedis = new(mockPkg.MockRedis)
	s.service = user.NewPost(s.mockRepo, s.mockRedis)
}

func (s *createUserTestSuite) TestCreateUserSuccess() {

	s.mockRepo.On("Create", mock.MatchedBy(func(u *entity.User) bool {
		return u.Username == "testuser" && u.Email == "test@example.com"
	})).Return(nil)

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, status)
}

func (s *createUserTestSuite) TestCreateUserUsernameExists() {
	s.mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("ERROR: duplicate key value violates unique constraint"))

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "existinguser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusBadRequest, status)
}

func (s *createUserTestSuite) TestCreateUserInternalServerError() {
	s.mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("database error"))

	_, status, err := s.service.CreateUser(request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	})

	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, status)
}

func (s *createUserTestSuite) TestCreateUserHashError() {
	s.mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

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
