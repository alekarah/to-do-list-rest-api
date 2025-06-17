package repeat

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NextDate вычисляет следующую дату для задачи в соответствии с правилом повторения
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Проверяем, что repeat не пустая строка
	if repeat == "" {
		return "", fmt.Errorf("правило повторения не может быть пустым")
	}

	// Парсим исходную дату
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("некорректная дата начала: %v", err)
	}

	// Разбиваем правило повторения на части
	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		// Правило для дней
		if len(parts) != 2 {
			return "", fmt.Errorf("неверный формат правила 'd': ожидается 'd <число>'")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("некорректное число дней: %v", err)
		}

		if interval <= 0 || interval > 400 {
			return "", fmt.Errorf("число дней должно быть от 1 до 400, получено: %d", interval)
		}

		// Увеличиваем дату на указанное количество дней до тех пор, пока она не станет больше now
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}

	case "y":
		// Правило для года
		if len(parts) != 1 {
			return "", fmt.Errorf("неверный формат правила 'y': не должно быть дополнительных параметров")
		}

		// Увеличиваем дату на год до тех пор, пока она не станет больше now
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

	case "w":
		// Правило для дней недели
		if len(parts) != 2 {
			return "", fmt.Errorf("неверный формат правила 'w': ожидается 'w <дни_недели>'")
		}

		// Парсим дни недели
		weekdays, err := parseWeekdays(parts[1])
		if err != nil {
			return "", fmt.Errorf("ошибка парсинга дней недели: %v", err)
		}

		// Ищем следующую дату, которая попадает на один из указанных дней недели
		date = findNextWeekday(date, now, weekdays)

	case "m":
		// Правило для дней месяца
		if len(parts) < 2 || len(parts) > 3 {
			return "", fmt.Errorf("неверный формат правила 'm': ожидается 'm <дни> [месяцы]'")
		}

		// Парсим дни месяца
		days, err := parseDays(parts[1])
		if err != nil {
			return "", fmt.Errorf("ошибка парсинга дней месяца: %v", err)
		}

		// Парсим месяцы (если указаны)
		var months []int
		if len(parts) == 3 {
			months, err = parseMonths(parts[2])
			if err != nil {
				return "", fmt.Errorf("ошибка парсинга месяцев: %v", err)
			}
		}

		// Ищем следующую дату, которая попадает на один из указанных дней месяца
		date = findNextMonthday(date, now, days, months)

	default:
		return "", fmt.Errorf("неподдерживаемый формат правила: %s", rule)
	}

	// Преобразуем дату в строку и возвращаем
	return date.Format("20060102"), nil
}

// afterNow проверяет, что первая дата больше второй (игнорируя время)
func afterNow(date, now time.Time) bool {
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return dateOnly.After(nowOnly)
}

// ShouldDelete проверяет, должна ли задача быть удалена при завершении
// Возвращает true, если правило повторения не указано (пустая строка)
func ShouldDelete(repeat string) bool {
	return repeat == ""
}

// ! затем в основном коде нужно будет сделать проверку, если ShouldDelete() true, то удаляем задачу
// если нет, то вычисляем следующую дату NextDate()
