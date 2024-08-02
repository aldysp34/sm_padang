package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type BarangOutRepository struct {
	Db *gorm.DB
}

func NewBarangOutRepository(db *gorm.DB) *BarangOutRepository {
	return &BarangOutRepository{
		Db: db,
	}
}

func (sr *BarangOutRepository) CreateBarangOut(satuan model.BarangOut) (model.BarangOut, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.BarangOut{}, err
	}
	return satuan, nil
}

func (sr *BarangOutRepository) GetBarangOutByID(id uint) (model.BarangOut, error) {
	var satuan model.BarangOut
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.BarangOut{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangOutRepository) GetAllBarangOuts(db *gorm.DB) ([]model.BarangOut, error) {
	var satuans []model.BarangOut
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BarangOutRepository) UpdateBarangOut(satuan model.BarangOut) (model.BarangOut, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.BarangOut{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *BarangOutRepository) DeleteBarangOut(id uint) error {
	if err := sr.Db.Delete(&model.BarangOut{}, id).Error; err != nil {
		return err
	}
	return nil
}
