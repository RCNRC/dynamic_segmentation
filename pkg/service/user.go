package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	dynamicsegmentation "github.com/RCNRC/dynamic_segmentation"
	"github.com/RCNRC/dynamic_segmentation/pkg/repository"
)

const reportsFilePath = "reports"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Update(update dynamicsegmentation.UserUpdate) error {
	var err error
	var segmentId int
	for _, segment := range update.SegmentsAdd {
		if segmentId, err = s.repo.GetSegmentId(segment.Service); err != nil {
			return err
		}
		if err = s.repo.AddSegmentRecord(update.UserId, segmentId, segment.TTL, "i"); err != nil {
			return err
		}
	}
	for _, segment := range update.SegmentsDelete {
		if segmentId, err = s.repo.GetSegmentId(segment.Service); err != nil {
			return err
		}
		if err = s.repo.AddSegmentRecord(update.UserId, segmentId, segment.TTL, "d"); err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) GetUsersCurrentSegments(userId int) ([]string, error) {
	return s.repo.GetUsersCurrentSegments(userId)
}

func (s *UserService) GetUsersSegmentsHistory(dateFrom string, dateTo string) (string, error) {
	fileName := fmt.Sprintf("%s:%s(%s).csv", dateFrom, dateTo, time.Now().Format("2006-01-02 15:04:05"))
	if err := os.MkdirAll(reportsFilePath, 0770); err != nil {
		return "", err
	}
	file, err := os.Create(reportsFilePath + "/" + fileName)
	if err != nil {
		return "", err
	}
	csv := csv.NewWriter(file)
	defer csv.Flush()

	csvRaws, err := s.repo.GetUsersSegmentsHistory(dateFrom, dateTo)
	if err != nil {
		return "", err
	}
	for _, csvRaw := range csvRaws {
		err = csv.Write([]string{strconv.Itoa(csvRaw.UsersId), csvRaw.Segments, csvRaw.Operations, csvRaw.DateTimes})
		if err != nil {
			return "", err
		}
	}
	return fileName, nil
}

func (s *UserService) GetReportsPath() string {
	return reportsFilePath
}
