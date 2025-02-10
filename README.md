# Transaction App

This is a simple API to create accounts and transactions using Go, Gin, and PostgreSQL.

## Prerequisites

### Bare Metal

- Go 1.23 or later
- PostgreSQL
- Git

### Docker

- Docker
- Docker Compose

## Setup

### Bare Metal

1. Clone the repository:

   ```sh
   git clone https://github.com/HeitorMC/pismo-backend.git
   cd pismo-backend
   ```

2. Copy the `.env.example` file to `.env`:

   ```sh
   cp .env.example .env
   ```

3. Update the `.env` if you want to use different configuration.

4. Install dependencies:

   ```sh
   go mod download
   ```

5. Initialize the database:

   ```sh
   go run migrate/migrate.go
   ```

6. Run the application:

   ```sh
   go run cmd/main.go
   ```

### Docker

1. Build and run the application using Docker Compose:

   ```sh
   docker-compose up --build
   ```

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make fmt`: Formats and examines Go source code and reports suspicious constructs.
- `make run`: Generates the swagger documentation, runs the migration, and then runs the application.
- `make build`: Builds both the migration and server, creating executable files named `migration` and `server`.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove both executables and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## API Documentation

The API documentation is available at `/swagger/index.html` when the application is running.

## Endpoints

- `POST /accounts`: Create a new account
- `GET /accounts/{accountId}`: Find an account by its ID
- `POST /transactions`: Create a new transaction

## Environment Variables

- `PORT`: The port on which the application will run (default: 3000)
- `DATABASE_URL`: The URL for connecting to the PostgreSQL database

## License

This project is licensed under the MIT License.
