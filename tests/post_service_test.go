package tests

import (
	"errors"
	"testing"
	"web/example/internal/domain"
	handlermodel "web/example/internal/http/handler_model"
	"web/example/internal/repository/mocks"
	"web/example/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestPostService_CreatePostService(t *testing.T) {
	tests := []struct {
		name       string
		request    *handlermodel.CreatePostRequest
		setupMocks func(*mocks.MockPostRepositoryInterface, *mocks.MockUserRepositoryInterface)
		wantErr    bool
		errMsg     string
	}{
		{
			name: "success",
			request: &handlermodel.CreatePostRequest{
				UserEmail: "test@example.com",
				Title:     "Test Title",
				Content:   "Test Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("test@example.com").Return(&domain.User{
					Id:       1,
					Email:    "test@example.com",
					Username: "testuser",
				}, nil)

				PostRepo.EXPECT().CreatePost(&domain.Post{
					UserId:  1,
					Title:   "Test Title",
					Content: "Test Content",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			request: &handlermodel.CreatePostRequest{
				UserEmail: "notfound@example.com",
				Title:     "Test Title",
				Content:   "Test Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("notfound@example.com").Return(nil, errors.New("user not found"))
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name: "create post fails",
			request: &handlermodel.CreatePostRequest{
				UserEmail: "test@example.com",
				Title:     "Test Title",
				Content:   "Test Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("test@example.com").Return(&domain.User{
					Id:       1,
					Email:    "test@example.com",
					Username: "testuser",
				}, nil)

				PostRepo.EXPECT().CreatePost(&domain.Post{
					UserId:  1,
					Title:   "Test Title",
					Content: "Test Content",
				}).Return(errors.New("database error"))
			},
			wantErr: true,
			errMsg:  "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
			mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

			tt.setupMocks(mockPostRepo, mockUserRepo)

			service := &services.PostService{
				PostRepo: mockPostRepo,
				UserRepo: mockUserRepo,
			}

			err := service.CreatePostService(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPostService_UpdatePostService(t *testing.T) {
	tests := []struct {
		name       string
		request    *handlermodel.UpdatePostRequest
		setupMocks func(*mocks.MockPostRepositoryInterface, *mocks.MockUserRepositoryInterface)
		wantErr    bool
		errMsg     string
	}{
		{
			name: "success",
			request: &handlermodel.UpdatePostRequest{
				Id:         1,
				UserEmail:  "test@example.com",
				NewTitle:   "Updated Title",
				NewContent: "Updated Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("test@example.com").Return(&domain.User{
					Id:       1,
					Email:    "test@example.com",
					Username: "testuser",
				}, nil)

				PostRepo.EXPECT().ReadPost(1).Return(&domain.Post{
					Id:      1,
					UserId:  1,
					Title:   "Old Title",
					Content: "Old Content",
				}, nil)

				PostRepo.EXPECT().UpdatePost(1, "Updated Title", "Updated Content").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user does not own post",
			request: &handlermodel.UpdatePostRequest{
				Id:         1,
				UserEmail:  "other@example.com",
				NewTitle:   "Updated Title",
				NewContent: "Updated Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("other@example.com").Return(&domain.User{
					Id:       2,
					Email:    "other@example.com",
					Username: "otheruser",
				}, nil)

				PostRepo.EXPECT().ReadPost(1).Return(&domain.Post{
					Id:      1,
					UserId:  1,
					Title:   "Title",
					Content: "Content",
				}, nil)
			},
			wantErr: true,
			errMsg:  "user does not own this post",
		},
		{
			name: "post not found",
			request: &handlermodel.UpdatePostRequest{
				Id:         999,
				UserEmail:  "test@example.com",
				NewTitle:   "Updated Title",
				NewContent: "Updated Content",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("test@example.com").Return(&domain.User{
					Id:       1,
					Email:    "test@example.com",
					Username: "testuser",
				}, nil)

				PostRepo.EXPECT().ReadPost(999).Return(nil, errors.New("post not found"))
			},
			wantErr: true,
			errMsg:  "post not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
			mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

			tt.setupMocks(mockPostRepo, mockUserRepo)

			service := &services.PostService{
				PostRepo: mockPostRepo,
				UserRepo: mockUserRepo,
			}

			err := service.UpdatePostService(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPostService_DeletePostService(t *testing.T) {
	tests := []struct {
		name       string
		request    *handlermodel.DeletePostRequest
		setupMocks func(*mocks.MockPostRepositoryInterface, *mocks.MockUserRepositoryInterface)
		wantErr    bool
		errMsg     string
	}{
		{
			name: "success",
			request: &handlermodel.DeletePostRequest{
				Id:        1,
				UserEmail: "test@example.com",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("test@example.com").Return(&domain.User{
					Id:       1,
					Email:    "test@example.com",
					Username: "testuser",
				}, nil)

				PostRepo.EXPECT().ReadPost(1).Return(&domain.Post{
					Id:      1,
					UserId:  1,
					Title:   "Title",
					Content: "Content",
				}, nil)

				PostRepo.EXPECT().DeletePost(1).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user does not own post",
			request: &handlermodel.DeletePostRequest{
				Id:        1,
				UserEmail: "other@example.com",
			},
			setupMocks: func(PostRepo *mocks.MockPostRepositoryInterface, UserRepo *mocks.MockUserRepositoryInterface) {
				UserRepo.EXPECT().ReadUser("other@example.com").Return(&domain.User{
					Id:       2,
					Email:    "other@example.com",
					Username: "otheruser",
				}, nil)

				PostRepo.EXPECT().ReadPost(1).Return(&domain.Post{
					Id:      1,
					UserId:  1,
					Title:   "Title",
					Content: "Content",
				}, nil)
			},
			wantErr: true,
			errMsg:  "user does not own this post",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
			mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

			tt.setupMocks(mockPostRepo, mockUserRepo)

			service := &services.PostService{
				PostRepo: mockPostRepo,
				UserRepo: mockUserRepo,
			}

			err := service.DeletePostService(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPostService_ReadPost(t *testing.T) {
	expectedPost := &domain.Post{
		Id:      1,
		UserId:  1,
		Title:   "Test Title",
		Content: "Test Content",
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
		mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

		mockPostRepo.EXPECT().ReadPost(1).Return(expectedPost, nil)

		service := &services.PostService{
			PostRepo: mockPostRepo,
			UserRepo: mockUserRepo,
		}

		post, err := service.ReadPost(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedPost, post)
	})

	t.Run("post not found", func(t *testing.T) {
		mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
		mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

		mockPostRepo.EXPECT().ReadPost(999).Return(nil, errors.New("post not found"))

		service := &services.PostService{
			PostRepo: mockPostRepo,
			UserRepo: mockUserRepo,
		}

		post, err := service.ReadPost(999)

		assert.Error(t, err)
		assert.Nil(t, post)
		assert.Contains(t, err.Error(), "post not found")
	})
}

func TestPostService_ReadAllPosts(t *testing.T) {
	expectedPosts := []domain.Post{
		{
			Id:      1,
			UserId:  1,
			Title:   "First Post",
			Content: "Content 1",
		},
		{
			Id:      2,
			UserId:  2,
			Title:   "Second Post",
			Content: "Content 2",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
		mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

		mockPostRepo.EXPECT().ReadAllPosts().Return(expectedPosts, nil)

		service := &services.PostService{
			PostRepo: mockPostRepo,
			UserRepo: mockUserRepo,
		}

		posts, err := service.ReadAllPosts()

		assert.NoError(t, err)
		assert.Equal(t, expectedPosts, posts)
	})

	t.Run("empty result", func(t *testing.T) {
		mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
		mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

		mockPostRepo.EXPECT().ReadAllPosts().Return([]domain.Post{}, nil)

		service := &services.PostService{
			PostRepo: mockPostRepo,
			UserRepo: mockUserRepo,
		}

		posts, err := service.ReadAllPosts()

		assert.NoError(t, err)
		assert.Empty(t, posts)
	})

	t.Run("database error", func(t *testing.T) {
		mockPostRepo := mocks.NewMockPostRepositoryInterface(t)
		mockUserRepo := mocks.NewMockUserRepositoryInterface(t)

		mockPostRepo.EXPECT().ReadAllPosts().Return(nil, errors.New("database connection failed"))

		service := &services.PostService{
			PostRepo: mockPostRepo,
			UserRepo: mockUserRepo,
		}

		posts, err := service.ReadAllPosts()

		assert.Error(t, err)
		assert.Nil(t, posts)
		assert.Contains(t, err.Error(), "database connection failed")
	})
}
