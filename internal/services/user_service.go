package services

import (
	"database/sql"
	"encoding/base64"
	"web/example/internal/domain"
	handlermodel "web/example/internal/http/handler_model"
	"web/example/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService handles all business logic for users
type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

// NewUserService creates a new instance of UserService with repository
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(db),
	}
}

func (s *UserService) CreateUserService(req *handlermodel.CreateUserRequest) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	if err != nil {
		return err
	}

	pwh := base64.StdEncoding.EncodeToString(bytes)

	user := &domain.User{
		Email:         req.Email,
		Username:      req.Username,
		Password_hash: pwh,
	}

	return s.UserRepo.CreateUser(user)
}

func (s *UserService) LoginUserService(req *handlermodel.LoginUserRequest) (string, error) {

	usr, err := s.UserRepo.ReadUser(req.Email)

	if err != nil {
		return "", err
	}

	if decoded, err := base64.StdEncoding.DecodeString(usr.Password_hash); err != nil {
		return "", err
	} else if err := bcrypt.CompareHashAndPassword([]byte(decoded), []byte(req.Password)); err != nil {
		return "", err
	}

	// For simplicity sake, we use a fixed string intead of a secure token
	return "MOCK_VALID_JWT", nil

}

func (s *UserService) DeleteUser(email string) error {
	return s.UserRepo.DeleteUser(email)
}

func (s *UserService) ReadUser(email string) (*domain.User, error) {
	return s.UserRepo.ReadUser(email)
}

func (s *UserService) UpdateUsername(email string, newUsername string) error {
	return s.UserRepo.UpdateUsername(email, newUsername)
}
