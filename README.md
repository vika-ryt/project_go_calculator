# proect_calculator_go
Сервер для подсчёта арифметических выражений
Пользователь отправляет запрос в формате {
    "expression": "выражение, которое ввёл пользователь"
} , и должен получать ответ в формате :
{
    "result": "результат выражения"
} и кодом 200, в случае ввода данных, которые не соответствуют требованиям выражения должен возращать:
{
    "error": "Expression is not valid"
} и кодом 422, в случае иной ошибки сервер должен возращать:
{
    "error": "Internal server error"
} и код 500
Примеры работы:
Ввод:
{
    "expression": "1 + 1"
}
Вывод: {
    "result": "2"
}, 200
Ввод:
{
    "expression": "1 + a"
} Вывод:

Вывод: {
    "error": "Expression is not valid"
}, 422
Curl
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Вывод:
{"result":"6"}
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2ab + c - d"
}'

Вывод:
{"error": "Expression is not valid"}
Для запуска проекта:
go run ./cmd/main.go