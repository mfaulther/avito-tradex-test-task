# avito-tradex-test-task
Тестовое задание для стажёра Backend в команду Trade Marketing.

* [Перед запуском](#before-launch)
* [Запуск](#launch-app)
* [Реализованные обработчики REST-API](#handlers)
    * [1: POST /stat](#post-stat)
    * [2: GET /stat](#get-stat)
    * [3: DELETE /stat](#del-stat)


## <a name="before-launch"> </a> Перед запуском    

Если будете запускать вручную, без `docker-compose`, то нужно запустить PostgreSQL и указать переменные окружения в файле `.env`


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
### <a name="post-stat"> </a> 1. POST /stats

Метод сохранения статистики

Принимает на вход набор данных в формате JSON со следующими полями:

Поле|Тип|Значение
---|---|---|
date|Дата|Дата события
views|Целое| Кол-во просмотров
clicks|Целое| Кол-во кликов
cost|Вещественное| Цена

Поле `date` обязательное, причем должно быть в формате `YYYY-MM-DD`, иначе вернется
HTTP-ответ со статусом `400 (Bad Request)`.

Поле `cost` должно быть задано с точностью до рубля (что означает не больше 2-х чисел после запятой), иначе
также `Bad Request`.

Остальные поля опциональны.

    POST /stats
    {
        "date": "2010-10-15",
        "views": 45,
        "clicks": 35,
        "cost": 40.23,
    }

В случае успеха возвращается ответ с HTTP статусом `201 Created`

    HTTP 201 {
      "answer": "statistics successfully added !"
    }

Пример запроса с невалидными данными:

    POST /stats
    {
        "date": "15 January, 2007",
        "views": 45,
        "clicks": 35,
        "cost": 40.23,
    }


В этом случае возвращает ответ с HTTP-статусом `400 Bad Request`

    HTTP 400 {
        "error": "Date: must be in format YYYY-MM-DD."
    }

### <a name="get-stat"> </a> 2. GET /stats

Метод показа статистики

Принимает на вход параметры `from` и `to` и возвращает статистикой,
остортированной по дате.

      GET /stats?from=1995-05-10&&to=2007-05-12

Результат:
      
      
      HTTP 200 {
          "data": [
              {
                  "date": "2000-05-05",
                  "views": 0,
                  "clicks": 45,
                  "cost": 100.43,
                  "cpc": 2.2317777,
                  "cpm": 0
              }
              {
                  "date": "2003-05-11",
                  "views": 120,
                  "clicks": 20,
                  "cost": 200.0,
                  "cpc": 10,
                  "cpm": 0.0016666667
              }
         
         
          ]
      }

      


### <a name=del-stat> </a> 3. DELETE /stats

Метод удаления всей сохраненной статистики

Возвращает обычный HTTP-ответ со статусом  `200`


    HTTP 200 {
        "message": "all statistics has been successfully deleted"
    }

