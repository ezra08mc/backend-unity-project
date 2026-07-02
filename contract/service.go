package contract

import (
    "github.com/ezra08mc/backend-unity-project/dto"
)

type Service struct {
    Auth AuthService
    Todo TodoService
}

type AuthService interface {
    Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
    Login(req dto.LoginRequest) (*dto.LoginResponse, error)
    GetProfile(userID int) (*dto.ProfileResponse, error)
}

type TodoService interface {
    CreateTodo(userID int, req dto.TodoRequest) (*dto.TodoResponse, error)

    // Admin methods (Global access)
    GetAllActive(limit, offset int) (*dto.TodoListResponse, error)
    GetAllTrash(limit, offset int) (*dto.TodoListResponse, error)
    GetByID(id int) (*dto.TodoResponse, error)
    Update(id int, req dto.TodoRequest) (*dto.TodoResponse, error)
    SoftDelete(id int) error
    Restore(id int) error
    PermanentDelete(id int) error

    // User methods (Scoped access)
    GetActiveByUserID(userID int, limit, offset int) (*dto.TodoListResponse, error)
    GetTrashByUserID(userID int, limit, offset int) (*dto.TodoListResponse, error)
    GetByIDAndUserID(id int, userID int) (*dto.TodoResponse, error)
    UpdateByUserID(id int, userID int, req dto.TodoRequest) (*dto.TodoResponse, error)
    SoftDeleteByUserID(id int, userID int) error
    RestoreByUserID(id int, userID int) error
    PermanentDeleteByUserID(id int, userID int) error
}