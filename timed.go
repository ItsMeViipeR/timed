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

func Until(date any) (string, error) {
	var targetTime time.Time

	switch v := date.(type) {
	default:
		return "", errors.New("unable to parse date")
	case time.Time:
		targetTime = v
	case string:
		parsedTime, err := StringToTime(v)
		if err != nil {
			return "", err
		}
		targetTime = parsedTime
	}

	duration := time.Until(targetTime)

	return formatUntilDuration(duration), nil
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
	if d < time.Minute {
		return fmt.Sprintf("%d seconds ago", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(d.Hours()))
	}

	days := int(d.Hours() / 24)
	return fmt.Sprintf("%d days ago", days)
}

func formatUntilDuration(d time.Duration) string {
	if d < 0 {
		d = -d
		if d < time.Minute {
			return fmt.Sprintf("%d seconds ago", int(d.Seconds()))
		}
		if d < time.Hour {
			return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
		}
		if d < 24*time.Hour {
			return fmt.Sprintf("%d hours ago", int(d.Hours()))
		}
		days := int((d + time.Hour).Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	}

	if d < time.Minute {
		return fmt.Sprintf("In %d seconds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("In %d minutes", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("In %d hours", int(d.Hours()))
	}

	days := int((d + time.Hour).Hours() / 24)
	return fmt.Sprintf("In %d days", days)
}
