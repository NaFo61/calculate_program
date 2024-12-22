# calculate_program

# Простой веб-сервис для вычисления арифметических выражений

## Описание
Данный проект представляет собой веб-сервис, позволяющий вычислять арифметические выражения, отправленные пользователем посредством HTTP-запросов.
## Структура проекта

- `cmd/` — точка входа приложения.
- `custom_errors` - спецальные ситуативные случаи ошибок.
- `internal/` — внутренняя логика и модули приложения.
- `pkg/` — вспомогательные пакеты и утилиты.


## Запуск сервиса

1. Установите [Go](https://go.dev/dl/).
2. Склонируйте проект с GitHub:
    ```bash
    git clone https://github.com/NaFo61/calculate_program.git
    ```
3. Перейдите в папку проекта и запустите сервер:
    ```bash
    go run ./cmd/main.go
    ```
4. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

### Альтернативный запуск
Вы можете использовать скрипты для сборки и запуска:
- **Для Linux/MacOS:**
    ```bash
    ./build/build.sh
    ```
- **Для Windows:**
    ```powershell
    .\build\build.bat
    ```

## Эндпоинты

### `POST /api/v1/calculate`

#### Описание
Эндпоинт принимает JSON с математическим выражением.

### Пример запроса с использованием PowerShell

```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
```

## Пример успешного ответа:
```
{
  "result": "6.000000"
}
```

## Пример ошибки 422
Если в выражении присутствуют недопустимые символы, например, знак $, сервер вернёт ошибку с кодом 422:
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "1+2$"}'
```

## Пример ответа:
```
{
"error": "Expression is not valid"
}
```

### Тестирование
Для выполнения тестов используйте следующую команду:


```
go test ./...
```
Примеры запросов с использованием cURL
## Успешный запрос:

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "1 + 2"}'
```
## Запрос с некорректным выражением:

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "invalid"}'
```
## Запрос с неверным JSON:

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{invalid}'
```
## Запрос с выражением, вызывающим ошибку (например, деление на ноль):

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "10/0"}'
```
## Пустой запрос:

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d ''
```
## Запрос без заголовка Content-Type:

```
curl -X POST http://localhost:8080/api/v1/calculate/ \
-d '{"expression": "1 + 2"}'
```
