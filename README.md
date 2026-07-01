<div align="center">
  <img src="https://unityunsrat.dev/logo-unity.svg" alt="Logo Unity">
  <h1>Mini Project Backend Unity</h1>
  <h3><b>Secure Todo List API</b></h3>  
  <p>A clean architecture REST API built with Go for the Mini Project from Unity, <br>    
     featuring secure Todo list management with JWT authentication.</p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://go.dev/)
  [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
  [![Status](https://img.shields.io/badge/Status-Active-success.svg)]()
</div>

## 🚀 Overview
**Mini Project Backend Unity** is a robust, production-ready REST API built with Go. Following the principles of **Clean Architecture**, this project provides a modular and maintainable foundation for a **Task Management System (To-Do List)**. 

This project is part of the **Backend Development** division assignment. The expected outcome for candidates is to **understand the Client-Server concept and demonstrate the ability to create basic logic flows** within a structured backend environment.

### 🎯 Key Objectives
1. **Clean Architecture:** Decoupled layers (Controller, Service, Repository) for easy testing and maintenance.
2. **Secure Access:** JWT-based authentication to protect sensitive endpoints from unauthorized access.
3. **Scalable Design:** A modular structure that allows for seamless future feature expansion.

## ✨ Core Features
- 🔐 **Security & Auth** — Robust authentication utilizing **JWT (JSON Web Tokens)**.
- 📝 **Todo List** — Full capabilities to Create, Read, and Delete todo items.
- 📄 **Pagination** — Efficient data handling using `limit` and `offset` queries for retrieving lists.
- 🏗️ **Clean Architecture** — Strict separation of concerns between business logic, data access, and API handling.
- ⚡ **High Performance** — Built with Go for speed, concurrency, and reliability.

## 📋 API Documentation
All endpoints are prefixed with `/api`.

| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/todos` | Create a new todo item | Yes |
| **GET** | `/api/todos` | Get all todos (with pagination) | Yes |
| **DELETE** | `/api/todos/:id` | Delete a todo by ID | Yes |

## 📂 Project Structure
- `config/` — Configuration files and environment setup.
- `database/` — Database connection and initialization logic.
- `dto/` — Data Transfer Objects for request and response validation.
- `repository/` — Data access layer for interacting with the database.
- `service/` — Business logic layer.
- `controller/` — HTTP request handlers.
- `contract/` — Interfaces for decoupling layers.
- `main.go` — Application entry point.

## 🛠️ Tech Stack
Built with **Go (Golang)**, **Gin Web Framework**, and **PostgreSQL**.

## 📖 Guide

### 1. Prerequisites
Ensure your development environment has the following installed:
- **Go** (version 1.21 or higher)
- **PostgreSQL**
- **OpenSSL** (for generating JWT keys)

### 2. Setup Guide

#### 2.1 Clone the Repository
```bash
git clone https://github.com/ezra08mc/backend-unity-project
cd backend-unity-project
```

#### 2.2 Install Dependencies
Initialize and download the required Go modules:
```bash
go mod tidy
```

#### 2.3 Environment Configuration
1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
2. Edit the `.env` file and update the database credentials (`DB_USER`, `DB_PASS`, `DB_NAME`, `DB_HOST`, `DB_PORT`) to match your local PostgreSQL configuration.

#### 2.4 Generate JWT Keys
The application requires RSA key pairs for JWT authentication. Generate them in the root directory:
```bash
openssl genrsa -out private_key.pem 2048
openssl rsa -in private_key.pem -pubout -out public_key.pem
```

#### 2.5 Database Migration & Seeding
Run the initial database migration and seed the default admin account:
```bash
go run main.go migrate
go run main.go seed
```

#### 2.6 Running the Server
Start the application:
```bash
go run main.go
```
The server will be accessible at `http://localhost:8080`.

## 3. API Usage Examples

### 3.1 Authentication
To access protected endpoints, you must first log in to obtain a JWT token.

**Login:**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@unityunsrat.dev",
    "password": "miniprojectbackendunity2026"
  }'
```
*Copy the `token` from the response data.*

### 3.2 Todo List Endpoints
All Todo endpoints require the `Authorization: Bearer <token>` header.

**1. Create a Todo:**
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN>" \
  -d '{
    "title": "Learning REST API", 
    "description": "Creating a simple endpoint"
  }'
```

**2. Get All Todos (with Pagination):**
Use limit and offset query parameters.
```bash
curl -X GET "http://localhost:8080/api/todos?limit=10&offset=0" \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

**3. Delete a Todo:**
```bash
curl -X DELETE http://localhost:8080/api/todos/1 \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

## 5. Troubleshooting
- **Database Connection:** Ensure PostgreSQL is running and your .env credentials are correct.
- **Missing Keys:** Ensure private_key.pem and public_key.pem exist in the root directory.
- **401 Unauthorized:** Ensure you are sending the correct Authorization: Bearer <token> header.
