package util

import (
	"GoStartCode/GinExample/pkg/setting"
	"time"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

var weekdays = map[time.Weekday]string{
	time.Sunday:    "星期日",
	time.Monday:    "星期一",
	time.Tuesday:   "星期二",
	time.Wednesday: "星期三",
	time.Thursday:  "星期四",
	time.Friday:    "星期五",
	time.Saturday:  "星期六",
}

var weekdayList = map[string]time.Weekday{
	"星期日": time.Sunday,
	"星期一": time.Monday,
	"星期二": time.Tuesday,
	"星期三": time.Wednesday,
	"星期四": time.Thursday,
	"星期五": time.Friday,
	"星期六": time.Saturday,
}

func GetWeekdayName(weekdayInt int) string {
	return weekdays[time.Weekday(weekdayInt)]
}

func GetWeekdayInt(weekdayName string) int {
	return int(weekdayList[weekdayName])
}

// GetTimeByString
// @Description: 字符串转日期,适配多种日期格式
// @param dateString 日期格式
// @return time.Time 返回time
func GetTimeByString(dateString string) time.Time {
	if dateString == "" {
		return time.Now()
	}

	// 适配多种日期格式
	layoutArr := []string{"2006-1-2", "2006-01-2", "2006-1-02", "2006-01-02"}

	var dateTime time.Time
	var err error

	for _, l := range layoutArr {
		dateTime, err = time.Parse(l, dateString)
		if err == nil {
			break
		}
	}

	if err != nil {
		return time.Now()
	}
	return dateTime
}
