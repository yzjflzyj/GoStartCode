package v1

import (
	app "GoStartCode/GinExample/pkg/app"
	"GoStartCode/GinExample/pkg/e"
	"GoStartCode/GinExample/service/study_log_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AddStudyLogForm struct {
	UserId    int    `form:"User_Id"`
	StudyTime int    `form:"Study_Time"`
	Content   string `form:"content"`
	DateTime  string `form:"Date_Time"`
}

func AddStudyLog(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddStudyLogForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	// 获取日期
	dateTime := func(dateString string) time.Time {
		if dateString == "" {
			return time.Now()
		}
		layout := "2006-01-02" // 指定日期的格式
		dateTime, _ := time.Parse(layout, dateString)
		return dateTime
	}(form.DateTime)

	// 增加学习记录
	var studyLog = study_log_service.StudyLog{
		UserId:     form.UserId,
		DayOfWeek:  int(dateTime.Weekday()),
		StudyTime:  form.StudyTime,
		DateTime:   dateTime,
		Content:    form.Content,
		CreatedBy:  "时海徜徉",
		ModifiedBy: "时海徜徉",
		State:      1,
		ModifiedOn: dateTime,
		CreatedOn:  dateTime,
		DeletedOn:  dateTime,
	}
	err := studyLog.Add()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func QueryStudyLogPage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form study_log_service.StudyLog
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	studyLogs, _ := form.QueryStudyLogPage()
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": studyLogs,
		"total": 100,
	})
}
