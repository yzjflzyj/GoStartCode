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

	// 增加学习记录
	var studyLog = study_log_service.StudyLog{
		UserId:     form.UserId,
		DayOfWeek:  int(time.Now().Weekday()),
		StudyTime:  form.StudyTime,
		DateTime:   time.Now(),
		Content:    form.Content,
		CreatedBy:  "时海徜徉",
		ModifiedBy: "时海徜徉",
		State:      1,
		ModifiedOn: time.Now(),
		CreatedOn:  time.Now(),
		DeletedOn:  time.Now(),
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
