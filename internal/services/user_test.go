package services_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/mocks"
	"github.com/wisaitas/standard-golang/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		input          request.CreateUserRequest
		mockError      error
		expectedStatus int
		expectedError  error
	}{
		{
			name: "Success",
			input: request.CreateUserRequest{
				Username:        "testuser",
				Email:           "test@example.com",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedError:  nil,
		},
		{
			name: "Username Already Exists",
			input: request.CreateUserRequest{
				Username:        "existinguser",
				Email:           "test@example.com",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockError:      errors.New("ERROR: duplicate key value violates unique constraint"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  errors.New("username already exists"),
		},
		{
			name: "Internal Server Error",
			input: request.CreateUserRequest{
				Username:        "testuser",
				Email:           "test@example.com",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			mockRepo := new(mocks.MockUserRepository)

			// Set up mock expectations
			mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(tt.mockError)

			// Create service with mock repository
			service := services.NewUserService(mockRepo)

			// Call the service method
			resp, status, err := service.CreateUser(tt.input)

			// Assertions
			assert.Equal(t, tt.expectedStatus, status)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp.ID)
				assert.Equal(t, tt.input.Username, resp.Username)
				assert.Equal(t, tt.input.Email, resp.Email)
			}

			// Verify that mock expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}
