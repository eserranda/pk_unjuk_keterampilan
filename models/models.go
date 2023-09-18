package models

import (
	"time"
)

type Produk struct {
	ID         uint      `json:"id" gorm:"primaryKey autoIncrement"`
	NamaProduk string    `json:"nama_produk"`
	Harga      int       `json:"harga"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Produk) TableName() string {
	return "produk"
}

type Invoice struct {
	ID            uint            `json:"id" gorm:"primaryKey autoIncrement"`
	CodeInvoice   string          `json:"code_invoice"`
	TotalHarga    int             `json:"total_harga"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	ProdukInvoice []ProdukInvoice `json:"invoices_produks" gorm:"foreignKey:InvoiceID"`
}

func (Invoice) TableName() string {
	return "invoice"
}

type ProdukInvoice struct {
	ID         uint      `json:"id" gorm:"primaryKey autoIncrement"`
	NamaProduk string    `json:"nama_produk"`
	Harga      int       `json:"harga"`
	InvoiceID  uint      `json:"invoices_id"`
	ProdukID   uint      `json:"produk_id"`
	Jumlah     int       `json:"jumlah"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (ProdukInvoice) TableName() string {
	return "produk_invoice"
}
