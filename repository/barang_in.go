package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.BarangIn{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangInRepository) GetAllBarangIns() ([]model.BarangIn, error) {
	var satuans []model.BarangIn
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

func (sr *BarangInRepository) GetAllBarangInsWithDate(startDate, endDate time.Time) ([]model.BarangIn, error) {
	var satuans []model.BarangIn
	start := startDate.AddDate(0, 0, -1)
	end := endDate.AddDate(0, 0, 1)
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).Where("created_at BETWEEN ? AND ?", start, end).Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BarangInRepository) UpdateBarangIn(satuan model.BarangIn) (model.BarangIn, error) {
	satuan.UpdatedAt = time.Now()
	var data model.BarangIn
	if err := sr.Db.Where("id = ?", satuan.ID).First(&data).Error; err != nil {
		return model.BarangIn{}, err
	}

	data.TotalBarang = satuan.TotalBarang
	data.UpdatedAt = satuan.UpdatedAt
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
