package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type BarangRepository struct {
	Db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) *BarangRepository {
	return &BarangRepository{
		Db: db,
	}
}

func (sr *BarangRepository) CreateBarang(satuan model.Barang) (model.Barang, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Barang{}, err
	}
	return satuan, nil
}

func (sr *BarangRepository) GetBarangByID(id uint) (model.Barang, error) {
	var satuan model.Barang
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.Barang{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangRepository) GetAllBarangs(db *gorm.DB) ([]model.Barang, error) {
	var satuans []model.Barang
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BarangRepository) UpdateBarang(satuan model.Barang) (model.Barang, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Barang{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *BarangRepository) DeleteBarang(id uint) error {
	if err := sr.Db.Delete(&model.Barang{}, id).Error; err != nil {
		return err
	}
	return nil
}
