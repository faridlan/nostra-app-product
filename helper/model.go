package helper

import (
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {

	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: product.Description,
		Image:       product.Image,
		Category: &web.CategoryResponse{
			Id:   product.Category.Id,
			Name: product.Category.Name,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	productResponses := []web.ProductResponse{}

	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	cateogryResponses := []web.CategoryResponse{}

	for _, category := range categories {
		cateogryResponses = append(cateogryResponses, ToCategoryResponse(category))
	}

	return cateogryResponses
}

//User

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Image:     user.Image,
		RoleId:    user.RoleId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	userResponses := []web.UserResponse{}

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:        role.Id,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	roleResponses := []web.RoleResponse{}

	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}

	return roleResponses
}

func ToLoginResponse(user domain.User) web.LoginResponse {
	return web.LoginResponse{
		User: &web.UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Image:    user.Image,
			RoleId:   user.RoleId,
		},
	}
}
