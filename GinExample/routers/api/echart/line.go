package echart

import (
	"GoStartCode/GinExample/service/study_log_service"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"strconv"
	"time"
)

// 组装bar所需数据
func generateAllItemsLine() ([]opts.LineData, []string, int) {
	var studyLog = study_log_service.StudyLog{}
	studyLogs, _ := studyLog.QueryStudyLogByDateTime(time.Now().AddDate(-6, 0, 0), time.Now())
	//展示列表
	showStrList := make([]string, 0)
	items := make([]opts.LineData, 0)
	totalStudyTime := 0
	if len(studyLogs) == 0 {
		return items, showStrList, totalStudyTime
	}
	//星期常量map
	weekMap := make(map[int]string)
	weekMap[0] = "Sun"
	weekMap[1] = "Mon"
	weekMap[2] = "Tue"
	weekMap[3] = "Wed"
	weekMap[4] = "Thu"
	weekMap[5] = "Fri"
	weekMap[6] = "Sat"

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
			items = append(items, opts.LineData{Value: studyLogs[index].StudyTime})
			totalStudyTime += studyLogs[index].StudyTime
			index++
		} else {
			items = append(items, opts.LineData{Value: 0})
		}
	}
	return items, showStrList, totalStudyTime
}

func getLine() *charts.Line {
	// 获取数据
	items, showStrList, totalStudyTime := generateAllItemsLine()
	earliestStudyDate, lastStudyDate := getStudyDateStr(showStrList)
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title:    "总学习时长： " + strconv.Itoa(totalStudyTime) + " min",
			Subtitle: "最早学习时间：" + earliestStudyDate + "  最近学习时间：" + lastStudyDate,
			Link:     "https://github.com/go-echarts/examples",
			Right:    "70%",
		}),
		//charts.WithColorsOpts(opts.Colors{"blue"}), // 设置颜色
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "10%"}),
		// 调整图的大小
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "2000px",
			Height: "800px",
		}),
	)

	// Put data into instance
	line.SetXAxis(showStrList).
		AddSeries("学习记录", items).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			// 标志点
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
			//title and label options
			charts.WithLineChartOpts(opts.LineChart{ShowSymbol: true}),
			charts.WithLabelOpts(opts.Label{Show: true}),
		)

	return line
}

var (
	itemCntLine = 6
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and label options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		)
	return line
}

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems(),
		charts.WithLabelOpts(
			opts.Label{Show: true},
		))
	return line
}

func lineStep() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "step style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Step: true,
			}),
		)
	return line
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "smooth style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "area options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: true,
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				}),
		)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineOverlap() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	line.Overlap(esEffectStyle())
	line.Overlap(scatterBase())
	return line
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category  A", generateLineItems()).
		AddSeries("Category  B", generateLineItems()).
		AddSeries("Category  C", generateLineItems()).
		AddSeries("Category  D", generateLineItems())
	return line
}

func lineDemo() *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Search Time: Hash table vs Binary search",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Cost time(ns)",
			SplitLine: &opts.SplitLine{
				Show: false,
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Elements",
		}),
	)

	line.SetXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddSeries("map", generateLineItems(),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "bottom"})).
		AddSeries("slice", generateLineData([]float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top"})).
		SetSeriesOptions(
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "Average",
				Type: "average",
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: true,
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show:      true,
					Formatter: "{a}: {b}",
				},
			}),
		)

	return line
}

func lineSymbols() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: false, ShowSymbol: true},
		))

	return line
}
