package dto

type UserResponse struct {
	Id       uint   `json:"id,omitempty"`
	Name     string `json:"nama,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type BarangResponse struct {
	Id     uint   `json:"id,omitempty"`
	Name   string `json:"nama_barang,omitempty"`
	Brand  string `json:"brand,omitempty"`
	Stok   int    `json:"stok"`
	Satuan string `json:"satuan,omitempty"`
}

type ReqNewBarang struct {
	Id       uint   `json:"id,omitempty" form:"id"`
	Name     string `json:"nama_barang,omitempty" form:"nama_barang"`
	Brand    int    `json:"brand,omitempty" form:"brand"`
	Satuan   int    `json:"satuan,omitempty" form:"satuan"`
	Supplier int    `json:"supplier,omitempty" form:"supplier"`
}

type ReqDate struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
type ReqNewBarangIn struct {
	Id    uint `json:"barang_id"`
	Stock int  `json:"jumlah"`
}

type ReqNewBrandSatuan struct {
	Id   uint   `json:"id,omitempty" form:"id"`
	Name string `json:"nama,omitempty" form:"nama"`
}

type BrandSatuanResponse struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"nama,omitempty"`
}

type ReqNewSupplier struct {
	Id            uint   `json:"id,omitempty" form:"id"`
	Name          string `json:"nama_supplier,omitempty" form:"nama_supplier"`
	Address       string `json:"alamat,omitempty" form:"alamat"`
	ContactNumber string `json:"nomor_kontak,omitempty" form:"nomor_kontak"`
}

type SupplierResponse struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"nama_supplier,omitempty"`
	Address       string `json:"alamat,omitempty"`
	ContactNumber string `json:"nomor_kontak,omitempty"`
}

type RequestResponse struct {
	Id          uint   `json:"id,omitempty"`
	UserName    string `json:"nama_user,omitempty"`
	RequestDate string `json:"tanggal_request,omitempty"`
	Status      string `json:"status"`
	BarangName  string `json:"nama_barang"`
	Amount      int    `json:"jumlah"`
	Satuan      string `json:"satuan"`
}

type BarangInOutResponse struct {
	Id           uint   `json:"id,omitempty"`
	UserName     string `json:"nama_user,omitempty"`
	BarangName   string `json:"nama_barang,omitempty"`
	SupplierName string `json:"nama_supplier,omitempty"`
	Amount       uint   `json:"jumlah_keluar"`
	Satuan       string `json:"satuan,omitempty"`
}

type ReqApproval struct {
	Id     uint `form:"id"`
	Status uint `form:"status"`
}
