package main

import (
	"fmt"
	"net/http"

	"github.com/faridlan/nostra-api-product/app"
	"github.com/faridlan/nostra-api-product/controller"
	"github.com/faridlan/nostra-api-product/exception"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/middleware"
	"github.com/faridlan/nostra-api-product/repository"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	db := app.NewDatabase()
	validate := validator.New()

	//Upload
	upload := service.NewUploadS3AWS()

	//User
	userRepository := repository.NewUserRepository()
	userService := service.NewAuthService(userRepository, db, validate)
	userController := controller.NewAuthController(userService, upload)

	//User Seeder
	router.POST("/api/users/seeder", userController.CreateMany)
	router.DELETE("/api/users/seeder", userController.DeleteAll)

	//User CRUD
	router.POST("/api/users", userController.Register)
	router.PUT("/api/users/:userId", userController.Update)
	router.GET("/api/users", userController.FindAll)
	router.POST("/api/users/image", userController.UploadIamge)
	router.GET("/api/users/profile/:userId", userController.FindById)
	router.GET("/api/users/profile", userController.Profile)

	//auth user
	router.POST("/api/users/login", userController.Login)

	//Role
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	//Seeder
	router.POST("/api/roles/seeder", roleController.SeederCreate)
	router.DELETE("/api/roles/seeder", roleController.SeederDelete)

	//CRUS
	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)

	//Product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService, upload)

	//CRUD
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	// router.DELETE("/api/products/:productId", productController.Delete)
	router.GET("/api/products/:productId", productController.FindById)
	router.GET("/api/products", productController.FindAll)
	router.POST("/api/products/image", productController.UploadImage)

	//Product Seeder
	router.POST("/api/products/seeder", productController.SeederCreate)
	router.DELETE("/api/products/seeder", productController.SeederDelete)

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	//CRUD
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	//Category Seeder
	router.POST("/api/categories/seeder", categoryController.SeederCreate)
	router.DELETE("/api/categories/seeder", categoryController.SeederDelete)

	router.PanicHandler = exception.ExceptionError

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
		// Handler: router,
	}

	fmt.Println("server running at Port 8080")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
