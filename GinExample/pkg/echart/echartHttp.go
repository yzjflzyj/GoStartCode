package echart

import (
	"GoStartCode/GinExample/service/study_log_service"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart
func generateItemsByWeek() ([]opts.LineData, []string) {
	var studyLog = study_log_service.StudyLog{}
	//studyLogs, _ := studyLog.QueryStudyLogPage()
	studyLogs, _ := studyLog.QueryStudyLogByDateTime(time.Now().AddDate(0, 0, -7), time.Now())
	//星期常量map
	weekMap := make(map[int]string)
	weekMap[0] = "Sun"
	weekMap[1] = "Mon"
	weekMap[2] = "Tue"
	weekMap[3] = "Wed"
	weekMap[4] = "Thu"
	weekMap[5] = "Fri"
	weekMap[6] = "Sat"
	//展示列表
	showStrList := make([]string, 0)
	items := make([]opts.LineData, 0)

	for _, study := range studyLogs {
		showStrList = append(showStrList, weekMap[study.DayOfWeek])
		items = append(items, opts.LineData{Value: study.StudyTime})
	}
	return items, showStrList
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	items, showStrList := generateItemsByWeek()
	// Put data into instance
	line.SetXAxis(showStrList).
		AddSeries("郭翠霞是2b", items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func Setup() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8000", nil)
}
