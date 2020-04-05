package utils

import (
	"fmt"
	"time"
)

const (
	DATE_FORMAT_YYYY_MM_DD_HH_MM_SS = "1970-01-01 08:00:00"
	DATE_FORMAT_YYYY_MM_DD          = "1970-01-01"
	DATE_FORMAT_HH_MM_SS            = "12:00:00"
	DATE_FROMAT_YYYYMMDDHHMMSS      = "20060102150405"
)

func GetYYYYMMDDDate(date time.Time) string {
	return date.Format(DATE_FORMAT_YYYY_MM_DD)
}

func GetYesterdayBeginning() time.Time {
	return FormatDayBeginning(GetYesterdayDate())
}

func FormatDayBeginning(day time.Time) time.Time {
	beginning := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())

	return beginning
}

func IsSameDay(firstTime time.Time, secondTime time.Time) bool {
	return time.Date(firstTime.Year(), firstTime.Month(), firstTime.Day(), 0, 0, 0, 0, firstTime.Location()).Equal(
		time.Date(secondTime.Year(), secondTime.Month(), secondTime.Day(), 0, 0, 0, 0, secondTime.Location()))
}

func GetYesterdayDate() time.Time {
	currentTime := time.Now()

	return currentTime.AddDate(0, 0, -1)
}
func GetOffsetToday(offset int) time.Time {
	currentTime := time.Now()

	return currentTime.AddDate(0, 0, offset)
}
func GetBeforeYesterday6DateTime() time.Time {
	yesterday := GetYesterdayDate()

	return yesterday.AddDate(0, 0, -1).Add(time.Hour * 6)
}

func GetYesterday6DateTime() time.Time {
	yesterday := GetYesterdayDate()

	return yesterday.Add(time.Hour * 6)
}

func GetToday6DateTime() time.Time {
	currentTime := time.Now()

	return currentTime.Add(time.Hour * 6)
}

func GetTimeDaysAgo(n int) time.Time {
	return time.Now().AddDate(0, 0, -n)
}

func GetTimeHoursAgo(n int) time.Time {
	temp, _ := time.ParseDuration(fmt.Sprintf("-%dh", n))
	return time.Now().Add(temp)
}

func LimitDateTime(date time.Time, limitDate string) time.Time {
	timeString := LimitDate(TimeToString(date), limitDate)
	return TimeStringToTime(timeString)
}
func LimitDate(date string, limitDate string) string {
	//"2019-12-01"
	loc, _ := time.LoadLocation("Local")
	limit, _ := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD, limitDate, loc)
	dateT, err := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD, date, loc)
	if err != nil {
		return date
	}

	if dateT.Unix() < limit.Unix() {
		return limitDate
	}
	return date
}

func CompareDate(date string, limitDate string) bool {
	//"2019-12-01"
	loc, _ := time.LoadLocation("Local")
	limit, _ := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD, limitDate, loc)
	dateT, err := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD, date, loc)
	if err != nil {
		return false
	}

	if dateT.Unix() < limit.Unix() {
		return false
	}
	return true
}

// 字符串时间转unix
func TimeStringToUnix(datetime string) int64 {
	//日期转化为时间戳
	if datetime == "" {
		datetime = DATE_FORMAT_YYYY_MM_DD_HH_MM_SS
	}
	loc, err := time.LoadLocation("Asia/Shanghai") //获取时区
	if err != nil {
		loc = nil
	}
	tmp, err := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD_HH_MM_SS, datetime, loc)
	if err != nil {
		tmp = time.Time{}
	}
	return tmp.Unix() //转化为时间戳 类型是int64
}
func TimeStringToTime(datetime string) time.Time {
	//日期转化为时间戳
	if datetime == "" {
		datetime = DATE_FORMAT_YYYY_MM_DD_HH_MM_SS
	}
	loc, err := time.LoadLocation("Asia/Shanghai") //获取时区
	if err != nil {
		loc = nil
	}
	tmp, err := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD_HH_MM_SS, datetime, loc)
	if err != nil {
		tmp = time.Time{}
	}
	return tmp //转化为时间戳 类型是int64
}
func DateStringToTime(datetime string) time.Time {
	//日期转化为时间戳
	if datetime == "" {
		datetime = DATE_FORMAT_YYYY_MM_DD
	}
	loc, err := time.LoadLocation("Asia/Shanghai") //获取时区
	if err != nil {
		loc = nil
	}
	tmp, err := time.ParseInLocation(DATE_FORMAT_YYYY_MM_DD, datetime, loc)
	if err != nil {
		tmp = time.Time{}
	}
	return tmp //转化为时间戳 类型是int64
}
func UnixToTimeString(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format(DATE_FORMAT_YYYY_MM_DD_HH_MM_SS)
}
func TimeToString(ti time.Time) string {
	return ti.Format(DATE_FORMAT_YYYY_MM_DD_HH_MM_SS)
}

func GetTimeNow() string {
	return time.Now().Format(DATE_FORMAT_YYYY_MM_DD_HH_MM_SS)
}
