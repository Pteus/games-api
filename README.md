# **Game Library API**

A RESTful API for managing a game library, built with:

- **Go 1.23.3**
- Repository pattern
- In-memory or PostgreSQL database support
- `goose` for database migrations

This project reflects my effort to follow best practices in both code quality and project structure, while also enhancing my Go skills.

## **API Endpoints**

| Method | Endpoint        | Description             |
|--------|-----------------|-------------------------|
| GET    | `/game`        | Get all games           |
| GET    | `/game/{id}`   | Get a game by ID        |
| POST   | `/game`        | Create a new game       |
| PUT    | `/game/{id}`   | Update a game by ID     |
| DELETE | `/game/{id}`   | Delete a game by ID     |

## **Error Handling**

Errors are logged to the console by default. The API returns standard HTTP status codes and appropriate error messages in the response body. For example:

- **400 Bad Request** for invalid input
- **404 Not Found** when a game doesn't exist
- **500 Internal Server Error** for server-side issues

## **Dependencies**

- `uuid` for generating unique identifiers
  - `go get github.com/google/uuid`      
- `goose` for database migrations
  -  `brew install goose`
- `pq` for PostgreSQL driver
  - `go get github.com/lib/pq`      

## **Future Improvements**

- Add authentication and authorization (JWT, OAuth, etc.)
- Improve error handling and validation
- Add tests for API endpoints
