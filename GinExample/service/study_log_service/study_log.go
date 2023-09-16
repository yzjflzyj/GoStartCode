package study_log_service

import (
	"GoStartCode/GinExample/models"
	"GoStartCode/GinExample/pkg/export"
	"GoStartCode/GinExample/pkg/file"
	"GoStartCode/GinExample/pkg/util"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
	"io"
	"math"
	"strconv"
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

func (studyLog *StudyLog) GetAll() ([]models.StudyLog, error) {
	return models.QueryStudyLogPage(0, math.MaxInt64)
}

func (studyLog *StudyLog) QueryStudyLogByDateTime(begin time.Time, end time.Time) ([]models.StudyLog, error) {
	return models.QueryStudyLogByDateTime(begin, end)
}

func (studyLog *StudyLog) Export() (string, error) {
	studyLogList, err := studyLog.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("全量学习记录")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "用户id", "星期", "学习时间（min）", "日期", "内容", "创建者", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, study := range studyLogList {
		values := []string{
			strconv.Itoa(study.ID),
			strconv.Itoa(study.UserId),
			util.GetWeekdayName(study.DayOfWeek),
			strconv.Itoa(study.StudyTime),
			study.DateTime.Format("2006-01-02 15:04:05"),
			study.Content,
			study.CreatedBy,
			study.CreatedOn.Format("2006-01-02 15:04:05"),
			study.ModifiedBy,
			study.ModifiedOn.Format("2006-01-02 15:04:05"),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	nowTime := strconv.Itoa(int(time.Now().Unix()))
	filename := "study_log" + nowTime + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (studyLog *StudyLog) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("全量学习记录")
	for iRow, row := range rows {
		if iRow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			userId, _ := strconv.Atoi(data[1])
			dayOfWeek := util.GetWeekdayInt(data[2])
			studyTime, _ := strconv.Atoi(data[3])
			var studyLog = StudyLog{
				UserId:     userId,
				DayOfWeek:  dayOfWeek,
				StudyTime:  studyTime,
				DateTime:   time.Now(),
				Content:    data[5],
				CreatedBy:  data[6],
				CreatedOn:  time.Now(),
				ModifiedBy: data[8],
				ModifiedOn: time.Now(),
				State:      1,
				DeletedOn:  time.Now(),
			}
			studyLog.Add()
		}
	}

	return nil
}
