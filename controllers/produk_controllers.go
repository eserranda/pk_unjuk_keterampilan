package controllers

import (
	"go-prakerja/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProdukList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var produks []models.Produk
		db.Find(&produks)
		return c.JSON(http.StatusOK, produks)
	}
}

func DeleteProdukByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		db.Where("id = ?", id).Delete(&models.Produk{})

		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Data produk berhasil dihapus"})
	}
}

func CreateProduk(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var produk models.Produk
		if err := c.Bind(&produk); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		db.Create(&produk)

		return c.JSON(http.StatusCreated, produk)
	}
}
