package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"web/example/internal/domain"
)

type PostRepositoryInterface interface {
	CreatePost(post *domain.Post) error
	ReadPost(id int) (*domain.Post, error)
	UpdatePost(id int, title string, content string) error
	DeletePost(id int) error
	ReadAllPosts() ([]domain.Post, error)
}

// PostRepository handles all database operations for posts
type PostRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new instance of PostRepository
func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(post *domain.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (user_id, title, content) values (?,?,?)", post.UserId, post.Title, post.Content)

	return err
}

func (r *PostRepository) ReadPost(id int) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx,
		"SELECT * FROM posts WHERE id == ?", id)

	var p domain.Post

	if err := row.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.CreatedAt); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostRepository) UpdatePost(id int, title string, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := r.db.ExecContext(ctx, "UPDATE posts SET title = ?, content = ? WHERE id == ?", title, content, id)

	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n <= 0 {
		return fmt.Errorf("nothing was updated")
	}

	return nil
}

func (r *PostRepository) DeletePost(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id == ?", id)

	return err
}

func (r *PostRepository) ReadAllPosts() ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx,
		"SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		var p domain.Post

		if err := rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
