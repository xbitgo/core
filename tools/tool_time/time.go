package tool_time

import "time"

// TimeToDateTimeString .
func TimeToDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseDateTime .
func ParseDateTime(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

// TimeToDateString .
func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDate .
func ParseDate(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", str, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

// ParseDateEnd .
func ParseDateEnd(str string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", str, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t.Add(24 * time.Hour).Add(-1 * time.Second)
}
