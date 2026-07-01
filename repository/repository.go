package repository

import (
	"github.com/ezra08mc/backend-unity-project/contract"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		Auth: ImplAuthRepository(db),
		Todo: ImplTodoRepository(db),
	}
}
