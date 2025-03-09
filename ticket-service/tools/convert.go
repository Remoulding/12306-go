package tools

import (
	"fmt"
	"time"
)

var layout = "2006-01-02"

func ConvertStrToTime(dateStr string) time.Time {
	t, _ := time.Parse(layout, dateStr)
	return t
}

func ConvertTimeToStr(date time.Time) string {
	return date.Format(layout)
}

func CalculateHourDifference(startTime, endTime time.Time) string {
	duration := endTime.Sub(startTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
