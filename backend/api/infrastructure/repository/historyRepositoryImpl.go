package repository

import (
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	orm_model "work-management-app/infrastructure/orm"

	"gorm.io/gorm"
)

type HistoryRepoImpl struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return &HistoryRepoImpl{
		db: db,
	}
}

func (s *HistoryRepoImpl) ReadAllHistory(useId int, year int) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("userId = ?", useId).Where("year = ?", year).Find(&activities).Error
	
	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		resItem := model.Attendance{
			Id:             activity.Id,
			UserId:         activity.UserId,
			AttendanceType: model.IntToActionEnum(activity.AttendanceType),
			Time:           activity.Time,
			Date:           activity.Date,
		}
		res = append(res, resItem)
	}
	return res, nil
}

func (s *HistoryRepoImpl) ReadHistoryByDate(useId int, date string) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("userId = ?", useId).Where("date = ?", date).Find(&activities).Error

	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		resItem := model.Attendance{
			Id:             activity.Id,
			UserId:         activity.UserId,
			AttendanceType: model.IntToActionEnum(activity.AttendanceType),
			Time:           activity.Time,
			Date:           activity.Date,
		}
		res = append(res, resItem)
	}
	return res, nil
}
