package repository

import (
	"context"
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		Db: db,
	}
}

func (rr *RoleRepository) CreateRole(ctx context.Context, role model.Role) error {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	if err := rr.Db.Create(&role).Error; err != nil {
		return err
	}
	return nil

}

func (rr *RoleRepository) GetRoleByID(ctx context.Context, role model.Role) (model.Role, error) {
	var getRole model.Role
	if err := rr.Db.First(&role, role.ID).Error; err != nil {
		return model.Role{}, err
	}
	return getRole, nil
}

func (rr *RoleRepository) GetAllRoles(ctx context.Context) []model.Role {
	var roles []model.Role
	if err := rr.Db.Find(&roles).Error; err != nil {
		return nil
	}
	return roles
}

func (rr *RoleRepository) UpdateRole(ctx context.Context, req model.Role) (model.Role, error) {
	req.UpdatedAt = time.Now()
	if err := rr.Db.Save(&req).Error; err != nil {
		return model.Role{}, err
	}

	var role model.Role
	if err := rr.Db.First(&role, req.ID).Error; err != nil {
		return model.Role{}, err
	}
	return role, nil
}

func (ur *RoleRepository) DeleteRole(ctx context.Context, req model.Role) error {
	if err := ur.Db.Delete(&model.Role{}, req.ID).Error; err != nil {
		return err
	}
	return nil
}
