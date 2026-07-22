# Authentication API

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00AB66)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?logo=postgresql)
![JWT](https://img.shields.io/badge/JWT-Authentication-black)
![Swagger](https://img.shields.io/badge/Swagger-OpenAPI-85EA2D?logo=swagger)

RESTful Authentication API built with **Go (Golang)**, **Gin**, and **PostgreSQL**.

This project implements user authentication and management using **JWT Authentication**, **Refresh Token**, **Layered Architecture**, and **Manual Dependency Injection**.

---

# Features

- User Registration
- User Login
- JWT Authentication
- Refresh Token Authentication
- Authorization Middleware
- User CRUD
- Upload Profile Picture
- Password Hashing (bcrypt)
- PostgreSQL Integration
- Database Migration
- Swagger Documentation
- Makefile Automation
- Layered Architecture
- Manual Dependency Injection

---

# Architecture

```
cmd/main.go                 → Application Entry Point

internal/
├── di/                     → Dependency Injection
├── dto/                    → Request & Response DTO
├── handlers/               → HTTP Handlers
├── lib/                    → JWT & Response Helpers
├── middleware/             → Authorization Middleware
├── models/                 → Models
├── repo/                   → Data Access Layer
└── service/                → Business Logic

migrations/
├── 000001_create_users_table
├── 000002_alter_table_users
└── 000003_create_refresh_tokens
```

---

# Database Schema

```mermaid
erDiagram

USERS {
    bigint id PK
    string username
    string email UK
    string password
    string phone
    string picture
    timestamp created_at
    timestamp updated_at
}

REFRESH_TOKENS {
    bigint id PK
    bigint user_id FK
    string token
    timestamp expires_at
    timestamp created_at
}

USERS ||--o{ REFRESH_TOKENS : owns
```

---

# Tech Stack

| Technology | Description          |
| ---------- | -------------------- |
| Go         | Programming Language |
| Gin        | HTTP Web Framework   |
| PostgreSQL | Database             |
| pgx/v5     | PostgreSQL Driver    |
| bcrypt     | Password Hashing     |
| JWT        | Authentication       |
| godotenv   | Environment Loader   |
| Swaggo     | API Documentation    |

---

# Prerequisites

- Go 1.24+
- PostgreSQL
- golang-migrate

Install migrate

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

---

# Installation

Clone repository

```bash
git clone <repository-url>
```

Go to project

```bash
cd backend1
```

Install dependencies

```bash
go mod tidy
```

---

# Environment Variables

Create `.env`

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/backend1?sslmode=disable

JWT_SECRET=your-secret-key

ACCESS_TOKEN_EXPIRED=15m
REFRESH_TOKEN_EXPIRED=168h
```

---

# Database Migration

Migration Up

```bash
migrate -path migrations \
-database "$DATABASE_URL" up
```

Migration Down

```bash
migrate -path migrations \
-database "$DATABASE_URL" down
```

---

# Make Commands

| Command             | Description                    |
| ------------------- | ------------------------------ |
| `make run`          | Run application                |
| `make swagger`      | Generate Swagger documentation |
| `make migrate-up`   | Run database migration         |
| `make migrate-down` | Rollback database migration    |

> Sesuaikan dengan target yang tersedia di `Makefile` milikmu.

---

# Running Application

Without Makefile

```bash
go run cmd/main.go
```

Using Makefile

```bash
make run
```

Server

```
http://localhost:8080
```

---

# Swagger Documentation

Open Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

# API Endpoints

## Authentication

| Method | Endpoint         | Description          |
| ------ | ---------------- | -------------------- |
| POST   | `/auth/register` | Register new user    |
| POST   | `/auth/login`    | Login user           |
| POST   | `/auth/refresh`  | Refresh access token |

---

## Users

| Method | Endpoint             | Description            |
| ------ | -------------------- | ---------------------- |
| GET    | `/users`             | Get all users          |
| GET    | `/users/{id}`        | Get user by ID         |
| POST   | `/users`             | Create new user        |
| PUT    | `/users/{id}`        | Update user            |
| DELETE | `/users/{id}`        | Delete user            |
| POST   | `/users/{id}/upload` | Upload profile picture |

---

# Authentication Flow

```text
Register
    │
    ▼
Login
    │
    ▼
Access Token + Refresh Token
    │
    ▼
Access Protected Endpoint
    │
    ▼
Access Token Expired
    │
    ▼
POST /auth/refresh
    │
    ▼
New Access Token
```

---

# Project Structure

```
.
├── cmd
├── docs
├── images
├── internal
│   ├── di
│   ├── dto
│   ├── handlers
│   ├── lib
│   ├── middleware
│   ├── models
│   ├── repo
│   └── service
├── migrations
├── uploads
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

---

# API Testing Results

## Register

Register endpoint tested using REST Client.

![Register Success](images/register-success.png)

---

## Login

Login endpoint tested using REST Client.

![Login Success](images/login-success.png)

---

# Security

- Password hashed using bcrypt
- JWT Access Token Authentication
- Refresh Token Authentication
- Authorization Middleware
- Protected API Endpoints

---

# Future Improvements

- Docker & Docker Compose
- Unit Testing
- Role-Based Access Control (RBAC)
- Email Verification
- Forgot Password
- Rate Limiting
