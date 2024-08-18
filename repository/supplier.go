package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SupplierRepository struct {
	Db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		Db: db,
	}
}

func (sr *SupplierRepository) CreateSupplier(satuan model.Supplier) (model.Supplier, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

func (sr *SupplierRepository) GetSupplierByID(id uint) (model.Supplier, error) {
	var satuan model.Supplier
	if err := sr.Db.Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *SupplierRepository) GetAllSuppliers() ([]model.Supplier, error) {
	var satuans []model.Supplier
	if err := sr.Db.Preload(clause.Associations).Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *SupplierRepository) UpdateSupplier(satuan model.Supplier) (model.Supplier, error) {
	satuan.UpdatedAt = time.Now()

	var data model.Supplier
	if err := sr.Db.Where("id = ?", satuan.ID).First(&data).Error; err != nil {
		return model.Supplier{}, err
	}

	data.SupplierName = satuan.SupplierName
	data.ContactNumber = satuan.ContactNumber
	data.Address = satuan.Address
	data.UpdatedAt = satuan.UpdatedAt

	if err := sr.Db.Save(&data).Error; err != nil {
		return model.Supplier{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *SupplierRepository) DeleteSupplier(id uint) error {
	if err := sr.Db.Delete(&model.Supplier{}, id).Error; err != nil {
		return err
	}
	return nil
}
