package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	goreddit "github.com/monkrus/pseudo-reddit.git"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

// queries a thread based based on its ID

func (s *PostStore) Post (id uuid.UUID) (goreddit.Post, error) {
	var t goreddit.Post
	if err := s.Get(&t, `SELECT * FROM posts WHERE id =$1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return t, nil
}
func (s *PostStore) PostsByThreads(threadID uuid.UUID) ([]goreddit.Post, error) {
	var tt []goreddit.Post
	if err := s.Select(&pp, `SELECT * FROM posts`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return tt, nil
}
func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `UPDATE threads VALUES($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}
func (s *ThreadStore) DeleteThread(t *goreddit.Thread) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1, id`); err != nil {
		return fmt.Errorf("error deleting thread: %w, err")
	}
	return nil
}
