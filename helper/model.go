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
		CategoryId:  product.CategoryId,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	productResponses := []web.ProductResponse{}

	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
