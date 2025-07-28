# Notes App

Простое приложение для заметок на Go + PostgreSQL.

## Стек

- Go
- PostgreSQL
- Docker
- Gin
- JWT

## Запуск проекта

```bash
git clone https://github.com/your-user/notes-app.git
cd notes-app
```

## .env файл

```
PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=kuralai
DB_NAME=notes_db
JWT_SECRET=Up#ZDZP7ecml9Ff0cXbA
```

## Запуск через Docker
````
docker compose up --build
````

### База данных
````
CREATE TABLE users (
    id serial PRIMARY KEY,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE TABLE notes (
    id serial NOT NULL PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users(id),
    title text NOT NULL,
    content text NOT NULL
);
````

## API
````
POST /register — регистрация

POST /login — вход

GET /notes — получить все заметки

POST /notes — создать заметку
````

#### Аутентификация через JWT (в заголовке Authorization: Bearer <token>).