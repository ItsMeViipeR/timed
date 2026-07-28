package timed

import (
	"errors"
	"fmt"
	"time"
)

func Since(date any) (string, error) {
	var targetTime time.Time

	switch v := date.(type) {
	default:
		return "", errors.New("Unable to parse date")
	case time.Time:
		targetTime = v
	case string:
		parsedTime, err := StringToTime(v)

		if err != nil {
			return "", err
		}

		targetTime = parsedTime
	}

	duration := time.Since(targetTime)

	fmt.Println(formatDuration(duration))

	return formatDuration(duration), nil
}

func StringToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, fmt.Errorf("Invalid date format (expected YYYY-MM-DD HH:MM:SS): %w", err)
	}

	return parsedTime, nil
}

func formatDuration(d time.Duration) string {
	if d < 0 {
		return "in the future"
	}
	if d < time.Minute {
		return fmt.Sprintf("%.0f seconds ago", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%.0f minutes ago", d.Minutes())
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%.0f hours ago", d.Hours())
	}

	days := int(d.Hours() / 24)

	return fmt.Sprintf("%d days ago", days)
}
