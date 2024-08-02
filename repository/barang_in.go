package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type BarangInRepository struct {
	Db *gorm.DB
}

func NewBarangInRepository(db *gorm.DB) *BarangInRepository {
	return &BarangInRepository{
		Db: db,
	}
}

func (sr *BarangInRepository) CreateBarangIn(satuan model.BarangIn) (model.BarangIn, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.BarangIn{}, err
	}
	return satuan, nil
}

func (sr *BarangInRepository) GetBarangInByID(id uint) (model.BarangIn, error) {
	var satuan model.BarangIn
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.BarangIn{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangInRepository) GetAllBarangIns(db *gorm.DB) ([]model.BarangIn, error) {
	var satuans []model.BarangIn
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BarangInRepository) UpdateBarangIn(satuan model.BarangIn) (model.BarangIn, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.BarangIn{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *BarangInRepository) DeleteBarangIn(id uint) error {
	if err := sr.Db.Delete(&model.BarangIn{}, id).Error; err != nil {
		return err
	}
	return nil
}
