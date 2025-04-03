# The_Enricher
Сервис, который будет получать по апи ФИО, из открытых апи обогащать ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в БД
___
# Оглавление
- [Как использовать](#Инструкция)
- [Endpoints](#endpoints)
- [Задание](#задание)
___
# Инструкция 
Чтобы начать работу с приложением:
- Перейдите в корневую папку репозитория
- Запустите docker compose с базой данных
- Создайте .env в app/database/postgres по подобию .env.example
- Запустите приложение `go run main.go` в папке app
- Пользуйтесь с помощью Postman или любого другого подобного инструмента
    - Эндпоинты в следующем разделе
___
# Endpoints
|Method| URL |Discription| Body / Query params| Result|
|------|-----|--------|-------|------|
|**POST**| localhost:8080/create_user | Creates an entry in DB| name: string <br> surname: string <br> patronymic: string|  JSON {<br>name: string <br> surname: string <br> patronymic: string <br> age: int <br> gender: string <br> nationality: string<br>}|
|**PUT**| localhost:8080/update_user | Updates an entry in DB by user_id and field name| user_id: int <br> field_to_update: string <br> new_value: string|  JSON {<br> message: string<br>}|
|**DELETE**| localhost:8080/delete_user | Deletes an entry in DB by user_id| user_id: int |  JSON {<br> message: string<br>}|
|**GET**| localhost:8080/get_users | Returns every entry in DB| _ |  JSON { <br>{<br>"UserID": int,<br>"Name": string,<br>"Surname": string,<br>"Patronymic": string,<br>"Age": int,<br>"Gender": string,<br>"Nationality": string<br>}<br>}|
|**GET**| localhost:8080/get_users_by_filter | Returns every entry in DB by filter| Query: <br> any_filter: any_value (Example: name: Denis)|  JSON { <br>{<br>"UserID": int,<br>"Name": string,<br>"Surname": string,<br>"Patronymic": string,<br>"Age": int,<br>"Gender": string,<br>"Nationality": string<br>}<br>}|
___
# Задание
Реализовать сервис, который будет получать по API ФИО, из открытых API обогащать
ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в
БД. По запросу выдавать инфу о найденных людях. Необходимо реализовать следующее
1. Выставить rest методы
    1. Для получения данных с различными фильтрами и пагинацией
    2. Для удаления по идентификатору
    3. Для изменения сущности
    4. Для добавления новых людей в формате
```
{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" // необязательно
}
```

2. Корректное сообщение обогатить
    1. Возрастом - https://api.agify.io/?name=Dmitriy
    2. Полом - https://api.genderize.io/?name=Dmitriy
    3. Национальностью - https://api.nationalize.io/?name=Dmitriy
3. Обогащенное сообщение положить в БД Postgres (структура БД должна быть создана
путем миграций)
4. Покрыть код debug- и info-логами
5. Вынести конфигурационные данные в .env
