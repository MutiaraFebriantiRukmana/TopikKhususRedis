package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"redis/domain"
)

type mockUserRepository struct {
	getByIDCalls   int
	saveCalls      int
	deleteCalls    int
	getByIDUser    *domain.User
	getByIDError   error
	saveError      error
	deleteError    error
}

func (m *mockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	m.getByIDCalls++
	return m.getByIDUser, m.getByIDError
}

func (m *mockUserRepository) Save(ctx context.Context, user *domain.User) error {
	m.saveCalls++
	return m.saveError
}

func (m *mockUserRepository) Delete(ctx context.Context, id string) error {
	m.deleteCalls++
	return m.deleteError
}

func TestUserUsecase_GetUser(t *testing.T) {
	mockRepo := &mockUserRepository{
		getByIDUser: &domain.User{ID: "1", Name: "John", Email: "john@example.com"},
	}
	usecase := NewUserUsecase(mockRepo)

	user, err := usecase.GetUser(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, 1, mockRepo.getByIDCalls)
	assert.Equal(t, "John", user.Name)
}

func TestUserUsecase_GetUser_InvalidID(t *testing.T) {
	mockRepo := &mockUserRepository{}
	usecase := NewUserUsecase(mockRepo)

	_, err := usecase.GetUser(context.Background(), "")

	assert.Error(t, err)
	assert.Equal(t, 0, mockRepo.getByIDCalls)
}

func TestUserUsecase_CreateUser(t *testing.T) {
	mockRepo := &mockUserRepository{}
	usecase := NewUserUsecase(mockRepo)
	user := &domain.User{ID: "1", Name: "John", Email: "john@example.com"}

	err := usecase.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, 1, mockRepo.saveCalls)
}

func TestUserUsecase_CreateUser_InvalidData(t *testing.T) {
	mockRepo := &mockUserRepository{}
	usecase := NewUserUsecase(mockRepo)
	user := &domain.User{ID: "", Name: "", Email: ""}

	err := usecase.CreateUser(context.Background(), user)

	assert.Error(t, err)
	assert.Equal(t, 0, mockRepo.saveCalls)
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	mockRepo := &mockUserRepository{}
	usecase := NewUserUsecase(mockRepo)

	err := usecase.DeleteUser(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, 1, mockRepo.deleteCalls)
}

func TestUserUsecase_DeleteUser_Error(t *testing.T) {
	expectedErr := errors.New("repo error")
	mockRepo := &mockUserRepository{deleteError: expectedErr}
	usecase := NewUserUsecase(mockRepo)

	err := usecase.DeleteUser(context.Background(), "1")

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, 1, mockRepo.deleteCalls)
}

