# Магазин мерча

Этот проект представляет собой бэкенд-сервис для внутреннего магазина мерча компании Авито, где сотрудники могут приобретать товары за монеты. Каждый новый сотрудник получает 1000 монет, которые можно использовать для покупки мерча или передачи другим сотрудникам.

## Содержание
- [Описание проекта](#описание-проекта)
- [Функциональность](#функциональность)
- [Установка](#установка)
- [Использование](#использование)
- [API Эндпоинты](#api-эндпоинты)
- [Тестирование](#тестирование)

## Функциональность
- Покупка мерча.
- Перевод монет другому юзеру.
- Просмотр истории транзакций монет.

## Установка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/YrWaifu/test_go_back
   cd test_go_back
   ```

2. Запустите докер (он создаст базу данных и запустит сервер):
   ```bash
   docker compose -f docker-compose.yaml up server --build
   ```

## Использование

* Доступ к сервису осуществляется по адресу http://localhost:8080.
* Используйте предоставленные API эндпоинты для взаимодействия с сервисом.

## API Эндпоинты
* GET /api/buy/{merch_name}: Купить мерч за монеты. 
* POST /api/auth: Залогинить пользователя или (в случае отсутствия) зарегестрировать нового.
* POST /api/sendCoin: Передать монеты другому сотруднику.
* GET /api/info: Просмотреть историю транзакций монет.

## Тестирование

* Проект покрыт юнит тестами более, чем на 65%.
* Включены интеграционные тесты для сценариев аутентификации, покупки мерча и передачи монет.
* Запуск тестов 
```bash
TEST_SQL_CONNECTION_STRING="postgres://postgres:12345@127.0.0.1:5432/test?sslmode=disable" go test ./...
go tool cover -html .coverage.out 
go tool cover -func .coverage.out 
```
#### Нагрузочное тестирование

Результаты тестирования можно посмотреть в `load_result/load_result.html`. Для нагрузочного тестирования использовалась библиотека `locust` из `python`.

Тестирование проводилось на `/api/info`
* Среднее время ответа при RPS 1090 равно 10.10мс
* SLI успешности ответа - 100%