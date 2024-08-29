package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/model"
	"github.com/aldysp34/sm_padang/repository"
	"github.com/xuri/excelize/v2"
)

type AdminUsecase struct {
	barangIn  *repository.BarangInRepository
	barangOut *repository.BarangOutRepository
	barang    *repository.BarangRepository
	brand     *repository.BrandRepository
	request   *repository.RequestRepository
	role      *repository.RoleRepository
	satuan    *repository.SatuanRepository
	supplier  *repository.SupplierRepository
	user      *repository.UserRepository
}

func NewAdminUsecase(barangIn *repository.BarangInRepository,
	barangOut *repository.BarangOutRepository,
	barang *repository.BarangRepository,
	brand *repository.BrandRepository,
	request *repository.RequestRepository,
	role *repository.RoleRepository,
	satuan *repository.SatuanRepository,
	supplier *repository.SupplierRepository,
	user *repository.UserRepository,
) *AdminUsecase {
	return &AdminUsecase{
		barangIn:  barangIn,
		barangOut: barangOut,
		barang:    barang,
		brand:     brand,
		request:   request,
		role:      role,
		satuan:    satuan,
		supplier:  supplier,
		user:      user,
	}
}

func (au *AdminUsecase) CreateBarangIn(ctx context.Context, req dto.ReqNewBarangIn) error {

	data := model.BarangIn{
		BarangID:    req.Id,
		TotalBarang: req.Stock,
	}

	if _, err := au.barangIn.CreateBarangIn(data); err != nil {

		return err
	}

	barang, err := au.barang.GetBarangByID(req.Id)
	if err != nil {
		return err
	}

	barang.Total += req.Stock
	if _, err := au.barang.UpdateBarang(barang); err != nil {

		return err
	}

	return nil
}
func (au *AdminUsecase) GetAllBarangIn(ctx context.Context) ([]dto.BarangInOutResponse, error) {
	data, err := au.barangIn.GetAllBarangIns()
	if err != nil {
		return nil, err
	}

	var res []dto.BarangInOutResponse
	for _, v := range data {
		formattedDate := v.CreatedAt.Format("02-01-2006")
		newRes := dto.BarangInOutResponse{
			Id:           v.ID,
			SupplierName: v.Barang.Supplier.SupplierName,
			Amount:       uint(v.TotalBarang),
			Brand:        v.Barang.Brand.BrandName,
			BarangName:   v.Barang.BarangName,
			Satuan:       v.Barang.Satuan.Satuan,
			Date:         formattedDate,
		}

		res = append(res, newRes)
	}

	return res, nil
}

func (au *AdminUsecase) DeleteBarangIn(ctx context.Context, req dto.ReqNewBarang) error {
	if err := au.barangIn.DeleteBarangIn(req.Id); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) GetAllBarangOut(ctx context.Context) ([]dto.BarangInOutResponse, error) {
	data, err := au.barangOut.GetAllBarangOuts()
	if err != nil {
		return nil, err
	}

	var res []dto.BarangInOutResponse
	for _, v := range data {
		formattedDate := v.CreatedAt.Format("02-01-2006")
		newRes := dto.BarangInOutResponse{
			Id:           v.ID,
			UserName:     v.Request.User.Nama,
			SupplierName: v.Barang.Supplier.SupplierName,
			Amount:       uint(v.TotalBarang),
			BarangName:   v.Barang.BarangName,
			Satuan:       v.Barang.Satuan.Satuan,
			Brand:        v.Barang.Brand.BrandName,
			Date:         formattedDate,
		}

		res = append(res, newRes)
	}

	return res, nil
}

