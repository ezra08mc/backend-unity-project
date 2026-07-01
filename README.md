<div align="center">
  <h1>Mini Project Backend Unity</h1>
  <h3><b>Secure Todo List API</b></h3>  
  <p>A clean, scalable REST API built with Go, designed to fulfill the requirements of the Backend Development Division mini project.</p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://go.dev/)
  [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
  [![Status](https://img.shields.io/badge/Status-Active-success.svg)]()
</div>

## 🚀 Overview
**Mini Project Backend Unity** is a robust, production-ready REST API built with Go. Following the principles of **Clean Architecture**, this project provides a modular and maintainable foundation for a **Task Management System (To-Do List)**. 

This project is part of the **Back-End Development** division assignment. The expected outcome for candidates is to **understand the Client-Server concept and demonstrate the ability to create basic logic flows** within a structured backend environment.

### 🎯 Key Objectives
1. **Clean Architecture:** Decoupled layers (Controller, Service, Repository) for easy testing and maintenance.
2. **Secure Access:** JWT-based authentication to protect sensitive endpoints from unauthorized access.
3. **Scalable Design:** A modular structure that allows for seamless future feature expansion.

## ✨ Core Features

- 🔐 **Authentication** — Secure endpoints ensuring only authorized users can manage tasks.
- 📝 **Todo List** — Full capabilities to Create, Read, and Delete tasks.
- 🏗️ **Clean Architecture** — Strict separation of concerns between business logic, data access, and API handling.
- ⚡ **High Performance** — Built with Go for speed, concurrency, and reliability.

## 📂 Project Structure

The project follows a clean architecture pattern to ensure maintainability:

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

---
*Created with ❤️ by a candidate member.*
