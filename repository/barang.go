package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := sr.Db.Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.Barang{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *BarangRepository) GetAllBarangs() ([]model.Barang, error) {
	var satuans []model.Barang
	if err := sr.Db.Preload(clause.Associations).Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *BarangRepository) UpdateBarang(req model.Barang) (model.Barang, error) {
	var satuan model.Barang
	if err := sr.Db.Where("id = ?", req.ID).First(&satuan).Error; err != nil {
		return model.Barang{}, err
	}

	satuan.BarangName = req.BarangName
	satuan.BrandID = req.BrandID
	satuan.SatuanID = req.SatuanID
	satuan.SupplierID = req.SupplierID
	satuan.Total = req.Total
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

func (sr *BarangRepository) UpdateBarangAmount(tx *gorm.DB, id uint, jumlah int) (*gorm.DB, error) {
	var data model.Barang
	if err := tx.First(&data, id).Error; err != nil {
		return tx, err
	}

	total := data.Total + jumlah

	if err := tx.Model(&data).Updates(model.Barang{Total: total, UpdatedAt: time.Now()}).Error; err != nil {
		return tx, err
	}
	return tx, nil

}
