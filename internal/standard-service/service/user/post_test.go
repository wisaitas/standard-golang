package user_test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/service/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type postTestSuite struct {
	suite.Suite
	mockUserRepo *repository.MockBaseRepository[entity.User]
	mockRedis    *redisPkg.MockRedis
	service      user.Post
	mockDB       *gorm.DB
}

func (s *postTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	s.Require().NoError(err)
	s.mockDB = db

	s.mockUserRepo = new(repository.MockBaseRepository[entity.User])
	s.mockRedis = redisPkg.NewMockRedis()

	s.service = user.NewPost(
		s.mockUserRepo,
		s.mockRedis,
	)
}

func (s *postTestSuite) TestCreateUserSuccess() {
	createUserRequest := request.CreateUserRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*entity.User)

		user.Id = uuid.New()
	})

	resp, statusCode, err := s.service.CreateUser(createUserRequest)

	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, statusCode)
	s.Require().Equal("testuser", resp.Username)
	s.Require().Equal("test@example.com", resp.Email)
	s.Require().NotEmpty(resp.Id)

	s.mockUserRepo.AssertExpectations(s.T())
}

func TestPostTestSuite(t *testing.T) {
	suite.Run(t, new(postTestSuite))
}
