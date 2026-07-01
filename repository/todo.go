package repository

import (
	"github.com/ezra08mc/backend-unity-project/contract"
	"github.com/ezra08mc/backend-unity-project/database"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func ImplTodoRepository(db *gorm.DB) contract.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *database.Todo) error {
	return r.db.Create(todo).Error
}

// Admin methods (Global access)
func (r *todoRepository) GetAllActive(limit, offset int) ([]database.Todo, error) {
	var todos []database.Todo
	err := r.db.Limit(limit).Offset(offset).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetAllTrash(limit, offset int) ([]database.Todo, error) {
	var todos []database.Todo
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Limit(limit).Offset(offset).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByID(id int) (*database.Todo, error) {
	var todo database.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(id int, todo *database.Todo) error {
	return r.db.Model(&database.Todo{}).Where("id = ?", id).Updates(todo).Error
}

func (r *todoRepository) SoftDelete(id int) error {
	result := r.db.Delete(&database.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) Restore(id int) error {
	result := r.db.Unscoped().Model(&database.Todo{}).Where("id = ?", id).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) PermanentDelete(id int) error {
	result := r.db.Unscoped().Delete(&database.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// User methods (Scoped access)
func (r *todoRepository) GetActiveByUserID(userID int, limit, offset int) ([]database.Todo, error) {
	var todos []database.Todo
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetTrashByUserID(userID int, limit, offset int) ([]database.Todo, error) {
	var todos []database.Todo
	err := r.db.Unscoped().Where("user_id = ? AND deleted_at IS NOT NULL", userID).Limit(limit).Offset(offset).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByIDAndUserID(id int, userID int) (*database.Todo, error) {
	var todo database.Todo
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error
	return &todo, err
}

func (r *todoRepository) UpdateByUserID(id int, userID int, todo *database.Todo) error {
	result := r.db.Model(&database.Todo{}).Where("id = ? AND user_id = ?", id, userID).Updates(todo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) SoftDeleteByUserID(id int, userID int) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&database.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) RestoreByUserID(id int, userID int) error {
	result := r.db.Unscoped().Model(&database.Todo{}).Where("id = ? AND user_id = ?", id, userID).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) PermanentDeleteByUserID(id int, userID int) error {
	result := r.db.Unscoped().Where("id = ? AND user_id = ?", id, userID).Delete(&database.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}