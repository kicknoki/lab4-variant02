// Package wallpapercalc реализует расчёт количества обоев, необходимых
// для оклейки помещения, а также формирование итогового отчёта по расходам.
//
// Пакет содержит три обязательные функции по заданию лабораторной работы:
//   - Perimeter      — вычислительная функция (F1),
//   - RollsNeeded     — вспомогательная вычислительная функция,
//   - AddExtraRolls  — функция с указателем, изменяющая значение (F2),
//   - FormatWallpaperReport — функция форматированного вывода (F3).
package wallpapercalc

import "fmt"

// Perimeter вычисляет периметр прямоугольного помещения по его длине и
// ширине (в метрах).
//
// Возвращает ошибку, если length или width меньше либо равны нулю,
// так как отрицательные или нулевые размеры комнаты физически невозможны.
func Perimeter(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		return 0, fmt.Errorf("длина и ширина должны быть положительными: length=%.2f, width=%.2f", length, width)
	}
	return 2 * (length + width), nil
}

// RollsNeeded вычисляет минимальное количество рулонов обоев, необходимое
// для оклейки стен по периметру perimeter и высоте height, при известной
// ширине rollWidth и длине rollLength одного рулона (все размеры в метрах).
//
// Логика: из одного рулона получается несколько полос высотой height
// (stripsPerRoll = rollLength / height). Всего нужно оклеить периметр,
// разбитый на полосы шириной rollWidth (stripsNeeded). Количество рулонов —
// это stripsNeeded, округлённое вверх и делённое на stripsPerRoll.
//
// Возвращает ошибку, если любой из параметров не положителен, либо если
// длина рулона меньше высоты стены (в этом случае не получится вырезать
// ни одной полосы).
func RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error) {
	if perimeter <= 0 || height <= 0 || rollWidth <= 0 || rollLength <= 0 {
		return 0, fmt.Errorf("все параметры должны быть положительными: perimeter=%.2f, height=%.2f, rollWidth=%.2f, rollLength=%.2f",
			perimeter, height, rollWidth, rollLength)
	}

	stripsPerRoll := int(rollLength / height)
	if stripsPerRoll <= 0 {
		return 0, fmt.Errorf("длина рулона (%.2f м) меньше высоты стены (%.2f м)", rollLength, height)
	}

	stripsNeeded := int(perimeter/rollWidth) + 1 // +1 — запас на округление в большую сторону
	rolls := stripsNeeded / stripsPerRoll
	if stripsNeeded%stripsPerRoll != 0 {
		rolls++
	}
	return rolls, nil
}

// AddExtraRolls добавляет запасное количество рулонов extra (например, на
// подрезку рисунка или возможный брак) к уже рассчитанному количеству rolls.
//
// Значение изменяется через указатель *rolls, то есть функция модифицирует
// переменную вызывающего кода напрямую, не возвращая новое значение.
//
// Возвращает ошибку, если extra отрицательно.
func AddExtraRolls(rolls *int, extra int) error {
	if extra < 0 {
		return fmt.Errorf("количество запасных рулонов не может быть отрицательным: %d", extra)
	}
	*rolls += extra
	return nil
}

// FormatWallpaperReport формирует читаемую строку итогового отчёта по
// расчёту обоев для комнаты room: сколько рулонов нужно (rolls) и их
// суммарная стоимость при цене rollPrice за один рулон.
//
// Возвращает ошибку, если rolls или rollPrice отрицательны.
func FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error) {
	if rolls < 0 {
		return "", fmt.Errorf("количество рулонов не может быть отрицательным: %d", rolls)
	}
	if rollPrice < 0 {
		return "", fmt.Errorf("цена рулона не может быть отрицательной: %.2f", rollPrice)
	}
	total := float64(rolls) * rollPrice
	report := fmt.Sprintf("Комната: %-15s | Рулонов: %3d | Цена/рулон: %10.2f тг | Итого: %12.2f тг",
		room, rolls, rollPrice, total)
	return report, nil
}
