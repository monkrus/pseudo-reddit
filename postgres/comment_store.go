package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	goreddit "github.com/monkrus/pseudo-reddit.git"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (goreddit.Comment, error) {
	var t goreddit.Comment
	if err := s.Get(&c, `SELECT * FROM threads WHERE id =$1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return c, nil
}
func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]goreddit.Comment, error) {
	var cc []goreddit.Comment
	if err := s.Select(&cc, `SELECT * FROM post_id =$1`, postID); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return cc, nil
}
func (s *CommentStore) CreateComment(c *goreddit.Thread) error {
	if err := s.Get(c, `INSERT INTO threads VALUES($1, $2, $3) RETURNING *`,
		c.ID,
		c.PostID,
		c.Content,
		c.Votes); err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(t *goreddit.Comment) error {
	if err := s.Get(t, `UPDATE comments SET post_id = $1, content = $2, votes = $3 WHERE id = $4 RETURNING *`,
		c.PostID,
		c.Content,
		c.Votes,
		c.ID); err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}
func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1, id`); err != nil {
		return fmt.Errorf("error deleting comment: %w, err")
	}
	return nil
}
