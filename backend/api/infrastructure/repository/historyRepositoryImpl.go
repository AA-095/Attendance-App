package repository

import (
	"fmt"
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

func (s *HistoryRepoImpl) ReadAllHistory(userId int, year int) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("user_id = ?", userId).Where("year = ?", year).Find(&activities).Error
	
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
	fmt.Println(res)
	return res, nil
}

func (s *HistoryRepoImpl) ReadHistoryByDate(userId int, date string) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("user_id = ?", userId).Where("date = ?", date).Find(&activities).Error

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
	fmt.Println(res)
	return res, nil
}
