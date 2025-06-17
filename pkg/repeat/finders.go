package repeat

import (
	"time"
)

// findNextWeekday находит следующую дату, которая попадает на один из указанных дней недели
func findNextWeekday(date, now time.Time, weekdays []int) time.Time {
	// Создаем массив допустимых дней недели
	var validDays [8]bool // индексы 1-7
	for _, day := range weekdays {
		validDays[day] = true
	}

	// Начинаем с даты начала и ищем следующий подходящий день
	for {
		// Проверяем, подходит ли текущий день недели
		// В Go Sunday = 0, Monday = 1, ..., Saturday = 6
		// Нам нужно Monday = 1, ..., Sunday = 7
		weekday := int(date.Weekday())
		if weekday == 0 {
			weekday = 7 // Sunday
		}

		if validDays[weekday] && afterNow(date, now) {
			return date
		}

		date = date.AddDate(0, 0, 1)
	}
}

// findNextMonthday находит следующую дату, которая попадает на один из указанных дней месяца
func findNextMonthday(date, now time.Time, days []int, months []int) time.Time {
	// Создаем массивы допустимых дней и месяцев
	var validDays [32]bool   // индексы 1-31
	var validMonths [13]bool // индексы 1-12

	for _, day := range days {
		if day > 0 {
			validDays[day] = true
		}
	}

	// Если месяцы не указаны, разрешаем все
	if len(months) == 0 {
		for i := 1; i <= 12; i++ {
			validMonths[i] = true
		}
	} else {
		for _, month := range months {
			validMonths[month] = true
		}
	}

	// Начинаем с даты начала и ищем следующий подходящий день
	for {
		if validMonths[int(date.Month())] && afterNow(date, now) {
			// Проверяем обычные дни месяца
			if validDays[date.Day()] {
				return date
			}

			// Проверяем отрицательные дни (-1, -2)
			for _, day := range days {
				if day < 0 && isValidNegativeDay(date, day) {
					return date
				}
			}
		}

		date = date.AddDate(0, 0, 1)
	}
}

// isValidNegativeDay проверяет, соответствует ли дата отрицательному дню месяца
func isValidNegativeDay(date time.Time, negDay int) bool {
	// Получаем последний день месяца
	firstOfNextMonth := time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)

	switch negDay {
	case -1:
		return date.Day() == lastOfMonth.Day()
	case -2:
		return date.Day() == lastOfMonth.Day()-1
	default:
		return false
	}
}
