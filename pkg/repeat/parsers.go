package repeat

import (
	"fmt"
	"strconv"
	"strings"
)

// parseWeekdays парсит строку с днями недели (1-7)
func parseWeekdays(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	weekdays := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		day, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("некорректный день недели: %s", part)
		}
		if day < 1 || day > 7 {
			return nil, fmt.Errorf("день недели должен быть от 1 до 7, получено: %d", day)
		}
		weekdays = append(weekdays, day)
	}

	return weekdays, nil
}

// parseDays парсит строку с днями месяца
func parseDays(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	days := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		day, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("некорректный день месяца: %s", part)
		}
		if day < -2 || day == 0 || day > 31 {
			return nil, fmt.Errorf("день месяца должен быть от 1 до 31 или -1, -2, получено: %d", day)
		}
		days = append(days, day)
	}

	return days, nil
}

// parseMonths парсит строку с месяцами
func parseMonths(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	months := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		month, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("некорректный месяц: %s", part)
		}
		if month < 1 || month > 12 {
			return nil, fmt.Errorf("месяц должен быть от 1 до 12, получено: %d", month)
		}
		months = append(months, month)
	}

	return months, nil
}
