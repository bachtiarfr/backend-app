# Dating App Backend

This repository contains the backend implementation for a Dating Mobile App, similar to Tinder or Bumble, built with Go (Golang). It includes RESTful APIs for user authentication, profile management, and swiping functionalities.

## Features

- User authentication (signup, login)
- Swiping profiles (like, pass)
- Premium features (purchase premium package)
- Swagger documentation
- Unit tests

## Tech Stack

- **Go (Golang)**: Programming language used for backend development.
- **PostgreSQL**: Database used for storing user data and profiles.
- **Docker / Docker Compose**: Containerization and orchestration.
- **Gin**: Web framework used for building the RESTful APIs.
- **GORM**: Object-relational mapping (ORM) library for interfacing with the PostgreSQL database.
- **Swagger**: API documentation tool.
- **Testify**: Library for unit testing in Go.
- **Migrate**: Database schema migration tool.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

- [Docker installation guide](https://docs.docker.com/get-docker/)
- [Docker Compose installation guide](https://docs.docker.com/compose/install/)

## Getting Started

Follow these steps to get the project up and running on your local computer.

1. **Clone the repository:**
   ```bash
   git clone https://github.com/bachtiarfr/backend-app.git
   cd <repository-folder>
   go mod tidy
   
2. **Set up configuration on config.yaml:**
   ```
   database:
    host: "localhost"
    port: "5432"
    user: "postgres"
    password: "yourpassword"
    name: "dating_app"
    
    secretKey: "secret"
   
3. **After you finish set up the configuration of database, do the migraiton:**
   ```bash
   go run cmd/migrate.go

4. **After successfully mirate the db, generate the swager to generate swager endpoint on docs:**
   ```bash
   swag init -g ./cmd/main.go -o ./docs

5. **Run the main app:**
   ```bash
   go run cmd/main.go

6. **Access the application:**
   Once the application are up and running, you can access the application at:
   ```
   http://localhost:8080

7. **API Documentation (Swagger):**
   API documentation is generated using Swagger. You can access the Swagger UI at:
    ```
   http://localhost:8080/swagger/index.html
   
8. **Run the unit testing**
    ```
   go test ./...