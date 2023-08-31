package service

import (
	dynamicsegmentation "github.com/RCNRC/dynamic_segmentation"
	"github.com/RCNRC/dynamic_segmentation/pkg/repository"
)

type User interface {
	Update(update dynamicsegmentation.UserUpdate) error
	GetUsersCurrentSegments(userId int) ([]string, error)
	GetUsersSegmentsHistory(string, string) (string, error)
	GetReportsPath() string
}

type Segment interface {
	CreateSegment(title string) error
	DeleteSegment(title string) error
}

type Service struct {
	User
	Segment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Segment: NewSegmentService(repo.Segment),
		User:    NewUserService(repo.User),
	}
}
