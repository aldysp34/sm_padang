package repository

import (
	"context"
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) CreateNewUser(ctx context.Context, req model.User) error {
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	if err := ur.Db.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, req model.User) (model.User, error) {
	var user model.User
	if err := ur.Db.First(&user, req.ID).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, req model.User) (model.User, error) {
	req.UpdatedAt = time.Now()
	if err := ur.Db.Save(&req).Error; err != nil {
		return model.User{}, err
	}

	var user model.User
	if err := ur.Db.First(&user, req.ID).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, req model.User) error {
	if err := ur.Db.Delete(&model.User{}, req.ID).Error; err != nil {
		return err
	}
	return nil
}
