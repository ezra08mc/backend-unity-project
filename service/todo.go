package service

import (
	"time"

	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/ezra08mc/backend-unity-project/contract"
	"github.com/ezra08mc/backend-unity-project/database"
	"github.com/ezra08mc/backend-unity-project/dto"
	"gorm.io/gorm"
)

type todoService struct {
	todoRepo contract.TodoRepository
}

func ImplTodoService(todoRepo contract.TodoRepository) contract.TodoService {
	return &todoService{todoRepo: todoRepo}
}

func mapToResponse(t database.Todo, message string) dto.TodoResponse {
	var deletedAt *time.Time
	if t.DeletedAt.Valid {
		deletedAt = &t.DeletedAt.Time
	}

	return dto.TodoResponse{
		Success:     true,
		Message:     message,
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsDone:      t.IsDone,
		CreatedAt:   t.CreatedAt,
		DeletedAt:   deletedAt,
	}
}

func (s *todoService) CreateTodo(userID int, req dto.TodoRequest) (*dto.TodoResponse, error) {
	todo := &database.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
		UserID:      userID,
	}
	if err := s.todoRepo.Create(todo); err != nil {
		return nil, err
	}
	
	resp := mapToResponse(*todo, "Todo created successfully")
	return &resp, nil
}

// Admin methods (Global access)
func (s *todoService) GetAllActive(limit, offset int) ([]dto.TodoResponse, error) {
	todos, err := s.todoRepo.GetAllActive(limit, offset)
	if err != nil {
		return nil, err
	}
	var response []dto.TodoResponse
	for _, t := range todos {
		response = append(response, mapToResponse(t, "Active todo retrieved"))
	}
	return response, nil
}

func (s *todoService) GetAllTrash(limit, offset int) ([]dto.TodoResponse, error) {
	todos, err := s.todoRepo.GetAllTrash(limit, offset)
	if err != nil {
		return nil, err
	}
	var response []dto.TodoResponse
	for _, t := range todos {
		response = append(response, mapToResponse(t, "Trashed todo retrieved"))
	}
	return response, nil
}

func (s *todoService) GetByID(id int) (*dto.TodoResponse, error) {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NotFound("Todo not found")
		}
		return nil, err
	}
	resp := mapToResponse(*todo, "Todo retrieved successfully")
	return &resp, nil
}

func (s *todoService) Update(id int, req dto.TodoRequest) (*dto.TodoResponse, error) {
	todo := &database.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
	}
	if err := s.todoRepo.Update(id, todo); err != nil {
		return nil, err
	}

	updatedTodo, _ := s.todoRepo.GetByID(id)
	resp := mapToResponse(*updatedTodo, "Todo updated successfully")
	return &resp, nil
}

func (s *todoService) SoftDelete(id int) error {
	err := s.todoRepo.SoftDelete(id)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found")
	}
	return err
}

func (s *todoService) Restore(id int) error {
	err := s.todoRepo.Restore(id)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found in trash")
	}
	return err
}

func (s *todoService) PermanentDelete(id int) error {
	err := s.todoRepo.PermanentDelete(id)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found")
	}
	return err
}

// User methods (Scoped access)
func (s *todoService) GetActiveByUserID(userID int, limit, offset int) ([]dto.TodoResponse, error) {
	todos, err := s.todoRepo.GetActiveByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	var response []dto.TodoResponse
	for _, t := range todos {
		response = append(response, mapToResponse(t, "Active todo retrieved"))
	}
	return response, nil
}

func (s *todoService) GetTrashByUserID(userID int, limit, offset int) ([]dto.TodoResponse, error) {
	todos, err := s.todoRepo.GetTrashByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	var response []dto.TodoResponse
	for _, t := range todos {
		response = append(response, mapToResponse(t, "Trashed todo retrieved"))
	}
	return response, nil
}

func (s *todoService) GetByIDAndUserID(id int, userID int) (*dto.TodoResponse, error) {
	todo, err := s.todoRepo.GetByIDAndUserID(id, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NotFound("Todo not found")
		}
		return nil, err
	}
	resp := mapToResponse(*todo, "Todo retrieved successfully")
	return &resp, nil
}

func (s *todoService) UpdateByUserID(id int, userID int, req dto.TodoRequest) (*dto.TodoResponse, error) {
	todo := &database.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
	}
	if err := s.todoRepo.UpdateByUserID(id, userID, todo); err != nil {
		return nil, err
	}

	updatedTodo, _ := s.todoRepo.GetByIDAndUserID(id, userID)
	resp := mapToResponse(*updatedTodo, "Todo updated successfully")
	return &resp, nil
}

func (s *todoService) SoftDeleteByUserID(id int, userID int) error {
	err := s.todoRepo.SoftDeleteByUserID(id, userID)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found")
	}
	return err
}

func (s *todoService) RestoreByUserID(id int, userID int) error {
	err := s.todoRepo.RestoreByUserID(id, userID)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found in trash")
	}
	return err
}

func (s *todoService) PermanentDeleteByUserID(id int, userID int) error {
	err := s.todoRepo.PermanentDeleteByUserID(id, userID)
	if err == gorm.ErrRecordNotFound {
		return errs.NotFound("Todo not found")
	}
	return err
}