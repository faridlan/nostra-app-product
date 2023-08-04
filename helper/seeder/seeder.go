package seeder

import (
	"time"

	"github.com/faridlan/nostra-api-product/helper/hash"
	"github.com/faridlan/nostra-api-product/helper/mysql"
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
)

func RoleCreateMany(request []web.RoleCreateReq) []domain.Role {

	roles := []domain.Role{}

	for _, role := range request {
		roleStruct := domain.Role{}
		roleStruct.Name = role.Name
		roleStruct.CreatedAt = time.Now().UnixMilli()
		roles = append(roles, roleStruct)
	}
	return roles

}

func UserCreateMany(request []web.UserCreateReq) []domain.User {

	users := []domain.User{}

	for _, req := range request {
		user := domain.User{}

		imageString := mysql.NewNullString(req.Image)
		hash := hash.HashAndSalt([]byte(req.Password))

		user.Username = req.Username
		user.Password = hash
		user.Email = req.Email
		user.Image = imageString
		user.Role.RoleId = req.RoleId
		user.CreatedAt = time.Now().UnixMilli()

		users = append(users, user)
	}

	return users

}

func CategoryCreateMany(request []web.CategoryCreateReq) []domain.Category {

	categories := []domain.Category{}

	for _, req := range request {
		category := domain.Category{}

		category.Name = req.Name
		category.CreatedAt = time.Now().UnixMilli()

		categories = append(categories, category)
	}

	return categories

}

func ProductCreateMany(request []web.ProductCreateReq) []domain.Product {

	products := []domain.Product{}

	for _, req := range request {
		product := domain.Product{}
		// imageString := mysql.NewNullString(req.Image)

		product.Name = req.Name
		product.Price = req.Price
		product.Quantity = req.Quantity
		product.Description = req.Description
		product.Image = req.Image
		product.Category.CategoryId = req.CategoryId
		product.CreatedAt = time.Now().UnixMilli()

		products = append(products, product)
	}

	return products

}

func ProductImageCreate(request []domain.Product) []domain.ProductImage {
	images := []domain.ProductImage{}

	for _, req := range request {

		for _, x := range req.Image {
			image := domain.ProductImage{}
			image.ProductId = req.Id
			image.ImageUrl = x

			images = append(images, image)
		}

	}

	return images
}
