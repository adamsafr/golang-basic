## Golang Training 1

Golang API which works with https://openweathermap.org API.

Supports: response timeout, request methods validation, panic recovering.

### Configuration

1. `cp .env.dist .env` and update parameters
2. `make run`

### API Endpoints:
1. `GET http://localhost:8081/current-weather?city=Novosibirsk`
2. `GET http://localhost:8081/daily-forecast?city=Novosibirsk`