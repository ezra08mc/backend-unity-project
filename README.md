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
- 🔐 **Security & Auth** — Robust authentication utilizing **JWT (JSON Web Tokens)** and secure **Password Hashing** (e.g., Bcrypt) for user credentials.
- 📝 **Complete CRUD** — Full capabilities to Create, Read, Update, and Delete todo items.
- 📄 **Pagination** — Efficient data handling using `limit` and `offset` queries for retrieving lists.
- 🗑️ **Soft Delete & Trash** — Safely remove records with soft-delete implementation, allowing users to view or recover items from the trash.
- 📖 **Self-Documenting** — Integrated Swagger/OpenAPI documentation for seamless API exploration and testing.
- 👥 **Role-Based Access** — Distinct access boundaries for regular **Users** (manage own data) and **Admins** (global management).
- 🛡️ **Protection & Limits** — Includes Rate Limiting to prevent abuse and ensure API stability
- 🏗️ **Clean Architecture** — Strict separation of concerns between business logic, data access, and API handling.
- ⚡ **High Performance** — Built with Go for speed, concurrency, and reliability.

## 🛠️ Tech Stack
Built with **Go (Golang)**, **Gin Web Framework**, **GORM** and **PostgreSQL**.

## 📂 Project Structure
- `config/` — Configuration files and environment setup.
- `database/` — Database connection and initialization logic.
- `dto/` — Data Transfer Objects for request and response validation.
- `repository/` — Data access layer for interacting with the database.
- `service/` — Business logic layer.
- `controller/` — HTTP request handlers.
- `contract/` — Interfaces for decoupling layers.
- `main.go` — Application entry point.

## 📋 API Documentation
All endpoints are prefixed with `/api`. Access is strictly controlled based on user roles and data ownership.

### 👤 User Scope (Manage Own Data)

| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/todos` | Create a new todo item | Yes |
| **GET** | `/api/todos` | Get active todos (paginated) | Yes |
| **GET** | `/api/todos/:id` | Get specific own active todo details | Yes |
| **PUT** | `/api/todos/:id` | Update own active todo | Yes |
| **DELETE** | `/api/todos/:id` | Soft delete own todo (Move to trash) | Yes |
| **GET** | `/api/todos/trash` | Get all own soft-deleted todos | Yes |
| **PATCH** | `/api/todos/:id/restore` | Restore todo from trash back to active | Yes |
| **DELETE** | `/api/todos/:id/permanent` | Force delete own todo permanently | Yes |

### 🛡️ Admin Scope (Global Management)

| Method | Endpoint | Description | Auth |
| :--- | :--- | :--- | :--- |
| **GET** | `/api/todos/admin` | Get all users' active todos | Yes (Admin) |
| **GET** | `/api/todos/admin/trash` | Get all users' deleted/trashed todos | Yes (Admin) |
| **GET** | `/api/todos/admin/:id` | Get any specific todo details | Yes (Admin) |
| **PUT** | `/api/todos/admin/:id` | Update any user's todo | Yes (Admin) |
| **PATCH** | `/api/todos/admin/:id/restore` | Restore any user's todo from trash | Yes (Admin) |
| **DELETE** | `/api/todos/admin/:id/permanent` | Force delete any todo permanently | Yes (Admin) |
  
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
go mod download

go mod tidy
```

#### 2.3 Environment Configuration
1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
2. Edit the `.env` file and update the database credentials (`DB_USER`, `DB_PASS`, `DB_NAME`, `DB_HOST`, `DB_PORT`) to match your local PostgreSQL configuration.

`.env`
```env

PORT=8080
IS_PRODUCTION=false
BASE_URL=http://localhost:8080

# Database Configuration
DB_USER=postgres
DB_PASS=password
DB_NAME=todoproject
DB_HOST=localhost
DB_PORT=5432
DB_TIME_ZONE=Asia/Jakarta

# Token Configuration
ACCESS_TOKEN_LIFE_TIME=3600
REFRESH_TOKEN_LIFE_TIME=86400
PRIVATE_KEY=private_key.pem
PUBLIC_KEY=public_key.pem

# Rate Limiting
RATE_LIMIT_RPS=10
RATE_LIMIT_BURST=20
```

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

## 3. API Documentation
This project includes interactive API documentation. Once the server is running, navigate to:
`http://localhost:8080/swagger/index.html`

**Note:** To test protected endpoints in Swagger UI, click the **Authorize** button and enter your token in the format: `Bearer <your_token>`.

### 3.1 Authentication
To access protected endpoints, you must first log in to obtain a JWT token.

Register:
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Your Name",
    "email": "user@example.com",
    "password": "password123",
    "password_confirmation": "password123"
  }'
```
Login:
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Login (Admin):
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
All Todo endpoints require the `Authorization: Bearer <token>` header. Below are the examples categorized by role.

### 👤 User Scope Examples

1. Create a Todo:
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN>" \
  -d '{
    "title": "Learning REST API", 
    "description": "Creating a simple endpoint"
  }'
```

2. Get All Own Todos (with Pagination):
Use limit and offset query parameters.
```bash
curl -X GET "http://localhost:8080/api/todos?limit=10&offset=0" \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

3. Get Specific Own Todo:
curl -X GET http://localhost:8080/api/todos/1 \
  -H "Authorization: Bearer <YOUR_TOKEN>"

4. Update Own Todo:
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN>" \
  -d '{
    "title": "Learning REST API - Advanced", 
    "description": "Adding pagination and soft delete"
  }'
```

5. Soft Delete a Todo (Move to Trash):
```bash
curl -X DELETE http://localhost:8080/api/todos/1 \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

6. Get Trashed Todos:
```bash
curl -X GET http://localhost:8080/api/todos/trash \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

7. Restore Todo from Trash:
```bash
curl -X PATCH http://localhost:8080/api/todos/1/restore \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

8. Permanently Delete Own Todo:
```bash
curl -X DELETE http://localhost:8080/api/todos/1/permanent \
  -H "Authorization: Bearer <YOUR_TOKEN>"
```

### 🛡️ Admin Scope Examples

1. Get All Users' Todos (with Pagination):
```bash
curl -X GET "http://localhost:8080/api/todos/admin?limit=20&offset=0" \
  -H "Authorization: Bearer <ADMIN_TOKEN>"
```

2. Get All Users' Trashed Todos:
```bash
curl -X GET http://localhost:8080/api/todos/admin/trash \
  -H "Authorization: Bearer <ADMIN_TOKEN>"
```

3. Get Specific Todo Detail
```bash
curl -X GET http://localhost:8080/api/todos/admin/1 \
  -H "Authorization: Bearer <ADMIN_TOKEN>"
```

4. Update Any Todo
```bash
curl -X PUT http://localhost:8080/api/todos/admin/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <ADMIN_TOKEN>" \
  -d '{"title": "Updated by Admin", "description": "..."}'
```

5. Restore Any User's Todo:
```bash
curl -X PATCH http://localhost:8080/api/todos/admin/1/restore \
  -H "Authorization: Bearer <ADMIN_TOKEN>"
```

6. Permanent Delete Any User's Todo:
```bash
curl -X DELETE http://localhost:8080/api/todos/admin/1/permanent \
  -H "Authorization: Bearer <ADMIN_TOKEN>"
```

## 5. Troubleshooting

- **Database Connection:** Ensure PostgreSQL is running and your .env credentials are
correct.
- **Missing Keys:** Ensure private_key.pem and public_key.pem exist in the root
directory.
- **401 Unauthorized:** Ensure you are sending the correct Authorization: Bearer
<token> header.
