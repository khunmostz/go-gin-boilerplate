package tests

import (
	"context"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/service"
	"go-gin-boilerplate/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFooService_CreateFoo(t *testing.T) {
	tests := []struct {
		name           string
		inputFoo       *domain.Foo
		mockSetup      func(*mocks.MockFooRepository)
		expectedResult *domain.Foo
		expectedError  string
	}{
		{
			name: "successful creation",
			inputFoo: &domain.Foo{
				Name: "Test Foo",
			},
			mockSetup: func(mockRepo *mocks.MockFooRepository) {
				expectedFoo := &domain.Foo{
					ID:   "507f1f77bcf86cd799439011",
					Name: "Test Foo",
				}
				mockRepo.On("CreateFoo", mock.Anything, mock.AnythingOfType("*domain.Foo")).Return(expectedFoo, nil)
			},
			expectedResult: &domain.Foo{
				ID:   "507f1f77bcf86cd799439011",
				Name: "Test Foo",
			},
			expectedError: "",
		},
		{
			name: "repository error",
			inputFoo: &domain.Foo{
				Name: "Test Foo",
			},
			mockSetup: func(mockRepo *mocks.MockFooRepository) {
				mockRepo.On("CreateFoo", mock.Anything, mock.AnythingOfType("*domain.Foo")).Return(nil, assert.AnError)
			},
			expectedResult: nil,
			expectedError:  assert.AnError.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(mocks.MockFooRepository)
			tt.mockSetup(mockRepo)

			fooService := service.NewFooService(mockRepo)

			// Execute
			result, err := fooService.CreateFoo(context.Background(), tt.inputFoo)

			// Assert
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Name, result.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
