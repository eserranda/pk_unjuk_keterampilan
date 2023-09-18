package routes

import (
	"go-prakerja/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/create", controllers.CreateProduk(db))
	e.GET("/produk/list", controllers.GetProdukList(db))
	e.DELETE("/produk/:id", controllers.DeleteProdukByID(db))

	e.POST("invoices/create", controllers.CreateInvoice(db))
	e.GET("/invoice/list", controllers.GetInvoiceList(db))

	e.GET("/invoice/:id", controllers.GetInvoiceByID(db))
	e.DELETE("/invoice/delete/:id", controllers.DeleteInvoiceByID(db))

	e.DELETE("/produkinvoice/delete/:id", controllers.DeleteProdukInvoice(db))
	e.GET("/produkinvoice/list", controllers.ShowProdukInvoice(db))
}
