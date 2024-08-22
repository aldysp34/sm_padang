package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BarangOutRepository struct {
	Db *gorm.DB
}

func NewBarangOutRepository(db *gorm.DB) *BarangOutRepository {
	return &BarangOutRepository{
		Db: db,
	}
}

func (sr *BarangOutRepository) CreateBarangOut(tx *gorm.DB, satuan model.BarangOut) (*gorm.DB, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := tx.Create(&satuan).Error; err != nil {
		return tx, err
	}
	return tx, nil
}

func (sr *BarangOutRepository) GetBarangOutByID(id uint) (model.BarangOut, error) {
	var satuan model.BarangOut
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload("Request.User").Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.BarangOut{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangOutRepository) GetAllBarangOuts() ([]model.BarangOut, error) {
	var satuans []model.BarangOut
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload("Request.User").Preload(clause.Associations).Order("created_at DESC").Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

func (sr *BarangOutRepository) GetAllBarangOutsWithDate(startDate, endDate time.Time) ([]model.BarangOut, error) {
	var satuans []model.BarangOut
	start := startDate.AddDate(0, 0, -1)
	end := endDate.AddDate(0, 0, 1)
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload("Request.User").Preload(clause.Associations).Where("created_at BETWEEN ? AND ?", start, end).Order("created_at DESC").Find(&satuans).Error; err != nil {
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
