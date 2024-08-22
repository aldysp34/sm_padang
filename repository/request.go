package repository

import (
	"time"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RequestRepository struct {
	Db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{
		Db: db,
	}
}

func (sr *RequestRepository) CreateRequets(satuan model.Request) (model.Request, error) {
	satuan.CreatedAt = time.Now()
	satuan.UpdatedAt = time.Now()
	satuan.Status = 3
	if err := sr.Db.Create(&satuan).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

func (sr *RequestRepository) GetRequetsyID(id uint) (model.Request, error) {
	var satuan model.Request
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).First(&satuan, id).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

// Get all satuans
func (sr *RequestRepository) GetAllRequest() ([]model.Request, error) {
	var satuans []model.Request
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).Where("status NOT IN (1,2)").Order("created_at DESC").Find(&satuans).Error; err != nil {
		return nil, err
	}
	return satuans, nil
}

func (sr *RequestRepository) GetRequestHistory() ([]model.Request, error) {
	var requests []model.Request
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).Where("status != ?", 0).Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}

// Update a satuan
func (sr *RequestRepository) UpdateRequest(satuan model.Request) (model.Request, error) {
	satuan.UpdatedAt = time.Now()
	if err := sr.Db.Save(&satuan).Error; err != nil {
		return model.Request{}, err
	}
	return satuan, nil
}

// Delete a satuan by ID
func (sr *RequestRepository) DeleteRequets(id uint) error {
	if err := sr.Db.Delete(&model.Request{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (sr *RequestRepository) UpdateStatus(tx *gorm.DB, id, status uint) (model.Request, error) {
	var request model.Request
	if err := tx.First(&request, id).Error; err != nil {
		return model.Request{}, err
	}

	request.Status = status
	tx.Save(&request)

	if err := tx.First(&request, id).Error; err != nil {
		return model.Request{}, err
	}

	return request, nil
}

func (sr *RequestRepository) GetRequestByUserID(id uint) ([]model.Request, error) {
	var requests []model.Request
	if err := sr.Db.Preload("Barang.Supplier").Preload("Barang.Satuan").Preload("Barang.Brand").Preload(clause.Associations).Where("user_id = ?", id).Order("created_at DESC").Find(&requests).Error; err != nil {
		return nil, err
	}

	return requests, nil
}
