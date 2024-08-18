package dto

type ReqNewRequest struct {
	BarangID       uint `json:"barang_id"`
	TotalRequested int  `json:"total_request"`
}
