package auth_test

// import (
// 	"errors"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
// 	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
// 	mockPkg "github.com/wisaitas/standard-golang/internal/standard-service/mock/pkg"
// 	mockRepo "github.com/wisaitas/standard-golang/internal/standard-service/mock/repository"
// 	"github.com/wisaitas/standard-golang/internal/standard-service/service/auth"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type authRegisterTestSuite struct {
// 	suite.Suite
// 	mockUserRepo        *mockRepo.MockUserRepository
// 	mockUserHistoryRepo *mockRepo.MockUserHistoryRepository
// 	mockRedis           *mockPkg.MockRedis
// 	mockBcrypt          *mockPkg.MockBcrypt
// 	mockJWT             *mockPkg.MockJwt
// 	mockTxManager       *mockPkg.MockTransactionManager
// 	service             auth.AuthService
// }

// func (s *authRegisterTestSuite) SetupTest() {
// 	s.mockUserRepo = new(mockRepo.MockUserRepository)
// 	s.mockUserHistoryRepo = new(mockRepo.MockUserHistoryRepository)
// 	s.mockRedis = new(mockPkg.MockRedis)
// 	s.mockBcrypt = new(mockPkg.MockBcrypt)
// 	s.mockJWT = new(mockPkg.MockJwt)
// 	s.mockTxManager = new(mockPkg.MockTransactionManager)

// 	s.service = auth.NewAuthService(
// 		s.mockUserRepo,
// 		s.mockUserHistoryRepo,
// 		s.mockRedis,
// 		s.mockBcrypt,
// 		s.mockJWT,
// 		s.mockTxManager,
// 	)
// }

// func (s *authRegisterTestSuite) TestRegisterSuccess() {
// 	testTime := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
// 	provinceID := uuid.New()
// 	districtID := uuid.New()
// 	subDistrictID := uuid.New()
// 	testAddress := "123 Test Street"

// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       testTime,
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses: []request.AddressRequest{
// 			{
// 				ProvinceID:    provinceID,
// 				DistrictID:    districtID,
// 				SubDistrictID: subDistrictID,
// 				Address:       &testAddress,
// 			},
// 		},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(nil).Run(func(args mock.Arguments) {
// 		txFunc := args.Get(0).(func(tx *gorm.DB) error)
// 		mockTx := &gorm.DB{}

// 		s.mockUserRepo.On("WithTx", mockTx).Return(s.mockUserRepo)
// 		s.mockUserHistoryRepo.On("WithTx", mockTx).Return(s.mockUserHistoryRepo)

// 		s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
// 			user := args.Get(0).(*entity.User)
// 			user.ID = uuid.New()
// 		})
// 		s.mockUserHistoryRepo.On("Create", mock.AnythingOfType("*entity.UserHistory")).Return(nil)

// 		txFunc(mockTx)
// 	})

// 	resp, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().NoError(err)
// 	s.Require().Equal(http.StatusCreated, statusCode)
// 	s.Require().Equal("testuser", resp.Username)
// 	s.Require().Equal("test@example.com", resp.Email)
// 	s.Require().Equal("John", resp.FirstName)
// 	s.Require().Equal("Doe", resp.LastName)
// 	s.Require().Equal(testTime, resp.BirthDate)
// 	s.Require().Len(resp.Addresses, 1)
// 	s.Require().Equal(provinceID, resp.Addresses[0].ProvinceID)
// 	s.Require().Equal(districtID, resp.Addresses[0].DistrictID)
// 	s.Require().Equal(subDistrictID, resp.Addresses[0].SubDistrictID)
// 	s.Require().Equal(&testAddress, resp.Addresses[0].Address)
// }

// func (s *authRegisterTestSuite) TestRegisterBcryptError() {
// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       time.Now(),
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses:       []request.AddressRequest{},
// 	}

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return([]byte(nil), errors.New("bcrypt error"))

// 	_, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().Error(err)
// 	s.Require().Equal(http.StatusInternalServerError, statusCode)
// 	s.mockTxManager.AssertNotCalled(s.T(), "ExecuteInTransaction")
// }

// func (s *authRegisterTestSuite) TestRegisterUsernameAlreadyExists() {
// 	registerRequest := request.RegisterRequest{
// 		Username:        "existinguser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       time.Now(),
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses:       []request.AddressRequest{},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(errors.New("ERROR: duplicate key value violates unique constraint"))

// 	_, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().Error(err)
// 	s.Require().Equal(http.StatusBadRequest, statusCode)
// 	s.Require().Contains(err.Error(), "username already exists")
// }

// func (s *authRegisterTestSuite) TestRegisterUserCreateError() {
// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       time.Now(),
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses:       []request.AddressRequest{},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(errors.New("database connection error")).Run(func(args mock.Arguments) {
// 		txFunc := args.Get(0).(func(tx *gorm.DB) error)
// 		mockTx := &gorm.DB{}

// 		s.mockUserRepo.On("WithTx", mockTx).Return(s.mockUserRepo)
// 		s.mockUserHistoryRepo.On("WithTx", mockTx).Return(s.mockUserHistoryRepo)
// 		s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(errors.New("database connection error"))

// 		txFunc(mockTx)
// 	})

// 	_, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().Error(err)
// 	s.Require().Equal(http.StatusInternalServerError, statusCode)
// }

// func (s *authRegisterTestSuite) TestRegisterUserHistoryCreateError() {
// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       time.Now(),
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses:       []request.AddressRequest{},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(errors.New("user history creation error")).Run(func(args mock.Arguments) {
// 		txFunc := args.Get(0).(func(tx *gorm.DB) error)
// 		mockTx := &gorm.DB{}

// 		s.mockUserRepo.On("WithTx", mockTx).Return(s.mockUserRepo)
// 		s.mockUserHistoryRepo.On("WithTx", mockTx).Return(s.mockUserHistoryRepo)

// 		s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
// 			user := args.Get(0).(*entity.User)
// 			user.ID = uuid.New()
// 		})
// 		s.mockUserHistoryRepo.On("Create", mock.AnythingOfType("*entity.UserHistory")).Return(errors.New("user history creation error"))

// 		txFunc(mockTx)
// 	})

// 	_, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().Error(err)
// 	s.Require().Equal(http.StatusInternalServerError, statusCode)
// }

// func (s *authRegisterTestSuite) TestRegisterTransactionError() {
// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       time.Now(),
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses:       []request.AddressRequest{},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)
// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(errors.New("transaction failed"))

// 	_, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().Error(err)
// 	s.Require().Equal(http.StatusInternalServerError, statusCode)
// }

// func (s *authRegisterTestSuite) TestRegisterWithMultipleAddresses() {
// 	testTime := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
// 	provinceID1 := uuid.New()
// 	districtID1 := uuid.New()
// 	subDistrictID1 := uuid.New()
// 	testAddress1 := "123 Test Street"

// 	provinceID2 := uuid.New()
// 	districtID2 := uuid.New()
// 	subDistrictID2 := uuid.New()
// 	testAddress2 := "456 Another Street"

// 	registerRequest := request.RegisterRequest{
// 		Username:        "testuser",
// 		Email:           "test@example.com",
// 		FirstName:       "John",
// 		LastName:        "Doe",
// 		BirthDate:       testTime,
// 		Password:        "password123",
// 		ConfirmPassword: "password123",
// 		Addresses: []request.AddressRequest{
// 			{
// 				ProvinceID:    provinceID1,
// 				DistrictID:    districtID1,
// 				SubDistrictID: subDistrictID1,
// 				Address:       &testAddress1,
// 			},
// 			{
// 				ProvinceID:    provinceID2,
// 				DistrictID:    districtID2,
// 				SubDistrictID: subDistrictID2,
// 				Address:       &testAddress2,
// 			},
// 		},
// 	}

// 	hashedPassword := []byte("hashed_password_123")

// 	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

// 	s.mockTxManager.On("ExecuteInTransaction", mock.AnythingOfType("func(*gorm.DB) error")).Return(nil).Run(func(args mock.Arguments) {
// 		txFunc := args.Get(0).(func(tx *gorm.DB) error)
// 		mockTx := &gorm.DB{}

// 		s.mockUserRepo.On("WithTx", mockTx).Return(s.mockUserRepo)
// 		s.mockUserHistoryRepo.On("WithTx", mockTx).Return(s.mockUserHistoryRepo)

// 		s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
// 			user := args.Get(0).(*entity.User)
// 			user.ID = uuid.New()
// 		})
// 		s.mockUserHistoryRepo.On("Create", mock.AnythingOfType("*entity.UserHistory")).Return(nil)

// 		txFunc(mockTx)
// 	})

// 	resp, statusCode, err := s.service.Register(registerRequest)

// 	s.Require().NoError(err)
// 	s.Require().Equal(http.StatusCreated, statusCode)
// 	s.Require().Equal("testuser", resp.Username)
// 	s.Require().Len(resp.Addresses, 2)

// 	s.Require().Equal(provinceID1, resp.Addresses[0].ProvinceID)
// 	s.Require().Equal(districtID1, resp.Addresses[0].DistrictID)
// 	s.Require().Equal(subDistrictID1, resp.Addresses[0].SubDistrictID)
// 	s.Require().Equal(&testAddress1, resp.Addresses[0].Address)

// 	s.Require().Equal(provinceID2, resp.Addresses[1].ProvinceID)
// 	s.Require().Equal(districtID2, resp.Addresses[1].DistrictID)
// 	s.Require().Equal(subDistrictID2, resp.Addresses[1].SubDistrictID)
// 	s.Require().Equal(&testAddress2, resp.Addresses[1].Address)
// }

// func TestAuthRegister(t *testing.T) {
// 	suite.Run(t, new(authRegisterTestSuite))
// }
