package services

import (
	"database/sql"
	"fmt"
	"web/example/internal/domain"
	handlermodel "web/example/internal/http/handler_model"
	"web/example/internal/repository"
)

// PostService handles all business logic for posts
type PostService struct {
	PostRepo repository.PostRepositoryInterface
	UserRepo repository.UserRepositoryInterface
}

// NewPostService creates a new instance of PostService with repositories
func NewPostService(db *sql.DB) *PostService {
	return &PostService{
		PostRepo: repository.NewPostRepository(db),
		UserRepo: repository.NewUserRepository(db),
	}
}

func (s *PostService) verifyUserOwnership(postId int, userEmail string) (*domain.Post, error) {
	// the following is a mock check, since we do not have claims on the token
	// we get the id from the user based on it's email
	// the correct approach would be to get the id from the validated claims from the JWT/Paseto token
	user, err := s.UserRepo.ReadUser(userEmail)
	if err != nil {
		return nil, err
	}

	post, err := s.PostRepo.ReadPost(postId)

	if err != nil {
		return nil, err
	}

	if user.Id != post.UserId {
		return nil, fmt.Errorf("user does not own this post")
	}

	return post, nil
}

func (s *PostService) CreatePostService(req *handlermodel.CreatePostRequest) error {

	// the following is a mock check, since we do not have claims on the token
	// we get the id from the user based on it's email
	// the correct approach would be to get the id from the validated claims from the JWT/Paseto token
	user, err := s.UserRepo.ReadUser(req.UserEmail)
	if err != nil {
		return err
	}

	return s.PostRepo.CreatePost(&domain.Post{UserId: user.Id, Title: req.Title, Content: req.Content})
}

func (s *PostService) UpdatePostService(req *handlermodel.UpdatePostRequest) error {
	post, err := s.verifyUserOwnership(req.Id, req.UserEmail)

	if err != nil {
		return err
	}

	return s.PostRepo.UpdatePost(post.Id, req.NewTitle, req.NewContent)
}

func (s *PostService) DeletePostService(req *handlermodel.DeletePostRequest) error {
	post, err := s.verifyUserOwnership(req.Id, req.UserEmail)

	if err != nil {
		return err
	}

	return s.PostRepo.DeletePost(post.Id)
}

func (s *PostService) ReadPost(id int) (*domain.Post, error) {
	return s.PostRepo.ReadPost(id)
}

func (s *PostService) ReadAllPosts() ([]domain.Post, error) {
	return s.PostRepo.ReadAllPosts()
}
