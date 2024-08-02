package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type BrandRepository struct {
	Db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{
		Db: db,
	}
}

func (sr *BrandRepository) CreateBrand(satuan model.Brand) (model.Brand, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Brand{}, err
	}
	return satuan, nil
}

func (sr *BrandRepository) GetBrandByID(id uint) (model.Brand, error) {
	var satuan model.Brand
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.Brand{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BrandRepository) GetAllBrands(db *gorm.DB) ([]model.Brand, error) {
	var satuans []model.Brand
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BrandRepository) UpdateBrand(satuan model.Brand) (model.Brand, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Brand{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *BrandRepository) DeleteBrand(id uint) error {
	if err := sr.Db.Delete(&model.Brand{}, id).Error; err != nil {
		return err
	}
	return nil
}
