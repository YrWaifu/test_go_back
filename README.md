# Магазин мерча

Этот проект представляет собой бэкенд-сервис для внутреннего магазина мерча компании Авито, где сотрудники могут приобретать товары за монеты. Каждый новый сотрудник получает 1000 монет, которые можно использовать для покупки мерча или передачи другим сотрудникам.

## Содержание
- [Описание проекта](#описание-проекта)
- [Функциональность](#функциональность)
- [Установка](#установка)
- [Использование](#использование)
- [API Эндпоинты](#api-эндпоинты)
- [Тестирование](#тестирование)
- [Линтер](#линтер)

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
* 50th percentile - 7ms
* 80th percentile - 15ms
* 90th percentile - 22ms
* 99th percentile - 44ms

Устройство, на котором проходило тестирование имеет следующие характеристики:
* CPU - Ryzen 7 2700x
* RAM - 16GB
* Windows11/WSL


## Линтер

Код проекта был проверен с использованием golangci-lint, и все проверки успешно пройдены, что подтверждает соответствие кода установленным стандартам качества и отсутствие критических ошибок.

**_Линтеры_**:
* **errcheck**: Проверяет, что все возвращаемые ошибки в коде обрабатываются. В конфигурации включена проверка утверждений типов (check-type-assertions: true).
* **goconst**: Ищет повторяющиеся строковые и числовые литералы, которые можно заменить на константы. В конфигурации установлены минимальная длина (min-len: 2) и минимальное количество повторений (min-occurrences: 3).
* **govet**: Выполняет статический анализ кода для выявления потенциальных ошибок и проблем.
* **staticcheck**: Проводит более глубокий анализ кода, выявляя устаревшие конструкции и потенциальные ошибки.
* **gofmt**: Проверяет, что код отформатирован в соответствии с рекомендациями Go.

**_Общие настройки_**
* **disable-all**: true: Отключает все линтеры по умолчанию, чтобы затем включить только те, которые явно указаны в списке enable.
* **issues-exit-code**: 1: Устанавливает код выхода 1, если найдены какие-либо проблемы, что может быть полезно для интеграции с CI/CD системами.