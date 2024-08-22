package model

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Username  string `gorm:"unique"`
	Password  string
	RoleID    uint `gorm:"index"`
	Role      Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Role represents a user role.
type Role struct {
	ID        uint `gorm:"primaryKey"`
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Satuan represents a unit of measurement.
type Satuan struct {
	ID        uint `gorm:"primaryKey"`
	Satuan    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Brand represents a brand of goods.
type Brand struct {
	ID        uint `gorm:"primaryKey"`
	BrandName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Barang represents an item in the inventory.
type Barang struct {
	ID         uint `gorm:"primaryKey"`
	BarangName string
	SatuanID   uint   `gorm:"index"`
	Satuan     Satuan `gorm:"foreignKey:SatuanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BrandID    uint   `gorm:"index"`
	Brand      Brand  `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total      int
	SupplierID uint     `gorm:"index"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// BarangIn represents incoming goods.
type BarangIn struct {
	ID          uint   `gorm:"primaryKey"`
	BarangID    uint   `gorm:"index"`
	Barang      Barang `gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalBarang int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BarangOut represents outgoing goods.
type BarangOut struct {
	ID          uint   `gorm:"primaryKey"`
	BarangID    uint   `gorm:"index"`
	Barang      Barang `gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalBarang int
	RequestID   uint    `gorm:"index"`
	Request     Request `gorm:"foreignKey:RequestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Supplier represents a supplier.
type Supplier struct {
	ID            uint `gorm:"primaryKey"`
	SupplierName  string
	Address       string
	ContactNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Request represents a request made by a user.
type Request struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"index"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BarangID       uint   `gorm:"index"`
	Barang         Barang `gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalRequested int
	Status         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
