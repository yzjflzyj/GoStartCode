package models

import (
	"gorm.io/gorm"
	"time"
)

type StudyLog struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	DeletedOn  time.Time `json:"deleted_on"`

	UserId     int       `json:"user_id"`
	DayOfWeek  int       `json:"day_of_week"`
	StudyTime  int       `json:"study_time"`
	DateTime   time.Time `json:"date_time"`
	Content    string    `json:"content"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	State      int       `json:"state"`
}

func QueryStudyLogPage(pageNum int, pageSize int) ([]StudyLog, error) {
	var (
		studyLogList []StudyLog
		err          error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Find(&studyLogList).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Find(&studyLogList).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return studyLogList, nil
}

func QueryStudyLogByDateTime(begin time.Time, end time.Time) ([]StudyLog, error) {
	var (
		studyLogList []StudyLog
		err          error
	)

	err = db.Where("date_time between ? and ?", begin, end).Find(&studyLogList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return studyLogList, nil
}

func AddStudyLog(studyLog *StudyLog) error {
	if err := db.Create(&studyLog).Error; err != nil {
		return err
	}
	return nil
}
