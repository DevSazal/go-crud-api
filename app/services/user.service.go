package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/DevSazal/go-crud-api/app/models"
	"github.com/DevSazal/go-crud-api/app/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepository.GetUsers()
}

func (s *UserService) GetUser(id primitive.ObjectID) (*models.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(id primitive.ObjectID, user *models.User) error {
	return s.userRepository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id primitive.ObjectID) error {
	return s.userRepository.DeleteUser(id)
}