package auth_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	jwtPkg "github.com/wisaitas/share-pkg/auth/jwt"
	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	bcryptPkg "github.com/wisaitas/share-pkg/crypto/bcrypt"
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/service/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type authTestSuite struct {
	suite.Suite
	mockUserRepo        *repository.MockBaseRepository[entity.User]
	mockUserHistoryRepo *repository.MockBaseRepository[entity.UserHistory]
	mockRedis           *redisPkg.MockRedis
	mockBcrypt          *bcryptPkg.MockBcrypt
	mockJWT             *jwtPkg.MockJwt
	service             auth.AuthService
	mockDB              *gorm.DB
}

func (s *authTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	s.Require().NoError(err)
	s.mockDB = db

	s.mockUserRepo = new(repository.MockBaseRepository[entity.User])
	s.mockUserHistoryRepo = new(repository.MockBaseRepository[entity.UserHistory])
	s.mockRedis = redisPkg.NewMockRedis()
	s.mockBcrypt = bcryptPkg.NewMockBcrypt()
	s.mockJWT = jwtPkg.NewMockJwt()

	s.service = auth.NewAuthService(
		s.mockUserRepo,
		s.mockUserHistoryRepo,
		s.mockRedis,
		s.mockBcrypt,
		s.mockJWT,
	)
}

func (s *authTestSuite) TestRegisterSuccess() {
	testTime := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	provinceID := uuid.New()
	districtID := uuid.New()
	subDistrictID := uuid.New()
	testAddress := "123 Test Street"

	registerRequest := request.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		FirstName:       "John",
		LastName:        "Doe",
		BirthDate:       testTime,
		Password:        "password123",
		ConfirmPassword: "password123",
		Addresses: []request.AddressRequest{
			{
				ProvinceID:    provinceID,
				DistrictID:    districtID,
				SubDistrictID: subDistrictID,
				Address:       &testAddress,
			},
		},
	}

	hashedPassword := []byte("hashed_password_123")

	s.mockBcrypt.On("GenerateFromPassword", "password123", bcrypt.DefaultCost).Return(hashedPassword, nil)

	s.mockUserRepo.On("GetDB").Return(s.mockDB)

	mockTxUserRepo := new(repository.MockBaseRepository[entity.User])
	mockTxUserHistoryRepo := new(repository.MockBaseRepository[entity.UserHistory])

	s.mockUserRepo.On("WithTx", mock.AnythingOfType("*gorm.DB")).Return(mockTxUserRepo)
	s.mockUserHistoryRepo.On("WithTx", mock.AnythingOfType("*gorm.DB")).Return(mockTxUserHistoryRepo)

	mockTxUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*entity.User)
		user.Id = uuid.New()
	})
	mockTxUserHistoryRepo.On("Create", mock.AnythingOfType("*entity.UserHistory")).Return(nil)

	resp, statusCode, err := s.service.Register(registerRequest)

	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, statusCode)
	s.Require().Equal("testuser", resp.Username)
	s.Require().Equal("test@example.com", resp.Email)
	s.Require().Equal("John", resp.FirstName)
	s.Require().Equal("Doe", resp.LastName)
	s.Require().Equal(testTime, resp.BirthDate)
	s.Require().Len(resp.Addresses, 1)
	s.Require().Equal(provinceID, resp.Addresses[0].ProvinceID)
	s.Require().Equal(districtID, resp.Addresses[0].DistrictID)
	s.Require().Equal(subDistrictID, resp.Addresses[0].SubDistrictID)
	s.Require().Equal(&testAddress, resp.Addresses[0].Address)

	s.mockBcrypt.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockUserHistoryRepo.AssertExpectations(s.T())
	mockTxUserRepo.AssertExpectations(s.T())
	mockTxUserHistoryRepo.AssertExpectations(s.T())
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(authTestSuite))
}
