package ps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func convertLstart(datetime string) (time.Time, error) {
	// The layout used when issuing: s -e -o lstart
	const lstart_layout = "Mon Jan 02 15:04:05 2006"

	t, err := time.Parse(lstart_layout, datetime)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func convertCpuTime(t string) (time.Duration, error) {
	parts := strings.Split(t, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format: %s, expected HH:MM:SS", t)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid duration hours: %v", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid duration minutes: %v", err)
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid duration seconds: %v", err)
	}

	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	return duration, nil
}
