# Effective_Mobile


# Тайм-трекер

## Функционал проекта

Проект представляет собой тайм-трекер, который позволяет пользователям отслеживать время, затраченное на задачи. Он предоставляет следующий набор функций:

- Учет и отображение времени, затраченного на задачи.
- Создание пользователя.
- Создание новой задачи для отслеживания времени.
- Фильтрация параметров пользователя по различным параметрам (имя,адрес статус).
- Получение списка всех задач с возможностью детального просмотра каждой задачи.
- Получение общей информации о времени, затраченном пользователем на задачи за определенный период.

Для хранения данных используется PostgreSQL.

## Запуск

 **Запуск проекта:** 
    ```
    make run
    ```
    или 
       ```
    go run cmd/server/main.go --config=./config/local.yaml
    ```
 **Доступ к API:** 
    После запуска проекта API будет доступно по адресу `http://localhost:8080`.

 **Доступ к Swagger:** 
       После запуска проекта swagger будет по адресу `http://localhost:8081`

## Использование API

После запуска проекта вы можете использовать любой HTTP клиент (например, cURL, Postman) для взаимодействия с API. Ниже приведены примеры использования основных эндпоинтов:

- Получить список: 
    ```http
    GET http://localhost:8080//paginated
    ```

- Создать нового пользователя: 
    ```http
    POST http://localhost:8080/users
    ```
    пример
    ```example
    Post http://localhost:8080/worklogs
    {
    "surname": "Иванов",
    "name": "Иван",
    "patronymic": "Иванович",
    "address": "г. Москва, ул. Ленина, д. 10, кв. 5",
    "passport_series": "1234",
    "passport_number": "567890"
    }
    ```
- Получить информацию о задаче по ID: 
    ```http
    GET http://localhost:8080/users/{id}
    ```

- Обновить информацию о задаче: 
    ```http
    PUT http://localhost:8080/users/{id}
    ```

- Удалить задачу: 
    ```http
    DELETE http://localhost:8080/users/{id}
    ```
    
- Получить список задач с поддержкой пагинации:
    ```http
    GET /users/paginated/{page}&{limit}&{status}
    ```
   
    ```example
    GET /users/paginated?page=1&limit=10&status=true
    ```
    
- Получить список задач с фильтрацией по дате и статусу:
    ```http
    GET /users/filtered?column1={name}&column2={Jhon}
    ```
    ```example
    GET /users/filtered?column1={name}&column2={Jhon}
    ```
    
- Запустить таймер:
    ```http
    Post /users/filtered?column1={name}&column2={Jhon}
    ```

    пример

    ```
    Post http://localhost:8080/worklogs
    {
    "userID": 1,
    "description": "Working on feature X"
    }
    ```

- Остановить таймер:
    ```http
    Post /worklogs/{user_id}
    ```

    пример

    ```
    Post http://localhost:8080/worklogs
    {
     "worklog_id": 1
    }
    ```
- Вывести результат:
    ```http
    GET /worklogs/user/{user_id}
    ```

    пример

    ```example
    Get http://localhost:8080/worklogs/user/1

    ```
    

## Не успел

В планах было сделать шифрование паспортных данных, но не успел к сожалению,стоило обеъдинить запросы к пользователю в один эндпоинт, а не делать каждый отедльно(Get user, paginated, filter) так же в качестве бд было бы предпочтительнее использовать MongoDB

