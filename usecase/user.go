package usecase

import (
	"context"
	"net/http"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/model"
	"github.com/aldysp34/sm_padang/repository"
)

type UserUsecase struct {
	request *repository.RequestRepository
	barang  *repository.BarangRepository
}

func NewUserUsecase(request *repository.RequestRepository, barang *repository.BarangRepository) *UserUsecase {
	return &UserUsecase{
		request: request,
		barang:  barang,
	}
}

func (uu *UserUsecase) CreateNewRequest(ctx context.Context, req dto.ReqNewRequest) error {
	newRequest := model.Request{
		UserID:         ctx.Value("id").(uint),
		BarangID:       req.BarangID,
		TotalRequested: req.TotalRequested,
		Status:         3,
	}

	if _, err := uu.request.CreateRequets(newRequest); err != nil {
		return apperr.NewCustomError(http.StatusBadRequest, "failed to create new request")
	}
	return nil
}

func (uu *UserUsecase) GetUserRequest(ctx context.Context) ([]dto.RequestResponse, error) {
	data, err := uu.request.GetRequestByUserID(ctx.Value("id").(uint))
	if err != nil {
		return nil, err
	}

	var res []dto.RequestResponse
	for _, v := range data {
		formattedDate := v.CreatedAt.Format("02012006")
		var status string
		switch v.Status {
		case 1:
			status = "approved"
		case 2:
			status = "rejected"
		case 3:
			status = "pending"
		}
		resp := dto.RequestResponse{
			Id:          v.ID,
			BarangName:  v.Barang.BarangName,
			UserName:    v.User.Username,
			RequestDate: formattedDate,
			Status:      status,
			Amount:      v.TotalRequested,
			Satuan:      v.Barang.Satuan.Satuan,
		}

		res = append(res, resp)
	}

	return res, nil
}

func (uu *UserUsecase) GetAllBarang(ctx context.Context) []dto.BarangResponse {
	data, err := uu.barang.GetAllBarangs()
	if err != nil {
		return nil
	}

	var resp []dto.BarangResponse

	for _, barang := range data {
		newBarang := dto.BarangResponse{
			Id:     barang.ID,
			Brand:  barang.Brand.BrandName,
			Stok:   barang.Total,
			Satuan: barang.Satuan.Satuan,
			Name:   barang.BarangName,
		}
		resp = append(resp, newBarang)
	}

	return resp
}
