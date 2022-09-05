package helmrepo

import (
	"time"

	"gorm.io/gorm"
)

// Repository
// Helm 仓库实体
type Repository struct {
	// 主键
	ID int
	// 仓库名称
	Name string `gorm:"unique"`
	// 仓库地址
	URL string
	// 仓库账号
	Username string
	// 仓库密码
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewRepository() *Repository {
	return &Repository{}
}

func (Repository) TableName() string {
	return "repository"
}
