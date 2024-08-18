package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SatuanRepository struct {
	Db *gorm.DB
}

func NewSatuanRepository(db *gorm.DB) *SatuanRepository {
	return &SatuanRepository{
		Db: db,
	}
}

func (sr *SatuanRepository) CreateSatuan(satuan model.Satuan) (model.Satuan, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Satuan{}, err
	}
	return satuan, nil
}

func (sr *SatuanRepository) GetSatuanByID(id uint) (model.Satuan, error) {
	var satuan model.Satuan
	if err := sr.Db.Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.Satuan{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *SatuanRepository) GetAllSatuans() ([]model.Satuan, error) {
	var satuans []model.Satuan
	if err := sr.Db.Preload(clause.Associations).Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *SatuanRepository) UpdateSatuan(satuan model.Satuan) (model.Satuan, error) {
	satuan.UpdatedAt = time.Now()

	var data model.Satuan
	if err := sr.Db.Where("id = ?", satuan.ID).First(&data).Error; err != nil {
		return model.Satuan{}, err
	}

	data.Satuan = satuan.Satuan
	data.UpdatedAt = satuan.UpdatedAt
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Satuan{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *SatuanRepository) DeleteSatuan(id uint) error {
	if err := sr.Db.Delete(&model.Satuan{}, id).Error; err != nil {
		return err
	}
	return nil
}
