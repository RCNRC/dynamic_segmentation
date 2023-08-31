package repository

import (
	"github.com/jmoiron/sqlx"
)

type User interface {
	AddSegmentRecord(int, int, string, string) error
	GetSegmentId(string) (int, error)
	GetUsersCurrentSegments(int) ([]string, error)
	GetUsersSegmentsHistory(string, string) ([]CsvRaw, error)
}

type Segment interface {
	CreateSegment(string) error
	DeleteSegment(string) error
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment: NewSegmentPostgres(db),
		User:    NewUserPostgres(db),
	}
}
