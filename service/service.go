package service

import "github.com/ezra08mc/backend-unity-project/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		Auth: ImplAuthService(repo.Auth),
		Todo: ImplTodoService(repo.Todo),
	}
}
