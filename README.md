# Chat API

Простой API для чатов и сообщений на Go с GORM и PostgreSQL, полностью запускается через Docker.

---

## Технологии

- Go 1.21  
- chi (роутер)  
- GORM (ORM)  
- PostgreSQL (Docker)  
- Docker Compose  

---

## Запуск проекта через Docker

1. Клонировать репозиторий:

```bash
git clone https://github.com/pavelluchenkov/chat-api.git
cd chat-api
```

2. Собрать и запустить контейнеры:

```bash
docker-compose up --build
```

3. Приложение будет доступно на:

```
http://localhost:8080
```

---

## Доступные эндпоинты

### 1. Создать чат
**POST /chats**

```bash
curl -X POST http://localhost:8080/chats \
-H "Content-Type: application/json" \
-d '{"title":"Test Chat"}'
```

Пример ответа:

```json
{
  "ID": 1,
  "Title": "Test Chat",
  "CreatedAt": "2026-01-29T12:00:00Z"
}
```

### 2. Отправить сообщение в чат
**POST /chats/{id}/messages**

```bash
curl -X POST http://localhost:8080/chats/1/messages \
-H "Content-Type: application/json" \
-d '{"text":"Hello, world!"}'
```

Пример ответа:

```json
{
  "ID": 1,
  "ChatID": 1,
  "Text": "Hello, world!",
  "CreatedAt": "2026-01-29T12:05:00Z"
}
```

### 3. Получить чат с сообщениями
**GET /chats/{id}?limit=N**

```bash
curl -X GET "http://localhost:8080/chats/1?limit=10"
```

Пример ответа:

```json
{
  "chat": {
    "ID": 1,
    "Title": "Test Chat",
    "CreatedAt": "2026-01-29T12:00:00Z"
  },
  "messages": [
    {
      "ID": 1,
      "ChatID": 1,
      "Text": "Hello, world!",
      "CreatedAt": "2026-01-29T12:05:00Z"
    }
  ]
}
```

- `limit` по умолчанию 20, максимальное 100.  
- Сообщения идут в порядке убывания по дате (`created_at desc`).  

### 4. Удалить чат
**DELETE /chats/{id}**

```bash
curl -i -X DELETE http://localhost:8080/chats/1
```

- Возвращает HTTP статус **204 No Content**, тело ответа пустое.  

---

## Полезные заметки

- Все данные хранятся в Docker volume `pgdata`, так что перезапуск контейнеров **не удаляет данные**.  
- Для «чистого старта» можно удалить volume:

```bash
docker-compose down -v
```

- Проект полностью работает через Docker, **не требует установки Go и PostgreSQL на хосте**.
