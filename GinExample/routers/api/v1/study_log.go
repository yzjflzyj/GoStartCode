package v1

import (
	app "GoStartCode/GinExample/pkg/app"
	"GoStartCode/GinExample/pkg/e"
	"GoStartCode/GinExample/pkg/export"
	"GoStartCode/GinExample/pkg/logging"
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

// ExportStudyLog @Summary Export StudyLog
// @Produce  json
// @Param name body string false "Name"
// @Param state body int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/export [post]
func ExportStudyLog(c *gin.Context) {
	appG := app.Gin{C: c}
	var studyLog = study_log_service.StudyLog{}
	filename, err := studyLog.Export()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}

// ImportStudyLog @Summary Import StudyLog
// @Produce  json
// @Param file body file true "Excel File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func ImportStudyLog(c *gin.Context) {
	appG := app.Gin{C: c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	var studyLog = study_log_service.StudyLog{}
	err = studyLog.Import(file)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_IMPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
