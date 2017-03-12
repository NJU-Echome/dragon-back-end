package util

import (
	"database/sql/driver"
	// "fmt"
	"time"
)

// GetCurrentTime 获取当前时间 string 格式
func GetCurrentTimeStr() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// GetCurrentTime 获取当前时间 time.Time 格式
func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeStr := time.Now().In(loc).Format("2006-01-02 15:04:05")
	return StrToTime(timeStr)
}

// GetTodayDate 获取当前日期 time.Time 格式
func GetTodayDate() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	// today_str := fmt.Sprintf("%d-%d-%d 00:00:00", year, month, day)
	todayTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return todayTime
}

func GetEndOfThatDay(t time.Time) time.Time {
	year, month, day := t.Date()
	thatDay := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return thatDay
}

// GetIntervalToNow 获取上一个时间点到当前的时间间隔 秒 lastTime< now 返回值大于0 否则返回值小于0
func GetIntervalToNow(lastTime time.Time) int {
	// fmt.Println(time.Now())
	// fmt.Println(lastTime)
	duration := time.Now().Sub(lastTime)
	return int(duration.Seconds())
}

func StrToTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
