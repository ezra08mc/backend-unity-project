package database

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;not null;<-:create"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;uniqueIndex;not null"`
	Password  string    `gorm:"column:password;not null"`
	Role      string    `gorm:"column:role;type:varchar(50);not null;default:'user'"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type Todo struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserID      int            `gorm:"column:user_id;not null"`
	Title       string         `gorm:"column:title;not null"`
	Description string         `gorm:"column:description"`
	IsDone      bool           `gorm:"column:is_done;default:false"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
