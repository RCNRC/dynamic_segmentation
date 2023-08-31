package service

import "github.com/RCNRC/dynamic_segmentation/pkg/repository"

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(title string) error {
	return s.repo.CreateSegment(title)
}

func (s *SegmentService) DeleteSegment(title string) error {
	return s.repo.DeleteSegment(title)
}
