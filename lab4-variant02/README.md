# lab4-variant02 — wallpapercalc (расчёт обоев)

Лабораторная работа №4: собственный пакет, документация и подключение
пакетов из GitHub. Вариант 2 — калькулятор обоев.

## Структура проекта

```
lab4-variant02/
├── go.mod
├── go.sum
├── README.md
├── cmd/app/main.go              # точка входа, использует пакет и внешние библиотеки
└── pkg/wallpapercalc/wallpapercalc.go   # сам пакет с 4 функциями
```

## Что делает пакет `wallpapercalc`

| Функция | Назначение |
|---|---|
| `Perimeter(length, width float64) (float64, error)` | считает периметр комнаты |
| `RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error)` | считает базовое число рулонов |
| `AddExtraRolls(rolls *int, extra int) error` | добавляет запас рулонов **через указатель** (меняет значение напрямую) |
| `FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error)` | форматирует итоговую строку отчёта |

Все функции проверяют входные данные и возвращают `error` через `fmt.Errorf`,
если что-то некорректно (отрицательные/нулевые значения и т. п.).

## Внешние пакеты (GitHub)

- `github.com/google/uuid` — генерирует уникальный ID заявки.
- `github.com/dustin/go-humanize` — форматирует итоговую сумму с разделителями
  разрядов (например, `63,000` вместо `63000`).

## Как запустить

Нужен установленный Go (1.22 или новее — `go version`).

```bash
cd lab4-variant02
go mod tidy      # скачает зависимости и обновит go.sum
go run ./cmd/app
```

Ожидаемый вывод:

```
=== Расчёт обоев (ID заявки: <случайный uuid>) ===
Периметр комнаты: 18.00 м
Базовое количество рулонов: 12
С учётом запаса (+2): 14 рулонов
Комната: Гостиная        | Рулонов:  14 | Цена/рулон:    4500.00 тг | Итого:     63000.00 тг
Итоговая сумма прописью: 63,000 тг
```

## Проверка документации

```bash
go doc ./pkg/wallpapercalc                  # список функций пакета + описание пакета
go doc ./pkg/wallpapercalc Perimeter         # документация по одной функции
```

Через веб-интерфейс (если `godoc` не установлен — `go install golang.org/x/tools/cmd/godoc@latest`):

```bash
godoc -http=:6060
# открыть в браузере: http://localhost:6060/pkg/lab4-variant02/pkg/wallpapercalc/
```

## Публикация на GitHub

```bash
git init
git add .
git commit -m "lab4 variant02: wallpapercalc package"
git branch -M main
git remote add origin https://github.com/<ваш_логин>/lab4-variant02.git
git push -u origin main
```

## Пункт 10 задания — импорт «чужого» пакета

Требование: импортировать готовый пакет одногруппника или ещё один публичный
пакет и вызвать хотя бы одну его функцию. В этом проекте эту роль уже играет
`github.com/dustin/go-humanize` (второй внешний пакет, помимо `google/uuid`).
Если преподаватель требует именно пакет **одногруппника**, замените
`go-humanize` на его модуль:

```bash
go get github.com/<логин_одногруппника>/<его_пакет>@latest
```

и вызовите одну из его экспортируемых функций в `main.go`.
