# avito-tradex-test-task
Тестовое задание для стажёра Backend в команду Trade Marketing.

* [Как запустить](#launch-app)
* [Реализованные обработчики REST-API](#handlers)
    * [1: POST /stat](#post-stats)
    * [2: GET /stat](#get-stats)
    * [3: DELETE /stat](#del-stats)
    


## <a name="launch-app"> </a> Как запустить

* С помощью Go

```
go build ./cmd/main.go
./main
```
Или
```
go run ./cmd/main.go
```

* С помощью docker-compose

```
docker-compose up
```

## <a name="handlers"> </a> Реализованные обработчики REST API
### 1. POST /stats <a name=post-stat> </a>"

Метод сохранения статистики

Принимает на вход набор данных в формате JSON со следующими полями:

Поле|Тип|Значение
---|---|---|
date|Дата|Дата события
views|Целое| Кол-во просмотров
clicks|Целое| Кол-во кликов
cost|Вещественное| Цена

Поле `date` обязательное, причем должно быть в формате `YYYY-MM-DD`, иначе вернется
HTTP-ответ со статусом 400 (Bad Request).

Остальные поля опциональны.

### 2. GET /stats <a name=get-stat> </a>

Метод показа статистики

### 3. DELETE /stats <a name=del-stat> </a>

Метод удаления всей сохраненной статистики