func (au *AdminUsecase) GetAllBarang(ctx context.Context) ([]dto.BarangResponse, error) {
	data, err := au.barang.GetAllBarangs()
	if err != nil {
		return nil, err
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

	return resp, nil
}

func (au *AdminUsecase) CreateNewBarang(ctx context.Context, req dto.ReqNewBarang) error {
	newBarang := model.Barang{
		BarangName: req.Name,
		SatuanID:   uint(req.Satuan),
		BrandID:    uint(req.Brand),
		Total:      0,
		SupplierID: uint(req.Supplier),
	}

	if _, err := au.barang.CreateBarang(newBarang); err != nil {
		return err
	}

	return nil
}
func (au *AdminUsecase) GetBarangByID(ctx context.Context, req dto.ReqNewBarang) (dto.BarangResponse, error) {
	barang, err := au.barang.GetBarangByID(req.Id)
	if err != nil {
		return dto.BarangResponse{}, err
	}

	resp := dto.BarangResponse{
		Id:     barang.ID,
		Brand:  barang.Brand.BrandName,
		Stok:   barang.Total,
		Satuan: barang.Satuan.Satuan,
		Name:   barang.BarangName,
	}

	return resp, nil
}

func (au *AdminUsecase) EditBarang(ctx context.Context, req dto.ReqNewBarang) error {
	newBarang := model.Barang{
		ID:         req.Id,
		BarangName: req.Name,
		SatuanID:   uint(req.Satuan),
		BrandID:    uint(req.Brand),
		SupplierID: uint(req.Supplier),
	}

	if _, err := au.barang.UpdateBarang(newBarang); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) DeleteBarang(ctx context.Context, req dto.ReqNewBarang) error {

	if err := au.barang.DeleteBarang(req.Id); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) GetAllUser(ctx context.Context) ([]dto.UserResponse, error) {
	data, err := au.user.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	var resp []dto.UserResponse
	for _, user := range data {
		newUser := dto.UserResponse{
			Id:       user.ID,
			Name:     user.Nama,
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role.RoleName,
		}

		resp = append(resp, newUser)
	}

	return resp, nil
}

func (au *AdminUsecase) CreateNewUser(ctx context.Context, req dto.ReqNewUser) error {
	newUser := model.User{
		Nama:     req.Name,
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.Role,
	}
	err := au.user.CreateNewUser(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) EditUser(ctx context.Context, req dto.ReqNewUser) (dto.UserResponse, error) {
	newUser := model.User{
		ID:       req.Id,
		Nama:     req.Name,
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.Role,
	}

	user, err := au.user.UpdateUser(ctx, newUser)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Id:       user.ID,
		Name:     user.Nama,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role.RoleName,
	}, nil
}

func (au *AdminUsecase) DeleteUser(ctx context.Context, req dto.ReqNewUser) error {
	err := au.user.DeleteUser(ctx, model.User{
		ID: req.Id,
	})

	if err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) CreateNewBrand(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	newBrand := model.Brand{
		BrandName: req.Name,
	}

	if _, err := au.brand.CreateBrand(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) GetAllBrand(ctx context.Context) ([]dto.BrandSatuanResponse, error) {
	var resp []dto.BrandSatuanResponse
	data, err := au.brand.GetAllBrands()
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		newBrand := dto.BrandSatuanResponse{
			Id:   v.ID,
			Name: v.BrandName,
		}

		resp = append(resp, newBrand)
	}

	return resp, nil
}

func (au *AdminUsecase) EditBrand(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	newBrand := model.Brand{
		ID:        req.Id,
		BrandName: req.Name,
	}

	if _, err := au.brand.UpdateBrand(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) DeleteBrand(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	err := au.brand.DeleteBrand(req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (au *AdminUsecase) CreateNewSatuan(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	newBrand := model.Satuan{
		Satuan: req.Name,
	}

	if _, err := au.satuan.CreateSatuan(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) GetAllSatuan(ctx context.Context) ([]dto.BrandSatuanResponse, error) {
	var resp []dto.BrandSatuanResponse
	data, err := au.satuan.GetAllSatuans()
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		newBrand := dto.BrandSatuanResponse{
			Id:   v.ID,
			Name: v.Satuan,
		}

		resp = append(resp, newBrand)
	}

	return resp, nil
}

func (au *AdminUsecase) EditSatuan(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	newBrand := model.Satuan{
		ID:     req.Id,
		Satuan: req.Name,
	}

	if _, err := au.satuan.UpdateSatuan(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) DeleteSatuan(ctx context.Context, req dto.ReqNewBrandSatuan) error {
	err := au.satuan.DeleteSatuan(req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (au *AdminUsecase) CreateNewSupplier(ctx context.Context, req dto.ReqNewSupplier) error {
	newBrand := model.Supplier{
		SupplierName:  req.Name,
		Address:       req.Address,
		ContactNumber: req.ContactNumber,
	}

	if _, err := au.supplier.CreateSupplier(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) GetAllSupplier(ctx context.Context) ([]dto.SupplierResponse, error) {
	var resp []dto.SupplierResponse
	data, err := au.supplier.GetAllSuppliers()
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		newBrand := dto.SupplierResponse{
			Id:            v.ID,
			Name:          v.SupplierName,
			Address:       v.Address,
			ContactNumber: v.ContactNumber,
		}

		resp = append(resp, newBrand)
	}

	return resp, nil
}

func (au *AdminUsecase) EditSupplier(ctx context.Context, req dto.ReqNewSupplier) error {
	newBrand := model.Supplier{
		ID:            req.Id,
		SupplierName:  req.Name,
		Address:       req.Address,
		ContactNumber: req.ContactNumber,
	}

	if _, err := au.supplier.UpdateSupplier(newBrand); err != nil {
		return err
	}

	return nil
}

func (au *AdminUsecase) DeleteSupplier(ctx context.Context, req dto.ReqNewSupplier) error {
	err := au.supplier.DeleteSupplier(req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (au *AdminUsecase) ApproveRejectRequest(ctx context.Context, req dto.ReqApproval) error {

	tx := au.request.Db.Begin()
	// Approve Status
	switch req.Status {
	case 1:
		request, err := au.request.GetRequetsyID(req.Id)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Check if the total available barang is less than the requested amount.
		if request.Barang.Total < 1 || request.Barang.Total < request.TotalRequested {
			// Update the request status to 2 (rejected) before rolling back.
			_, err := au.request.UpdateStatus(tx, req.Id, 2)
			if err != nil {
				tx.Rollback()
				return apperr.NewCustomError(http.StatusBadRequest, "error when updating status to rejected")
			}
			if err := tx.Commit().Error; err != nil {
				tx.Rollback()
				return err
			}
			return apperr.NewCustomError(http.StatusBadRequest, "jumlah request kurang dari persediaan, request ditolak")
		}

		barangOut := model.BarangOut{
			BarangID:    request.BarangID,
			TotalBarang: request.TotalRequested,
			RequestID:   request.ID,
		}
		tx, err = au.barangOut.CreateBarangOut(tx, barangOut)
		if err != nil {
			tx.Rollback()
			return apperr.NewCustomError(http.StatusBadRequest, "error when creating barang out")
		}

		tx, err = au.barang.UpdateBarangAmount(tx, request.BarangID, request.TotalRequested)
		if err != nil {
			tx.Rollback()
			return apperr.NewCustomError(http.StatusBadRequest, "error when updating barang amount")
		}

	case 2:
		_, err := au.request.GetRequetsyID(req.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
		// Additional logic for case 2 can go here.
	}

	result, err := au.request.UpdateStatus(tx, req.Id, req.Status)
	if err != nil {
		tx.Rollback()
		return err
	}
	log.Println(result)

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (au *AdminUsecase) GetAllRequest(ctx context.Context) ([]dto.RequestResponse, error) {
	data, err := au.request.GetAllRequest()
	if err != nil {
		return nil, err
	}

	var response []dto.RequestResponse
	for _, v := range data {
		formattedDate := v.CreatedAt.Format("02-01-2006")
		var status string
		switch v.Status {
		case 1:
			status = "approve"
		case 2:
			status = "reject"
		case 3:
			status = "none"
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

		response = append(response, resp)

	}

	return response, nil
}

func (au *AdminUsecase) DownloadXLSX(ctx context.Context, startDate, endDate string) (*excelize.File, error) {
	start, end, err := createDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	barangOuts, err := au.barangOut.GetAllBarangOutsWithDate(*start, *end)
	if err != nil {
		return nil, err
	}

	barangIns, err := au.barangIn.GetAllBarangInsWithDate(*start, *end)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()

	columns := []string{"A", "B", "C", "D", "E", "F", "G"}
	// Create BarangOut sheet
	sheetBarangOut := "BarangOut"
	index, err := f.NewSheet(sheetBarangOut)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)
	if err := f.MergeCell(sheetBarangOut, "A1", "G1"); err != nil {
		return nil, err
	}
	f.SetCellValue(sheetBarangOut, "A1", "BARANG KELUAR")
	header := excelize.Style{
		Font: &excelize.Font{
			Bold:      true,
			Underline: "single",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"},
			Pattern: 1,
		},
	}

	headerStyle, err := f.NewStyle(&header)
	if err != nil {
		return nil, err
	}

	if err := f.SetCellStyle(sheetBarangOut, "A1", "G1", headerStyle); err != nil {
		return nil, err
	}

	f.SetCellValue(sheetBarangOut, "A3", "NPSN")
	f.SetCellValue(sheetBarangOut, "B3", ": 20104455")
	f.SetCellValue(sheetBarangOut, "A4", "Nama Sekolah")
	f.SetCellValue(sheetBarangOut, "B4", ": SMKS Era Pembangunan")
	f.SetCellValue(sheetBarangOut, "A5", "Desa/Kelurahan")
	f.SetCellValue(sheetBarangOut, "B5", ": Jl. Gaga Alastua No. 122 Rt.03/04 Kel. Semanan")
	f.SetCellValue(sheetBarangOut, "A6", "Kecamatan")
	f.SetCellValue(sheetBarangOut, "B6", ": Kalideres")
	f.SetCellValue(sheetBarangOut, "A7", "Provinsi")
	f.SetCellValue(sheetBarangOut, "B7", ": DKI Jakarta")

	f.SetCellValue(sheetBarangOut, "A9", "Tanggal Keluar")
	f.SetCellValue(sheetBarangOut, "B9", "User Request")
	f.SetCellValue(sheetBarangOut, "C9", "Nama Barang")
	f.SetCellValue(sheetBarangOut, "D9", "Jumlah")
	f.SetCellValue(sheetBarangOut, "E9", "Satuan")
	f.SetCellValue(sheetBarangOut, "F9", "Brand")
	f.SetCellValue(sheetBarangOut, "G9", "Nama Supplier")

	color := excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",           // Pattern fill type
			Color:   []string{"#D9D9D9"}, // Fill color
			Pattern: 1,                   // Solid pattern
		},
	}

	// Create the style in Excelize
	colorStyle, err := f.NewStyle(&color)
	if err != nil {
		log.Fatalf("Failed to create style: %v", err)
	}

	// Apply the style to the empty range A10 to G10
	if err := f.SetCellStyle(sheetBarangOut, "A10", "G10", colorStyle); err != nil {
		log.Fatalf("Failed to set cell style: %v", err)
	}

	lastRow := 0

	for i, barang := range barangOuts {
		row := i + 11
		date, err := normalizedUploadDate(barang.CreatedAt.String())
		if err != nil {
			return nil, err
		}
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("A%d", row), date)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("B%d", row), barang.Request.User.Nama)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("C%d", row), barang.Barang.BarangName)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("D%d", row), barang.TotalBarang)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("E%d", row), barang.Barang.Satuan.Satuan)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("F%d", row), barang.Barang.Brand.BrandName)
		f.SetCellValue(sheetBarangOut, fmt.Sprintf("G%d", row), barang.Barang.Supplier.SupplierName)

		lastRow = row
	}
	tableBorderStyle := excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	}

	// Create the style for table borders
	tableBorderStyleID, err := f.NewStyle(&tableBorderStyle)
	if err != nil {
		log.Fatal(err)
	}

	// Apply the table border style from A9 to G{lastRow}
	if err := f.SetCellStyle(sheetBarangOut, "A9", fmt.Sprintf("G%d", lastRow), tableBorderStyleID); err != nil {
		log.Fatal(err)
	}

	f.SetCellValue(sheetBarangOut, fmt.Sprintf("G%d", lastRow+2), fmt.Sprintf("Jakarta, %s", createNowDate()))
	f.SetCellValue(sheetBarangOut, fmt.Sprintf("B%d", lastRow+3), "Mengetahui,")
	f.SetCellValue(sheetBarangOut, fmt.Sprintf("B%d", lastRow+4), "Kepala Sekolah")
	f.SetCellValue(sheetBarangOut, fmt.Sprintf("G%d", lastRow+3), "Petugas")
	f.SetCellValue(sheetBarangOut, fmt.Sprintf("B%d", lastRow+7), "(Nurul Hidayati, S.Pd.)")
	f.SetCellValue(sheetBarangOut, fmt.Sprintf("G%d", lastRow+6), "(Salman Farizi, S.Pd.)")

	for _, col := range columns {
		maxWidth := getMaxContentWidth(f, sheetBarangOut, col, lastRow)
		if err := f.SetColWidth(sheetBarangOut, col, col, maxWidth+2); err != nil {
			log.Fatal(err)
		}
	}

	// Create BarangIn sheet
	sheetBarangIn := "BarangIn"

	in, err := f.NewSheet(sheetBarangIn)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(in)
	if err := f.MergeCell(sheetBarangIn, "A1", "G1"); err != nil {
		return nil, err
	}
	f.SetCellValue(sheetBarangIn, "A1", "BARANG MASUK")

	headerStyleIn, err := f.NewStyle(&header)
	if err != nil {
		return nil, err
	}

	if err := f.SetCellStyle(sheetBarangIn, "A1", "G1", headerStyleIn); err != nil {
		return nil, err
	}

	f.SetCellValue(sheetBarangIn, "A3", "NPSN")
	f.SetCellValue(sheetBarangIn, "B3", ": 20104455")
	f.SetCellValue(sheetBarangIn, "A4", "Nama Sekolah")
	f.SetCellValue(sheetBarangIn, "B4", ": SMKS Era Pembangunan")
	f.SetCellValue(sheetBarangIn, "A5", "Desa/Kelurahan")
	f.SetCellValue(sheetBarangIn, "B5", ": Jl. Gaga Alastua No. 122 Rt.03/04 Kel. Semanan")
	f.SetCellValue(sheetBarangIn, "A6", "Kecamatan")
	f.SetCellValue(sheetBarangIn, "B6", ": Kalideres")
	f.SetCellValue(sheetBarangIn, "A7", "Provinsi")
	f.SetCellValue(sheetBarangIn, "B7", ": DKI Jakarta")

	f.SetCellValue(sheetBarangIn, "A9", "Tanggal Masuk")
	f.SetCellValue(sheetBarangIn, "B9", "Nama Barang")
	f.SetCellValue(sheetBarangIn, "C9", "Jumlah")
	f.SetCellValue(sheetBarangIn, "D9", "Satuan")
	f.SetCellValue(sheetBarangIn, "E9", "Brand")
	f.SetCellValue(sheetBarangIn, "F9", "Stok sisa")
	f.SetCellValue(sheetBarangIn, "G9", "Supplier Name")

	if err := f.SetCellStyle(sheetBarangIn, "A10", "G10", colorStyle); err != nil {
		return nil, err
	}

	lastRow = 0
	for i, barangIn := range barangIns {
		row := i + 11
		date, err := normalizedUploadDate(barangIn.CreatedAt.String())
		if err != nil {
			return nil, err
		}
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("A%d", row), date)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("B%d", row), barangIn.Barang.BarangName)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("C%d", row), barangIn.TotalBarang)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("D%d", row), barangIn.Barang.Satuan.Satuan)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("E%d", row), barangIn.Barang.Brand.BrandName)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("F%d", row), barangIn.Barang.Total)
		f.SetCellValue(sheetBarangIn, fmt.Sprintf("G%d", row), barangIn.Barang.Supplier.SupplierName)
		lastRow = row
	}

	// Apply the table border style from A9 to G{lastRow}
	if err := f.SetCellStyle(sheetBarangIn, "A9", fmt.Sprintf("G%d", lastRow), tableBorderStyleID); err != nil {
		log.Fatal(err)
	}

	f.SetCellValue(sheetBarangIn, fmt.Sprintf("G%d", lastRow+2), fmt.Sprintf("Jakarta, %s", createNowDate()))
	f.SetCellValue(sheetBarangIn, fmt.Sprintf("B%d", lastRow+3), "Mengetahui,")
	f.SetCellValue(sheetBarangIn, fmt.Sprintf("B%d", lastRow+4), "Kepala Sekolah")
	f.SetCellValue(sheetBarangIn, fmt.Sprintf("G%d", lastRow+3), "Petugas")
	f.SetCellValue(sheetBarangIn, fmt.Sprintf("B%d", lastRow+7), "(Nurul Hidayati, S.Pd.)")
	f.SetCellValue(sheetBarangIn, fmt.Sprintf("G%d", lastRow+6), "(Salman Farizi, S.Pd.)")

	for _, col := range columns {
		maxWidth := getMaxContentWidth(f, sheetBarangIn, col, lastRow)
		if err := f.SetColWidth(sheetBarangIn, col, col, maxWidth+2); err != nil {
			log.Fatal(err)
		}
	}

	return f, nil
}

func normalizedUploadDate(uploadDate string) (string, error) {
	timestamp, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", uploadDate)
	if err != nil {
		return "", err
	}

	formattedTimestamp := timestamp.Format("2006-01-02")
	return formattedTimestamp, nil
}

func createDate(startDate, endDate string) (*time.Time, *time.Time, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, nil, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, nil, err
	}

	return &start, &end, nil
}

func createNowDate() string {
	now := time.Now()
	formatTime := now.Format("02-01-2006")
	return formatTime
}

func getMaxContentWidth(f *excelize.File, sheet, col string, lastRow int) float64 {
	maxWidth := 0.0
	for row := 9; row <= lastRow; row++ {
		cell := fmt.Sprintf("%s%d", col, row)
		value, _ := f.GetCellValue(sheet, cell)
		if len(value) > int(maxWidth) {
			maxWidth = float64(len(value))
		}
	}
	return maxWidth * 1.2
}
