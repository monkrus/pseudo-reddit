package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	goreddit "github.com/monkrus/pseudo-reddit.git"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, error := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database")
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", error)
	}

	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil

}

type Store struct {
	goreddit.ThreadStore
	goreddit.PostStore
	goreddit.CommentStore
}
