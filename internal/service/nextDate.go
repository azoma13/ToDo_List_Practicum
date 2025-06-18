package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/configs"
)

func NextDate(now time.Time, dstart, repeat string) (string, error) {

	dstartTime, err := time.Parse(configs.DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("error dstart cannot be converted to a valid date")
	}

	splitRepeat := strings.Split(repeat, " ")

	switch splitRepeat[0] {
	case "d":
		nextDate, err := caseDay(now, dstartTime, splitRepeat)
		if err != nil {
			return "", err
		}

		return nextDate, nil
	case "y":
		if len(splitRepeat) != 1 {
			return "", fmt.Errorf("error parameter does not require additional clarifications.")
		}

		return caseYear(now, dstartTime), nil
	case "w":
		nextDate, err := caseWeekDay(now, dstartTime, splitRepeat)
		if err != nil {
			return "", err
		}

		return nextDate, nil
	case "m":
		return "", nil
	default:
		return "", fmt.Errorf("invalid character")
	}
}

func caseDay(now, dstartTime time.Time, splitRepeat []string) (string, error) {

	if len(splitRepeat) != 2 {
		return "", fmt.Errorf("error the interval in days is not specified or incorrectly")
	}

	nextDay, err := strconv.Atoi(splitRepeat[1])
	if err != nil {
		return "", fmt.Errorf("error converted atoi day: %v", err)
	}

	if nextDay > 400 {
		return "", fmt.Errorf("error the maximum allowed interval has been exceeded")
	}

	dstartTime = dstartTime.AddDate(0, 0, nextDay)
	for afterNow(dstartTime, now) {
		dstartTime = dstartTime.AddDate(0, 0, nextDay)
	}

	return dstartTime.Format(configs.DateFormat), nil
}

func caseYear(now, dstartTime time.Time) string {

	dstartTime = dstartTime.AddDate(1, 0, 0)
	for afterNow(dstartTime, now) {
		dstartTime = dstartTime.AddDate(1, 0, 0)
	}

	return dstartTime.Format(configs.DateFormat)
}

func caseWeekDay(now, dstartTime time.Time, splitRepeat []string) (string, error) {

	if len(splitRepeat) != 2 {
		return "", fmt.Errorf("error the weekday is not specified or incorrectly")
	}

	weekdays := strings.Split(splitRepeat[1], ",")
	if len(weekdays) == 0 {
		return "", fmt.Errorf("error incorrectly weekdays")
	}

	var arrWeekDay [7]bool
	for _, weekDay := range weekdays {
		wd, err := strconv.Atoi(weekDay)
		if err != nil {
			return "", fmt.Errorf("error converted atoi weekday: %v", err)
		}
		if wd > 7 || wd < 1 {
			return "", fmt.Errorf("error weekday specified incorrectly")
		}
		arrWeekDay[wd-1] = true
	}

	if afterNow(dstartTime, now) {
		dstartTime = now
	}

	dstartTime = dstartTime.AddDate(0, 0, 1)
	for {
		if arrWeekDay[int(dstartTime.Weekday())] {
			dstartTime = dstartTime.AddDate(0, 0, 1)
			break
		}
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	return dstartTime.Format(configs.DateFormat), nil
}

func afterNow(dstartTime, now time.Time) bool {
	oneDay := 24 * time.Hour

	dstartTime = dstartTime.Truncate(oneDay)
	now = now.Truncate(oneDay)

	return dstartTime.Before(now)
}
