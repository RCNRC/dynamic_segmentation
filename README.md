# dynamic_segmentation

Сервис, хранящий пользователя и сегменты, в которых он состоит.

## Развёртывание и запуск

При развёртывании укажите открытые порты под PostgreSQL и под сервер с сервисом в файлах "docker-compose.yml", "Makefile" и "configs/config.yml". Для переопределния подключения и/или развёртывания PostgreSQL поменяйте соответствующие параметры в этих же файлах.

Выполнить последовательно команды:
1. `make build`
   - перед первым запуском нужно дождаться создания папки ".database", затем произвести миграцию `make migrate` (для неё потребуется установленный [migrate](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/))
2. `make run`

## API endpoints

1. `/segment/` - добавление/удаление сегмента. Тело запроса должно содержать название сегмента в поле "slag". В теле ответа содержится обязательное поле "success" с значением "true" и дополнительным полем "new_segment" при положительном исходе или с значением "false" и дополнительным полем "error" при отрицательном исходе. Методы:
   - POST - добавляет сегмент.
   - DELETE - удаляет сегмент.
2. `/user/update/` - POST запрос, добавляет пользователя в одни сегменты и удаляет из других. Принимает массив сегментов в поле "segments_add" которые нужно добавить пользователю, массив сегментов в поле "segments_delete" которые нужно удалить у пользователя, id пользователя в поле "id". Каждый сегмент задаётся объектом с полями: обязательным "slag" - названием сегмента и необязательным "ttl" - датой, до которой будет действовать сегмент для данного пользователя (если не установлено, то по умолчанию сегмент действует бесконечно). В теле ответа содержится поле "success" с значением "true" при положительном исходе или с значением "false" при отрицательном исходе.
3. `/user/current/` - GET запрос, позволяет получить набор текущих действующих сегментов пользователя. Тело запроса должно содержать поле "id" - id пользователя. В теле ответа содержатся поля: "success" с значением "true" при положительном исходе или с значением "false" при отрицательном исходе, "error" с значением ошибки, если таковая присутсвует, "id" - id пользователя, "segments" - массив названий сегментов, которые в данный момент активны у пользователя.
4. `/user/history/` - GET запрос для получения ссылки на csv файл с историей добавлений и удалений сегментов. В теле запроса поля: "date_from" - дата с начала которой ищутся записи, "date_to" - дата до которой ищутся записи. Оба поля имеют формат "YYYY-MM", где "YYYY" - год, "MM" - месяц. В теле ответа содержатся поля: "success" с значением "true" при положительном исходе или с значением "false" при отрицательном исходе, "error" с значением ошибки, если таковая присутсвует, "reference" - ссылка на скачивание файла csv с историей добавлений и удалений сегментов.
5. `/user/report/{file_name}` - GET запрос отдающий файл с названием "file_name" из директории "./reports".

### Примеры запросов

Примеры запросов на localhost:8000  
1. Добавление сегмента.
   - Метод: POST
   - Адрес: http://127.0.0.1:8000/segment/

Тело запроса:
```
{
    "slag": "SEGMENT_NAME"
}
```
Положительный ответ:
```
{
    "new_segment": "SEGMENT_NAME",
    "success": "true"
}
```
Отрицательный ответ:
```
{
    "error": "cannot create segment",
    "success": "false"
}
```

2. Удаление сегмента.
   - Метод: DELETE
   - Адрес: http://127.0.0.1:8000/segment/

Тело запроса:
```
{
    "slag": "SEGMENT_NAME"
}
```
Положительный ответ:
```
{
    "deleted_segment": "SEGMENT_NAME",
    "success": "true"
}
```
Отрицательный ответ:
```
{
    "error": "cannot create segment",
    "success": "false"
}
```

3. Добавление пользователя с id = 1000 в одни сегменты и удаление из других.
   - Метод: POST
   - Адрес: http://127.0.0.1:8000/user/update/

Тело запроса:
```
{
    "id": 1000,
    "segments_add": [
        {
            "slag": "SEGMENT_ONE",
            "ttl": "2024-11-21 23:15:45"
        },
        {
            "slag": "SEGMENT_TWO",
        }
    ],
    "segments_delete": [
        {
            "slag": "SEGMENT_THREE"
        }
    ]
}
```
Положительный ответ:
```
{
    "success": "true"
}
```
Отрицательный ответ:
```
{
    "success": "false"
}
```

4. Получение текущих сегментов пользователя с id = 1000.
   - Метод: GET
   - Адрес: http://127.0.0.1:8000/user/current/

Тело запроса:
```
{
    "id": 1000
}
```
Положительный ответ:
```
{
    "success": "true",
    "error": "",
    "id": 1000,
    "segments": [
        "SEGMENT_ONE",
        "SEGMENT_TWO"
    ]
}
```
Отрицательный ответ:
```
{
    "success": "false",
    "error": "cannot update users segments",
    "id": 1000,
    "segments": nil
}
```

5. Получение ссылки на csv файл с записями о добавлении и удалении сегментов у пользователей за период с 2022-08 по 2023-09.
   - Метод: GET
   - Адрес: http://127.0.0.1:8000/user/history/

Тело запроса:
```
{
    "date_from": "2022-08",
    "date_to": "2023-09"
}
```
Положительный ответ:
```
{
    "success": "true",
    "error": "",
    "reference": "http://127.0.0.1:8000/user/report/2022-08:2023-09(2023-08-31 22:21:23).csv"
}
```
Отрицательный ответ:
```
{
    "success": "false",
    "error": "cannot get users history",
    "reference": "http://127.0.0.1:8000/user/report/"
}
```

6. Получение отчёта в виде csv файла с названием "2022-08:2023-09(2023-08-31 22:21:23).csv" с записями о добавлении и удалении сегментов у пользователей за период с 2022-08 по 2023-09.
   - Метод: GET
   - Адрес: http://127.0.0.1:8000/user/report/2023-08:2025-08(2023-08-31 21:54:39).csv

Положительный ответ:
```
1000,SEGMENT_ONE,добавление,2023-08-31T22:13:10.46605Z
1000,SEGMENT_TWO,добавление,2023-08-31T22:13:10.473382Z
1000,SEGMENT_THREE,удаление,2023-08-31T22:13:10.478604Z
```
Отрицательный ответ:
```

```