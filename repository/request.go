package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
)

type RequestRepository struct {
	Db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{
		Db: db,
	}
}

func (sr *RequestRepository) CreateBarangIn(satuan model.Request) (model.Request, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

func (sr *RequestRepository) GetBarangInByID(id uint) (model.Request, error) {
	var satuan model.Request
	if err := sr.Db.First(&satuan, id).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *RequestRepository) GetAllBarangIns(db *gorm.DB) ([]model.Request, error) {
	var satuans []model.Request
	if err := sr.Db.Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

// Update a satuan
func (sr *RequestRepository) UpdateBarangIn(satuan model.Request) (model.Request, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *RequestRepository) DeleteBarangIn(id uint) error {
	if err := sr.Db.Delete(&model.Request{}, id).Error; err != nil {
		return err
	}
	return nil
}
