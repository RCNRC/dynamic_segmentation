package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(title string) error {
	query := fmt.Sprintf("INSERT INTO %s (title) values ($1)", segmentsTable)
	if err := r.db.QueryRow(query, title); err != nil {
		return err.Err()
	}
	return nil
}

func (r *SegmentPostgres) DeleteSegment(title string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE title = $1", segmentsTable)
	if err := r.db.QueryRow(query, title); err != nil {
		return err.Err()
	}
	return nil
}
