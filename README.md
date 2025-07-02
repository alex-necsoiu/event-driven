## Database Integration

- User and Order services use PostgreSQL by default (see `USER_DB_CONN` and `ORDER_DB_CONN` in configs).
- On startup, each service runs a simple migration to create tables if they do not exist.
- For local dev, you can use the provided `docker/migrate.sql` with Docker Compose or psql.
- Example connection string:
  `postgres://user:pass@localhost:5432/users?sslmode=disable`

- To run migrations manually:
  ```sh
  psql -h localhost -U user -d users -f docker/migrate.sql
  ```

## Environment Configuration

- Each service loads configuration from a `.env` file or `/configs/*.env` (see example files in `configs/` and root `.env.example`).
- Copy `.env.example` to `.env` and adjust values as needed for local development.
- Environment variables control ports, DB connections, event bus URLs, and more.
- Example:
  ```sh
  cp .env.example .env
  cp configs/gateway.env.example configs/gateway.env
  # ...repeat for other services
  ```
- Services use [joho/godotenv](https://github.com/joho/godotenv) to load env files automatically.
