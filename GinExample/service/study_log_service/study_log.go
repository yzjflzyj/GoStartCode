package study_log_service

import (
	"GoStartCode/GinExample/models"
	"time"
)

type StudyLog struct {
	ID         int
	UserId     int
	DayOfWeek  int
	StudyTime  int
	DateTime   time.Time
	Content    string
	CreatedBy  string
	ModifiedBy string
	State      int
	ModifiedOn time.Time
	CreatedOn  time.Time
	DeletedOn  time.Time

	PageNum  int
	PageSize int
}

func (studyLog *StudyLog) Add() error {
	studyLogModel := models.StudyLog{
		UserId:     studyLog.UserId,
		ModifiedOn: studyLog.ModifiedOn,
		CreatedOn:  studyLog.CreatedOn,
		DeletedOn:  studyLog.DeletedOn,
		DayOfWeek:  studyLog.DayOfWeek,
		StudyTime:  studyLog.StudyTime,
		DateTime:   studyLog.DateTime,
		Content:    studyLog.Content,
		CreatedBy:  studyLog.CreatedBy,
		ModifiedBy: studyLog.ModifiedBy,
		State:      studyLog.State,
	}
	return models.AddStudyLog(&studyLogModel)
}

func (studyLog *StudyLog) QueryStudyLogPage() ([]models.StudyLog, error) {
	return models.QueryStudyLogPage(studyLog.PageNum, studyLog.PageSize)
}

func (studyLog *StudyLog) QueryStudyLogByDateTime(begin time.Time, end time.Time) ([]models.StudyLog, error) {
	return models.QueryStudyLogByDateTime(begin, end)

}
