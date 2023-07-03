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

	/**
	SELECT t.id,t.day_of_week, date(t.date_time),t.content, s.study_time
	FROM blog_study_log AS t
	JOIN (
	  SELECT sum(study_time) as study_time, MIN(id) as min_id
	  FROM blog_study_log
	  GROUP BY date(date_time)
	) AS s ON t.id = s.min_id
	order by date(t.date_time) asc;
	*/

	/*//GORM v2.0.0才支持设置别名，当前版本1.25.2
	subQuery := db.Model(&StudyLog{}).
		Select("SUM(study_time) AS study_time, MIN(id) AS min_id").
		Group("DATE(date_time)").SubQuery()
	db.Model(&StudyLog{}).
		Alias("t"). //GORM v2.0.0才支持设置别名
		Select("t.id, t.day_of_week, DATE(t.date_time), t.content, s.study_time").
		Joins("JOIN (?) AS s ON t.id = s.min_id", subQuery).
	    Order("DATE(t.date_time) ASC").
		Find(&studyLogList)
	*/

	// 原生sql的方式
	db.Table("blog_study_log AS t").
		Joins("JOIN (SELECT SUM(study_time) AS study_time, MIN(id) AS min_id FROM blog_study_log GROUP BY DATE(date_time)) AS s ON t.id = s.min_id").
		Select("t.id, t.day_of_week, t.date_time, t.content, s.study_time").
		Order("DATE(t.date_time) ASC").
		Scan(&studyLogList)

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
