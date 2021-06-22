package clock

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Talk translates a time specified in
func Talk(value string) (string, error) {
	hours, minutes := splitTime(value)

	if minutes == 0 {
		if hours == 0 {
			return capitalise(convertHoursToWord(hours)), nil
		}
		if hours == 12 {
			return capitalise("noon"), nil
		}

		return capitalise(fmt.Sprintf("%s o'clock", convertHoursToWord(hours))), nil
	}

	if minutes == 30 {
		return capitalise(fmt.Sprintf("half past %s", convertHoursToWord(hours))), nil
	}

	if minutes > 0 && minutes < 30 {
		return capitalise(fmt.Sprintf("%s past %s",
			convertMinutesToWord(minutes),
			convertHoursToWord(hours))), nil
	}

	if minutes > 30 && minutes < 60 {
		nextHour := getNextHour(hours)
		return capitalise(fmt.Sprintf("%s to %s",
			convertMinutesToWord(minutes),
			convertHoursToWord(nextHour))), nil
	}

	return "", fmt.Errorf("can't parse the date")
}

func splitTime(value string) (int, int) {
	values := strings.Split(value, ":")

	hours, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println("Cannot convert hours")
		os.Exit(1)
	}

	minutes, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println("Cannot convert minutes")
		os.Exit(1)
	}

	return hours, minutes
}

func capitalise(s string) string {
	return strings.Title(string(s[0])) + string(s[1:])
}

func GetCurrentTime() string {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()
	return fmt.Sprintf("%d:%d", currentHour, currentMinute)
}
