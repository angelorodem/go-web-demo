package services

import (
	"database/sql"
	"encoding/base64"
	"web/example/internal/domain"
	handlermodel "web/example/internal/http/handler_model"
	"web/example/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(req *handlermodel.CreateUserRequest, db *sql.DB) error {
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

	err = repository.CreateUser(db, user)

	if err != nil {
		return err
	}

	return nil
}

func LoginUserService(req *handlermodel.LoginUserRequest, db *sql.DB) (string, error) {

	usr, err := repository.GetUser(db, req.Email)

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
