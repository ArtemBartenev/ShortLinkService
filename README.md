# Short Link Service

Сервис предоставляет следующие возможности:
* Сокращает и сохраняет переданные ссылки.
* Перенаправляет на исходную ссылку при передаче сокращенной.

## API

### Сокращение и сохранение ссылки.

Успех - 200
<br>
Запрос
```http request
POST http://localhost:8080/v1/short/url
```
Тело запроса
```json
{
  "originalURL": "https://github.com/ArtemBartenev/ShortLinkService"
}
```

Ответ
```json
{"shortURL":  "http://localhost:8000/kfv2"}
```
#### Неуспешные
422
Тело запроса
```json
{
  "originalURL": ""
}
```
Ответ
```text
Field "originalURL" cant be empty.
```

400
```json
{
  "originalURL": "https://github.com/ArtemBartenev/ShortLinkService"
}
```
Ответ
```text
"Original URL already exists."
```

### Перенаправление на исходную ссылку при передаче сокращенной.
```http request
GET http://localhost:8080/v1/short/url?shortURL=http://localhost:8000/kfv2
```
Успешный
```http request
301 Moved Permamently 
```
#### Неуспешные

400
```http request
GET http://localhost:8080/v1/short/url
```
Ответ
```text
ShortURL parameter is missing
```

404
```http request
GET http://localhost:8080/v1/short/url?shortURL=http://localhost:8000/123

```
```text
Original URL not found
```

## Запуск

```text
make build - команда для сборки образа в docker контейнер
make run - запуск сервиса в dokcer-compose вместе с базой данных postgres 
make stop - останавливает работу сервиса и базы данных.
```
