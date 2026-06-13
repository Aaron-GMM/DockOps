package postgres

import (
	"context"

	"github.com/Aaron-GMM/DockOps/internal/core"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user core.User) error {
	model := UserModel{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*core.User, error) {
	var model UserModel
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &core.User{
		ID:       model.ID,
		Username: model.Username,
		Password: model.Password,
		Role:     model.Role,
	}, nil
}
