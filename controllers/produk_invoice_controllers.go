package controllers

import (
	"go-prakerja/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteProdukInvoice(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		produkInvoiceID := c.Param("id")

		var produkInvoice models.ProdukInvoice

		if err := db.First(&produkInvoice, produkInvoiceID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "ID tidak di temukan!"})
		}

		if err := db.Delete(&produkInvoice).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, map[string]string{"msg": "Produk Invoice berhasil dihapus", "status": "success"})

	}
}

func ShowProdukInvoice(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var produkInvoices []models.ProdukInvoice

		if err := db.Find(&produkInvoices).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, produkInvoices)
	}
}

func DeleteInvoiceByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		invoiceID := c.Param("id")

		if err := db.Delete(&models.Invoice{}, invoiceID).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}
