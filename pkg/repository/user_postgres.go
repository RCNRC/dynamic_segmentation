package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetSegmentId(title string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE title = $1", segmentsTable)
	raw := r.db.QueryRow(query, title)
	if raw.Err() != nil {
		return 0, raw.Err()
	}
	if err := raw.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) AddSegmentRecord(userId int, segmentId int, ttl string, actionType string) error {
	if ttl == "" {
		query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, action_type) values ($1, $2, $3)", usersSegmentsTable)
		if raw := r.db.QueryRow(query, userId, segmentId, actionType); raw.Err() != nil {
			return raw.Err()
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, ttl, action_type) values ($1, $2, $3, $4)", usersSegmentsTable)
		if raw := r.db.QueryRow(query, userId, segmentId, ttl, actionType); raw.Err() != nil {
			return raw.Err()
		}
	}
	return nil
}

func (r *UserPostgres) GetUsersCurrentSegments(userId int) ([]string, error) {
	var segments []string
	pre_pre_query := fmt.Sprintf("SELECT segment_id, MAX(created_time) as created_time FROM %s WHERE user_id = $1 GROUP BY segment_id", usersSegmentsTable)
	pre_query := fmt.Sprintf("SELECT segments.segment_id, segments.ttl, segments.action_type FROM (%s) as target_segments, %s as segments WHERE target_segments.segment_id = segments.segment_id AND target_segments.created_time = segments.created_time", pre_pre_query, usersSegmentsTable)
	query := fmt.Sprintf("SELECT segments.title FROM (%s) as users_segments, %s as segments WHERE users_segments.segment_id = segments.id AND users_segments.action_type = 'i' AND users_segments.ttl > now()", pre_query, segmentsTable)
	if err := r.db.Select(&segments, query, userId); err != nil {
		return []string{}, err
	}
	return segments, nil
}

type CsvRaw struct {
	UsersId    int    `db:"user_id"`
	Segments   string `db:"title"`
	Operations string `db:"action_type"`
	DateTimes  string `db:"created_time"`
}

func (r *UserPostgres) GetUsersSegmentsHistory(fromDate string, toDate string) ([]CsvRaw, error) {
	csvRaws := []CsvRaw{}
	query := fmt.Sprintf("SELECT user_segments.user_id, segments.title, CASE WHEN user_segments.action_type = 'i' THEN '%s' ELSE '%s' END AS action_type, user_segments.created_time FROM %s as user_segments, %s as segments WHERE user_segments.created_time >= TO_TIMESTAMP($1, 'YYYY-MM') AND user_segments.created_time <= TO_TIMESTAMP($2, 'YYYY-MM') AND segments.id = user_segments.segment_id", insertion, deletion, usersSegmentsTable, segmentsTable)
	if err := r.db.Select(&csvRaws, query, fromDate, toDate); err != nil {
		return []CsvRaw{}, err
	}
	return csvRaws, nil
}
