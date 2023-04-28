package main

import (
	"net/http"

	"github.com/faridlan/nostra-api-product/app"
	"github.com/faridlan/nostra-api-product/controller"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/repository"
	"github.com/faridlan/nostra-api-product/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	db := app.NewDatabase()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db)
	productController := controller.NewProductController(productService)

	//CRUD
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	// router.DELETE("/api/products/:productId", productController.Delete)
	router.GET("/api/products/:productId", productController.FindById)
	router.GET("/api/products", productController.FindAll)

	//Seeder
	router.POST("/api/products/seeder", productController.SeederCreate)
	router.DELETE("/api/products/seeder", productController.SeederDelete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
