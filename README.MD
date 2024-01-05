# GoCore - Golang API Starter Kit & Project Templates

GoCore is a comprehensive Golang API starter kit that provides a solid foundation for building scalable and maintainable APIs using the Go programming language. It comes with project templates, best practices, and a set of tools to streamline the development process.

## Features

- **Modular Structure:** Organize your project into modules for better code organization and maintainability.

- **Middleware Support:** Easily integrate middleware for handling cross-cutting concerns such as authentication, logging, and error handling.

- **Configuration Management:** Manage configuration settings efficiently with support for environment-specific configurations.

- **Database Integration:** Includes support for popular databases like PostgreSQL, MySQL, and SQLite.

- **API Documentation:** Automatic generation of API documentation using Swaggo.

- **Security Best Practices:** Enforce security best practices for handling authentication, authorization, and data protection.

- **Logging and Monitoring:** Integrated logging and monitoring for better visibility into the application's behavior

## Project Structure

root\
├── cmd               // first entry for run command\
├── config            // logic for process config\
├── docs              // auto gen swagger api docs\
├── functions         // function run independent with api service\
├── internal          // main logic of service\
├── pkg               // for reusable\
├── scripts           // bash script for run command\
└── vendor            // dependencies package

## Prerequisites

- [Go](https://go.dev/dl/) ^1.19.1
- [Docker](https://docs.docker.com/install/) ^4.12.0
- [Docker Compose](https://docs.docker.com/compose/install/) ^2.10.2
- [jq](https://stedolan.github.io/jq/download/) ^1.6.0
- [swaggo](https://github.com/swaggo/swag#getting-started) ^1.16.2
- [air](https://github.com/cosmtrek/air#installation) ^1.45.0

## Getting Started

1. Initialize the app for the first time:
   ```bash
   make provision
   ```
2. Generate swagger API docs:
   ```bash
   make specs
   ```
3. Run the development server:
   ```bash
   make dev
   ```

The application runs as an HTTP server at port 8080. You can log in as superadmin to the application by:

```http
POST /login HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "username": "superadmin",
    "password": "superadmin123!@#"
}
```

Then grab the `access_token` for authorization HTTP header:

```http
GET /v1/users HTTP/1.1
Host: localhost:8080
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ...
```

## Deployment
Comming soon

## Installation

```bash
go get -u github.com/vuduongtp/go-core
```
