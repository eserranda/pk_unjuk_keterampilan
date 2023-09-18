package controllers

import (
	"go-prakerja/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetInvoiceByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		invoiceID := c.Param("id")

		var invoice models.Invoice

		if err := db.Preload("ProdukInvoice").First(&invoice, invoiceID).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputData struct {
			CodeInvoice     string `json:"code_invoice"`
			InvoicesProduks []struct {
				ProdukID int `json:"produk_id"`
				Jumlah   int `json:"jumlah"`
			} `json:"invoices_produks"`
			TotalHarga int `json:"total_harga"`
		}

		if err := c.Bind(&inputData); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		totalHarga := 0

		for _, item := range inputData.InvoicesProduks {
			var produk models.Produk
			if err := db.First(&produk, item.ProdukID).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			totalHargaProduk := int(item.Jumlah) * int(produk.Harga)

			totalHarga += totalHargaProduk
		}

		invoice := models.Invoice{
			CodeInvoice: inputData.CodeInvoice,
			TotalHarga:  totalHarga,
		}

		if err := db.Create(&invoice).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		for _, item := range inputData.InvoicesProduks {
			produkInvoice := models.ProdukInvoice{
				InvoiceID: invoice.ID,
				ProdukID:  uint(item.ProdukID),
				Jumlah:    item.Jumlah,
			}
			if err := db.Create(&produkInvoice).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}

		return c.JSON(http.StatusCreated, invoice)
	}
}

func GetInvoiceList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var invoices []models.Invoice

		if err := db.Preload("ProdukInvoice", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, invoice_id, produk_id, jumlah, created_at, updated_at")
		}).Find(&invoices).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		for i := range invoices {
			for j := range invoices[i].ProdukInvoice {
				var produk models.Produk
				if err := db.First(&produk, invoices[i].ProdukInvoice[j].ProdukID).Error; err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}

				invoices[i].ProdukInvoice[j].NamaProduk = produk.NamaProduk
				invoices[i].ProdukInvoice[j].Harga = produk.Harga
			}
		}

		return c.JSON(http.StatusOK, invoices)
	}
}
