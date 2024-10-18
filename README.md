# Go REST API Template

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

This repository provides a minimalist template for building REST APIs in Go, utilizing a minimal number of external packages while maintaining essential functionality. It is designed to be lightweight, simple, and easy to extend as needed, using only a few necessary libraries.

## Key Features

- **Minimal External Dependencies**: 
  - [Gorilla Mux](https://github.com/gorilla/mux) - for routing.
  - [GoDotEnv](https://github.com/joho/godotenv) - for environment variable management.
  - [Google UUID](https://github.com/google/uuid) - for generating UUIDs.
- **Core Go Packages**: Leverages Go's built-in libraries (net/http, encoding/json, etc.).
- **Database-Ready**: Includes basic setup for database integration (e.g., PostgreSQL, MySQL, MongoDB).
- **RESTful Design**: Follows REST API principles for clean, stateless server-client communication.
- **Custom Logging**: Simple logging system to track API requests and errors.
- **Custom Error Handling**: Robust error-handling mechanism to handle common API errors consistently.

## Requirements

- Go 1.18 or higher
- Database (Optional: configure in `.env` file)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Beretta350/golang-rest-template.git
cd golang-rest-template
```

### Databases already configured

This branches has alterations made to support some databases:

- [MySQL](https://github.com/Beretta350/golang-rest-template/tree/mysql)
- [Postgres](https://github.com/Beretta350/golang-rest-template/tree/postgres)
- [MongoDB](https://github.com/Beretta350/golang-rest-template/tree/mongodb)


### Environment Configuration

This project uses [GoDotEnv](https://github.com/joho/godotenv) for managing environment variables. See the `local.env` file.

Set your database connection string or any other configuration options in the `.env` file.

### Running the API

Once you have set up your environment and make the changes to accept some go to deployments and run the following command to start the server:

```bash
docker compose -f ./<compose-file-name> up -d --build
```

The API will start at `http://localhost:8080` by default.

### Example Endpoints

Here are some basic endpoints provided by the template:

- `GET /users`: Fetch all resources.
- `GET /users/{id}`: Fetch a single resource by ID.
- `POST /users`: Create a new resource.
- `PUT /users/{id}`: Update an existing resource.
- `DELETE /users/{id}`: Delete a resource.

### Example Request

```bash
curl -X GET http://localhost:8080/users
```
## External Packages

- **[Gorilla Mux](https://github.com/gorilla/mux)**: A powerful URL router and dispatcher.
- **[GoDotEnv](https://github.com/joho/godotenv)**: Loads environment variables from `.env` files.
- **[Google UUID](https://github.com/google/uuid)**: For generating universally unique identifiers (UUIDs).

## Contributing

Feel free to submit issues, fork the repository, and open pull requests if you want to contribute to improving this template.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
