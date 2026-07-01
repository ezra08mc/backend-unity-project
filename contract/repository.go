package contract

import (
	"github.com/ezra08mc/backend-unity-project/database"
)

type Repository struct {
	Auth AuthRepository
	Todo TodoRepository
}

type AuthRepository interface {
	CreateUser(user *database.User) error
	FindByEmail(email string) (*database.User, error)
	FindByID(id int) (*database.User, error)
}

type TodoRepository interface {
	Create(todo *database.Todo) error

	// Admin methods (Global access)
	GetAllActive(limit, offset int) ([]database.Todo, error)
	GetAllTrash(limit, offset int) ([]database.Todo, error)
	GetByID(id int) (*database.Todo, error)
	Update(id int, todo *database.Todo) error
	SoftDelete(id int) error
	Restore(id int) error
	PermanentDelete(id int) error


	// User methods (Scoped access)
	GetActiveByUserID(userID int, limit, offset int) ([]database.Todo, error)
	GetTrashByUserID(userID int, limit, offset int) ([]database.Todo, error)
	GetByIDAndUserID(id int, userID int) (*database.Todo, error)
	UpdateByUserID(id int, userID int, todo *database.Todo) error
	SoftDeleteByUserID(id int, userID int) error
	RestoreByUserID(id int, userID int) error
	PermanentDeleteByUserID(id int, userID int) error
}
