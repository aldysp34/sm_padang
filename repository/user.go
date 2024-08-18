package repository

import (
	"context"
	"time"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		if err == gorm.ErrDuplicatedKey {
			return apperr.ErrUserAlreadyExists
		}
	}
	return nil
}

func (ur *UserRepository) GetAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := ur.Db.Preload(clause.Associations).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, req model.User) (model.User, error) {
	var user model.User
	if err := ur.Db.Preload(clause.Associations).First(&user, req.ID).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, req dto.ReqUser) (model.User, error) {
	var user model.User

	if err := ur.Db.Preload(clause.Associations).Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, req model.User) (model.User, error) {

	var userUpdate model.User
	if err := ur.Db.Where("id = ?", req.ID).First(&userUpdate).Error; err != nil {
		return model.User{}, nil
	}

	userUpdate.Username = req.Username
	userUpdate.Nama = req.Nama
	userUpdate.Password = req.Password
	userUpdate.RoleID = req.RoleID
	userUpdate.UpdatedAt = time.Now()
	if err := ur.Db.Save(&userUpdate).Error; err != nil {
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
