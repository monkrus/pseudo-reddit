package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	goreddit "github.com/monkrus/pseudo-reddit.git"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

// queries a thread based based on its ID
func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id =$1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
	
}
func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var tt []goreddit.Thread
	if err := s.Select(&tt, `SELECT * FROM threads`); err != nil {
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
