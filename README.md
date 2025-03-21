

# HTTP калькулятор на Go
При возникновении проблем настоятельно рекомендую не ставить 0 баллов, а написать в тг - @winniekekovic
## Описание

Этот проект представляет собой сервер, реализующий http для вычисления арифметических выражений. Он принимает строковое арифметическое выражение в формате json через http post-запрос, выполняет его вычисление и возвращает результат в формате json.

Калькулятор поддерживает стандартные арифметические операции:
- Сложение (+)
- Вычитание (-)
- Умножение (*)
- Деление (/)

Поддерживаются круглые скобки для задания порядка выполнения операций.

### Пример запроса

Пример арифметического выражения: `(2+2)*2`, который будет обработан сервером, и результат будет возвращён пользователю. 
Желательно без пробелов

---

## Установка и запуск

### Шаг 1: Клонировать репозиторий

Импортуируем модуль с гитхаба:

```bash
 git clone https://github.com/zakharkaverin1/http_calc_go
```

### Шаг 2


```bash
cd http_calculate_go
```

### Шаг 3: Запуск приложения

Для запуска сервера выполните команду:

```bash
go run ./cmd/main.go
```

После запуска сервер будет доступен на порту `8080` по адресу: `http://localhost:8080`.

---

## документация

### 1. `/api/v1/calculate`

**Метод:** POST

**Описание:** Принимает арифметическое выражение в JSON формате, выполняет вычисления и возвращает результат.

#### Пример запроса:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "(2+2)*2"}'
```

**Пример ответа (успешное выполнение):**

```json
{
  "result": "8.000000"
}
```

### Возможные ошибки

1. **Ошибка 422 (Unprocessable Entity):**

Это ошибка возникает, если передано недопустимое выражение (например, если в нем присутствуют посторонние символы или есть синтаксическая ошибка).

Пример запроса:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "2+x"}'
```

**Пример ответа (ошибка 422):**

```json
{
  "error": "invalid expression"
}
```

2. **Ошибка 500 (Internal Server Error):**

Эта ошибка возникает, если произошло что-то непредвиденное при выполнении вычислений, например, деление на ноль.

Пример запроса:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": 11}'
```

**Пример ответа (ошибка 500):**

```json
{
  "error": "something went wrong"
}
```

---

## Структура проекта

iternal/application
- `application.go` — файл, содержащий код для HTTP-сервера и обработки запросов
- `app_test.go` — файл с тестами для CalcHandler

pkg/calculation
- `calculation.go` — файл, содержащий логику для вычислений выражений
- `errors.go` — ошибки для calculation.go
- `calc_test.go` — тесты для калькулятора

cmd
- `main.go` — файл для запуска сервера


- `go.mod`  — файлы с зависимостями

---
