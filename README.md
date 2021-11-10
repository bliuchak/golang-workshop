# golang-workshop

Golang workshop to demonstrate how to create REST API.

## State

API endpoints:
- Home - `GET /`
- Health - homework (returns `200` in case Redis connection established othervise returns `503`)
- GetBooks - homework
- GetBook - `GET /book/{id}`
- CreateBook - `POST /book`
- DeleteBook - homework

## Usage

1. Run server

    ```bash
    docker-compose up --build -d
    ```

2. Create book

    ```bash
    curl -XPOST "http://0.0.0.0:3000/book" --data '{"id": "123", "title":"MegaBook"}' -v
    ```

3. Get book

    ```bash
    curl "http://0.0.0.0:3000/book/123"
    ```
