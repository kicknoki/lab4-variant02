// Command app демонстрирует работу пакета wallpapercalc: расчёт количества
// рулонов обоев для комнаты и вывод итогового отчёта в консоль.
package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"

	"lab4-variant02/pkg/wallpapercalc"
)

func main() {
	// Внешний пакет №1 (google/uuid) — генерируем идентификатор операции,
	// как будто это "номер заявки" в реальной системе учёта заказов.
	opID := uuid.New().String()
	fmt.Printf("=== Расчёт обоев (ID заявки: %s) ===\n", opID)

	// Исходные данные (в реальном проекте могли бы вводиться пользователем).
	const (
		length     = 5.0    // длина комнаты, м
		width      = 4.0    // ширина комнаты, м
		height     = 2.7    // высота стен, м
		rollWidth  = 0.53   // ширина рулона, м
		rollLength = 10.05  // длина рулона, м
		extraRolls = 2      // запас рулонов
		rollPrice  = 4500.0 // цена одного рулона, тг
		room       = "Гостиная"
	)

	// 1. F1: вычислительная функция.
	perimeter, err := wallpapercalc.Perimeter(length, width)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	fmt.Printf("Периметр комнаты: %.2f м\n", perimeter)

	rolls, err := wallpapercalc.RollsNeeded(perimeter, height, rollWidth, rollLength)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	fmt.Printf("Базовое количество рулонов: %d\n", rolls)

	// 2. F2: функция с указателем, изменяющая значение rolls "на месте".
	if err := wallpapercalc.AddExtraRolls(&rolls, extraRolls); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	fmt.Printf("С учётом запаса (+%d): %d рулонов\n", extraRolls, rolls)

	// 3. F3: формирование строки отчёта.
	report, err := wallpapercalc.FormatWallpaperReport(room, rolls, rollPrice)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	fmt.Println(report)

	// Внешний пакет №2 (dustin/go-humanize) — показываем итоговую сумму
	// "по-человечески" (с разделителями разрядов), например для счёта клиенту.
	total := float64(rolls) * rollPrice
	fmt.Printf("Итоговая сумма прописью: %s тг\n", humanize.Commaf(total))
}
