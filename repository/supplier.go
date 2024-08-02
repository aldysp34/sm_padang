package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	Db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		Db: db,
	}
}

func (sr *SupplierRepository) CreateBarangIn(satuan model.Supplier) (model.Supplier, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

func (sr *SupplierRepository) GetBarangInByID(id uint) (model.Supplier, error) {
	var satuan model.Supplier
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *SupplierRepository) GetAllBarangIns(db *gorm.DB) ([]model.Supplier, error) {
	var satuans []model.Supplier
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *SupplierRepository) UpdateBarangIn(satuan model.Supplier) (model.Supplier, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *SupplierRepository) DeleteBarangIn(id uint) error {
	if err := sr.Db.Delete(&model.Supplier{}, id).Error; err != nil {
		return err
	}
	return nil
}
