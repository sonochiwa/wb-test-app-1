# Запуск

Для запуска приложения необходимо выполнить команду:

```bash
go build -o ./bin cmd/main.go
go run bin/main.go
```

# Описание

Во многих местах приложение было осознанно упрощено, чтобы оставить
как можно меньше кода не относящегося к решению задачи, и уделить больше
внимания проблемам, на проверку понимания которых нацелено тестовое задание.

# Документация

API приложения поддерживает следующие методы:

- `/process-numbers` [POST]  
  Позволяет отправить множество чисел (в количестве от 1 до 10К) на
  обработку функцией GetResult  
  request:
  ```json
  {
    "numbers": [1, 2, 3]
  }
  ```
  response (успешная попытка):
  ```json
  {
    "results": {
      "1": 1,
      "2": 2,
      "3": 3
    }
  }
  ```
  response (не успели в timeout):
    ```json
  {
    "message": "Для некоторых из отправленных чисел подсчет незавершен и находится в процессе"
  }
  ```