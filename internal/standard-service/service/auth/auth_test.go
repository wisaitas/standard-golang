package auth_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	mockRepository "github.com/wisaitas/standard-golang/internal/standard-service/mocks/repository"
	mockUtil "github.com/wisaitas/standard-golang/internal/standard-service/mocks/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/service/auth"
	"github.com/wisaitas/standard-golang/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// MockDB implements a mock DB with Begin method
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Begin() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

type registerTestSuite struct {
	suite.Suite
	mockUserRepo          *mockRepository.MockUserRepository
	mockUserHistoryRepo   *mockRepository.MockUserHistoryRepository
	mockRedis             *mockUtil.MockRedis
	mockBcrypt            *mockUtil.MockBcrypt
	mockDB                *MockDB
	mockTxDB              *gorm.DB
	mockTxManager         *pkg.TxManager
	mockTxUserRepo        *mockRepository.MockUserRepository
	mockTxUserHistoryRepo *mockRepository.MockUserHistoryRepository
	service               auth.AuthService
}

func (s *registerTestSuite) SetupTest() {
	s.mockUserRepo = new(mockRepository.MockUserRepository)
	s.mockUserHistoryRepo = new(mockRepository.MockUserHistoryRepository)
	s.mockRedis = new(mockUtil.MockRedis)
	s.mockBcrypt = new(mockUtil.MockBcrypt)
	s.mockTxUserRepo = new(mockRepository.MockUserRepository)
	s.mockTxUserHistoryRepo = new(mockRepository.MockUserHistoryRepository)
	s.mockDB = new(MockDB)
	s.mockTxDB = &gorm.DB{}

	s.service = auth.NewAuthService(
		s.mockUserRepo,
		s.mockUserHistoryRepo,
		s.mockRedis,
		s.mockBcrypt,
	)
}

func (s *registerTestSuite) TestRegisterSuccess() {
	// Test data
	birthDate := time.Now()
	provinceID := uuid.New()
	districtID := uuid.New()
	subDistrictID := uuid.New()
	address := "123 Test St"

	registerReq := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
		Addresses: []request.AddressRequest{
			{
				ProvinceID:    provinceID,
				DistrictID:    districtID,
				SubDistrictID: subDistrictID,
				Address:       &address,
			},
		},
	}

	// Mock bcrypt password hashing
	hashedPassword := []byte("hashedpassword")
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return(hashedPassword, nil)

	// Mock DB and transaction setup
	s.mockDB.On("Begin").Return(s.mockTxDB)
	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	// Mock transaction repositories
	s.mockUserRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserHistoryRepo)

	// Mock first user history creation
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(nil)

	// Mock user creation
	s.mockTxUserRepo.On("Create", mock.MatchedBy(func(user *entity.User) bool {
		return user.Password == string(hashedPassword)
	})).Return(nil)

	// Mock second user history creation
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(nil).Once()

	// Mock transaction commit
	s.mockTxDB.Error = nil // Ensure no error for commit

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	s.mockTxUserRepo.AssertExpectations(s.T())
	s.mockTxUserHistoryRepo.AssertExpectations(s.T())
	s.mockDB.AssertExpectations(s.T())
}

func (s *registerTestSuite) TestRegisterFailedBcrypt() {
	// Test data
	birthDate := time.Now()
	registerReq := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	// Mock bcrypt password hashing failure
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return([]byte{}, errors.New("bcrypt error"))

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	// Ensure transaction was not started
	s.mockUserRepo.AssertNotCalled(s.T(), "WithTxManager")
}

func (s *registerTestSuite) TestRegisterFailedFirstHistoryCreate() {
	// Test data
	birthDate := time.Now()
	registerReq := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	// Mock bcrypt password hashing
	hashedPassword := []byte("hashedpassword")
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return(hashedPassword, nil)

	// Mock DB and transaction setup
	s.mockDB.On("Begin").Return(s.mockTxDB)
	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	// Mock transaction repositories
	s.mockUserRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserHistoryRepo)

	// Mock first user history creation failure
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(errors.New("history creation error"))

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	s.mockTxUserHistoryRepo.AssertExpectations(s.T())
	s.mockDB.AssertExpectations(s.T())
	// Ensure user creation was not called
	s.mockTxUserRepo.AssertNotCalled(s.T(), "Create")
}

func (s *registerTestSuite) TestRegisterFailedUserCreate() {
	// Test data
	birthDate := time.Now()
	registerReq := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	// Mock bcrypt password hashing
	hashedPassword := []byte("hashedpassword")
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return(hashedPassword, nil)

	// Mock DB and transaction setup
	s.mockDB.On("Begin").Return(s.mockTxDB)
	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	// Mock transaction repositories
	s.mockUserRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserHistoryRepo)

	// Mock first user history creation
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(nil).Once()

	// Mock user creation failure
	s.mockTxUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("database error"))

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	s.mockTxUserRepo.AssertExpectations(s.T())
	s.mockTxUserHistoryRepo.AssertExpectations(s.T())
	s.mockDB.AssertExpectations(s.T())
}

func (s *registerTestSuite) TestRegisterUserExistsError() {
	// Test data
	birthDate := time.Now()
	registerReq := request.RegisterRequest{
		Username:        "existinguser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	// Mock bcrypt password hashing
	hashedPassword := []byte("hashedpassword")
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return(hashedPassword, nil)

	// Mock DB and transaction setup
	s.mockDB.On("Begin").Return(s.mockTxDB)
	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	// Mock transaction repositories
	s.mockUserRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserHistoryRepo)

	// Mock first user history creation
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(nil).Once()

	// Mock user creation with unique constraint error
	s.mockTxUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("ERROR: duplicate key value violates unique constraint"))

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().Error(err)
	s.Require().Equal(http.StatusBadRequest, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	s.mockTxUserRepo.AssertExpectations(s.T())
	s.mockTxUserHistoryRepo.AssertExpectations(s.T())
	s.mockDB.AssertExpectations(s.T())
}

func (s *registerTestSuite) TestRegisterSecondHistoryCreateFailed() {
	// Test data
	birthDate := time.Now()
	registerReq := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "Test",
		LastName:        "User",
		BirthDate:       birthDate,
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	// Mock bcrypt password hashing
	hashedPassword := []byte("hashedpassword")
	s.mockBcrypt.On("GenerateFromPassword", mock.Anything, bcrypt.DefaultCost).Return(hashedPassword, nil)

	// Mock DB and transaction setup
	s.mockDB.On("Begin").Return(s.mockTxDB)
	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	// Mock transaction repositories
	s.mockUserRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTxManager", mock.AnythingOfType("*pkg.TxManager")).Return(s.mockTxUserHistoryRepo)

	// Mock first user history creation
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(nil).Once()

	// Mock user creation
	s.mockTxUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

	// Mock second user history creation failure
	s.mockTxUserHistoryRepo.On("Create", mock.MatchedBy(func(history *entity.UserHistory) bool {
		return history.Action == constants.Action.Create
	})).Return(errors.New("second history creation error")).Once()

	// Execute the test
	_, statusCode, err := s.service.Register(registerReq)

	// Assertions
	s.Require().Error(err)
	s.Require().Equal(http.StatusInternalServerError, statusCode)
	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	s.mockTxUserRepo.AssertExpectations(s.T())
	s.mockTxUserHistoryRepo.AssertExpectations(s.T())
	s.mockDB.AssertExpectations(s.T())
}

func TestRegister(t *testing.T) {
	suite.Run(t, new(registerTestSuite))
}
