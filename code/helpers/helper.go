package helpers

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func GeneratedUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func GetCurrentUnix() int64 {
	location, _ := time.LoadLocation("Asia/Bangkok")
	timeNow := time.Now().In(location)
	return timeNow.Unix()
}
func GetUnixTimestamp(date int64, timeStr string) (int64, error) {
	// โหลดโซนเวลาที่ต้องการ
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return 0, err
	}

	// แปลง date (ซึ่งเป็น Unix timestamp) เป็นเวลาในโซนที่กำหนด
	dateTime := time.Unix(date, 0).In(location)

	// แยกชั่วโมงและนาทีจาก timeStr
	var hour, minute int
	_, err = fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
	if err != nil {
		return 0, err
	}

	// สร้างเวลาที่รวมวันที่และเวลาจาก date และ timeStr ในโซนเวลาที่กำหนด
	fullTime := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), hour, minute, 0, 0, location)

	// แปลงเวลานั้นเป็น Unix timestamp
	return fullTime.Unix(), nil
}
func TimeNowStr() string {
	timeFormat := "2006.01.02 15:04:05"
	location, _ := time.LoadLocation("Asia/Bangkok")
	timeNow := time.Now().In(location)
	return fmt.Sprintf("%s", timeNow.Format(timeFormat))
}

func GetCurrentDateTime() (time.Time, *time.Location) {
	location, _ := time.LoadLocation("Asia/Bangkok")
	timeNow := time.Now().In(location)
	return timeNow, location
}

func GetStartDay() time.Time {
	now, location := GetCurrentDateTime()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return date
}

func GetEndDay() time.Time {
	now, location := GetCurrentDateTime()
	date := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)
	return date
}
