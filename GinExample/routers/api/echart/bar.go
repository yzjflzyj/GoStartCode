package echart

import (
	"GoStartCode/GinExample/service/study_log_service"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func generateItemsByWeekBar() ([]opts.BarData, []string, int) {
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
	items := make([]opts.BarData, 0)

	totalStudyTime := 0
	firstStudyLogDate := studyLogs[0].DateTime
	lastStudyLogDate := studyLogs[len(studyLogs)-1].DateTime
	duration := lastStudyLogDate.Sub(firstStudyLogDate)
	days := int(duration.Hours()/24) + 1
	index := 0
	// 将没有记录的，也展示
	for i := 0; i < days; i++ {
		date := firstStudyLogDate.AddDate(0, 0, i)
		showStrList = append(showStrList, weekMap[int(date.Weekday())]+"\n"+date.Format("2006-01-02"))
		//判断同一天date1.Year() == date2.Year() && date1.YearDay() == date2.YearDay()
		if studyLogs[index].DateTime.Year() == date.Year() && studyLogs[index].DateTime.YearDay() == date.YearDay() {
			items = append(items, opts.BarData{Value: studyLogs[index].StudyTime})
			totalStudyTime += studyLogs[index].StudyTime
			index++
		} else {
			items = append(items, opts.BarData{Value: 0})
		}
	}
	return items, showStrList, totalStudyTime
}

func getBar() *charts.Bar {
	// 获取数据
	items, showStrList, totalStudyTime := generateItemsByWeekBar()
	earliestStudyDate, lastStudyDate := getStudyDateStr(showStrList)
	// create a new line instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWalden}),
		charts.WithTitleOpts(opts.Title{
			Title:    "总学习时长： " + strconv.Itoa(totalStudyTime) + " min",
			Subtitle: "最早学习时间：" + earliestStudyDate + "  最近学习时间：" + lastStudyDate,
			Link:     "https://github.com/go-echarts/examples",
			Right:    "70%",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "10%"}),
		// 调整图的大小
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "2000px",
			Height: "800px",
		}),
	)

	// Put data into instance
	bar.SetXAxis(showStrList).
		AddSeries("学习记录", items).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			//在顶端展示数值
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top"}))
	return bar
}

func getStudyDateStr(showStrList []string) (string, string) {
	dateRegex := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

	// 提取匹配的日期字符串
	earliestStudyDate := dateRegex.FindStringSubmatch(showStrList[0])[0]
	lastStudyDate := dateRegex.FindStringSubmatch(showStrList[len(showStrList)-1])[0]
	return earliestStudyDate, lastStudyDate
}

func ChartHandler(c *gin.Context) {
	bar := getBar()
	// 渲染图表到一个HTML页面
	page := components.NewPage()
	page.AddCharts(bar,
		barBasic(),
		barTitle(),
		barTooltip(),
		barSetToolbox(),
		barShowLabel(),
		barXYName(),
		barXYFormatter(),
		barColor(),
		barSplitLine(),
		barGap(),
		barDataZoomInside(),
		barDataZoomSlider(),
		barReverse(),
		barStack(),
		barMarkPoints(),
		barMarkLines(),
		barOverlap(),
		barSize(),
		//条形图
		lineBase(),
		lineShowLabel(),
		lineSymbols(),
		lineMarkPoint(),
		lineSplitLine(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		lineOverlap(),
		lineMulti(),
		lineDemo())
	page.Render(c.Writer)

	c.Status(http.StatusOK)
}

// 以下为各种样式的示例
var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func barBasic() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic bar example", Subtitle: "This is the subtitle."}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTitle() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and legend options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
			Right:    "40%",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80%"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTooltip() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "tooltip options"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barSetToolbox() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "toolbox options"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					// set the language
					// Chinese version: ["数据视图", "关闭", "刷新"]
					Lang: []string{"data view", "turn off", "refresh"},
				},
			}},
		),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barShowLabel() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "label options"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)
	return bar
}

func barXYName() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "display the axes name",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barXYFormatter() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "customized the xaxis/yaxis formatter"}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} x-unit"},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} y-unit"},
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barColor() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "set user-defined colors",
		}),
		charts.WithColorsOpts(opts.Colors{"blue", "pink"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barSplitLine() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barGap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "set the gap of each bar",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.SetSeriesOptions(
		charts.WithBarChartOpts(opts.BarChart{
			BarGap: "150%",
		}),
	)
	return bar
}

func barDataZoomInside() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "datazoom options(inside)",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "inside",
			Start: 10,
			End:   50,
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barDataZoomSlider() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "datazoom options(slider)",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "slider",
			Start: 10,
			End:   50,
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "reverse xaxis and yaxis",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.XYReversal()
	return bar
}

func barStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "stack style",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
			Stack: "stackA",
		}))
	return bar
}

func barMarkPoints() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	special := generateBarItems()
	special[0].Value = 100

	bar.SetXAxis(weeks).
		AddSeries("Category A", special, charts.WithMarkPointNameCoordItemOpts(opts.MarkPointNameCoordItem{
			Name:       "special mark",
			Coordinate: []interface{}{"Mon", 100},
			Label: &opts.Label{
				Show:     true,
				Color:    "pink",
				Position: "inside",
			},
		})).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkPointNameTypeItemOpts(
			opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
		))
	return bar
}

func barMarkLines() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markline options",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkLineNameTypeItemOpts(
			opts.MarkLineNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkLineNameTypeItem{Name: "Avg", Type: "average"},
		))
	return bar
}

func barOverlap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkLineNameTypeItemOpts(
			opts.MarkLineNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkLineNameTypeItem{Name: "Avg", Type: "average"},
		))
	//bar.Overlap(lineBase())
	//bar.Overlap(scatterBase())
	return bar
}

func barSize() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "adjust canvas size",
			Subtitle: "I want a bigger canvas size :)",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "600px",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}
