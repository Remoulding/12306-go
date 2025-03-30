package tools

import (
	"fmt"
	"time"
)

var layoutDate = "2006-01-02"
var layoutDetail = "2006-01-02 15:04:05"

// GetDayRange 返回指定日期的起始时间和结束时间
func GetDayRange(dateStr string) (time.Time, time.Time, error) {
	// 解析输入的日期字符串
	loc, _ := time.LoadLocation("Local") // 使用本地时区
	startTime, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// 计算结束时间（23:59:59）
	endTime := startTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	return startTime, endTime, nil
}

func ConvertStrToTime(dateStr string, detail bool) time.Time {
	layout := layoutDate
	if detail {
		layout = layoutDetail
	}
	t, _ := time.Parse(layout, dateStr)
	return t
}

func ConvertTimeToStr(date time.Time, detail bool) string {
	layout := layoutDate
	if detail {
		layout = layoutDetail
	}
	return date.Format(layout)
}

func CalculateHourDifference(start, end string) string {
	startTime := ConvertStrToTime(start, true)
	endTime := ConvertStrToTime(end, true)
	duration := endTime.Sub(startTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%02dh%02dmin", hours, minutes)
}

// ComputeFullTimes 计算完整的出发时间和到达时间
func ComputeFullTimes(departureDate, departureTime, arrivalTime string, crossDay int) (string, string) {
	// 计算出发时间
	departureDateTime, _ := time.Parse(layoutDetail, departureDate+" "+departureTime)
	// 计算到达时间
	arrivalDateTime, _ := time.Parse(layoutDetail, departureDate+" "+arrivalTime)
	// 根据 CrossDay 调整到达时间
	arrivalDateTime = arrivalDateTime.Add(time.Duration(crossDay) * 24 * time.Hour)
	return departureDateTime.Format(layoutDetail), arrivalDateTime.Format(layoutDetail)
}
