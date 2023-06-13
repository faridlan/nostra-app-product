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

	//Role
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	//CRUD
	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)

	//User
	userRepository := repository.NewUserRepository()
	userService := service.NewAuthService(userRepository, roleRepository, db, validate)
	userController := controller.NewAuthController(userService, upload)

	//User CRUD
	router.POST("/api/users/register", userController.Register)
	router.PUT("/api/users/:userId", userController.Update)
	router.GET("/api/users", userController.FindAll)
	router.POST("/api/users/image", userController.UploadIamge)
	router.GET("/api/users/:userId", userController.FindById)
	router.GET("/api/users/profile", userController.Profile)

	//auth user
	router.POST("/api/users/login", userController.Login)

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

	//Category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	//CRUD
	router.POST("/api/products/categories", categoryController.Create)
	router.PUT("/api/products/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/products/categories/:categoryId", categoryController.FindById)
	router.GET("/api/products/categories", categoryController.FindAll)

	//SEEDER
	seederService := service.NewSeederService(db, roleRepository, userRepository, categoryRepository, productRepository)
	seederController := controller.NewSeederController(seederService)

	router.POST("/api/seeder", seederController.SaveMany)
	router.DELETE("/api/seeder", seederController.Delete)

	router.PanicHandler = exception.ExceptionError

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.NewAuthMiddleware(router),
		// Handler: router,
	}

	fmt.Println("server running at Port 8080")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
