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
)

type authTestSuite struct {
	suite.Suite
	mockUserRepo *repository.MockBaseRepository[entity.User]
	mockRedis    *redisPkg.MockRedis
	mockBcrypt   *bcryptPkg.MockBcrypt
	mockJWT      *jwtPkg.MockJwt
	service      auth.AuthService
}

func (s *authTestSuite) SetupTest() {
	s.mockUserRepo = new(repository.MockBaseRepository[entity.User])
	s.mockRedis = redisPkg.NewMockRedis()
	s.mockBcrypt = bcryptPkg.NewMockBcrypt()
	s.mockJWT = jwtPkg.NewMockJwt()

	s.service = auth.NewAuthService(
		s.mockUserRepo,
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

	s.mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*entity.User)
		user.Id = uuid.New()
	})

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
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(authTestSuite))
}
