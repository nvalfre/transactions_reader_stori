package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"transactions_reader_stori/validators"
)

// ParseAmount parses the amount string into a float64 value
func ParseAmount(amount string) float64 {
	value := strings.ReplaceAll(amount, "+", "")
	value = strings.ReplaceAll(value, "-", "")
	value = strings.TrimSpace(value)

	float := parseFloat(value)

	if validators.IsCredit(amount) {
		return float
	}
	return -float
}

func parseFloat(string string) float64 {
	i, err := strconv.ParseFloat(string, 10)
	if err != nil {
		log.Fatal("Failed to parsing amount id:", err)
	}
	return i
}

func ParseDate(date string) (string, error) {
	parsedDate, err := ParseDateTime(date)
	if err != nil {
		return date, err
	}

	sprint := fmt.Sprint(parsedDate)
	return sprint, err
}

func ParseDateTime(date string) (time.Time, error) {
	currentYear := time.Now().Year()
	dateWithYear := fmt.Sprintf("%d/%s", currentYear, date)
	parsedDate, err := ParseDateWithLayout("2006/1/2", dateWithYear)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

func ParseDateWithLayout(layout, dateWithYear string) (time.Time, error) {
	parsedDate, err := time.Parse(layout, dateWithYear)
	if err != nil {
		log.Println(err)
		return time.Time{}, err
	}
	return parsedDate, nil
}
