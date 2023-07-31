package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(product_id, name, price, quantity, description, category_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,UUID_TO_BIN(?),?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.Description, product.Category.CategoryId, product.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name = ?, price = ?, quantity = ?, description = ?, category_id = UUID_TO_BIN(?), updated_at = ? WHERE REPLACE(BIN_TO_UUID(product_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.Description, product.Category.CategoryId, product.UpdatedAt, product.ProductId)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE products WHERE REPLACE(BIN_TO_UUID(product_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ProductId)
	helper.PanicIfError(err)
}
func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	SQL := `SELECT p.id, REPLACE (BIN_TO_UUID(p.product_id),'-','') AS product_id, p.name, p.price, p.quantity, i.image_url, p.description, 
	REPLACE (BIN_TO_UUID(c.category_id), '-', '') AS category_id, c.name, c.created_at, p.created_at
	FROM products AS p 
	INNER JOIN categories AS c ON (c.category_id = p.category_id)
	LEFT JOIN product_images as i ON (p.id = i.product_id)
	WHERE REPLACE(BIN_TO_UUID(p.product_id), '-', '') = ?`

	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	products := make(map[int]domain.Product)

	for rows.Next() {

		var (
			Id              int
			ProductID       string
			Name            string
			Price           int
			Quantity        int
			ImageURL        sql.NullString
			Description     string
			CategoryId      string
			CategoryName    string
			CategoryCreated string
			// CategoryUpdate  string
			CreatedAt int64
			// UpdatedAt       *mysql.NullInt
		)

		// product := domain.Product{}
		// err := rows.Scan(&product.Id, &product.ProductId, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Category.CategoryId, &product.Category.Name, &product.Category.CreatedAt, &product.Category.UpdatedAt, &product.CreatedAt, &product.UpdatedAt)
		err := rows.Scan(&Id, &ProductID, &Name, &Price, &Quantity, &ImageURL, &Description, &CategoryId, &CategoryName, &CategoryCreated, &CreatedAt)
		helper.PanicIfError(err)

		product, ok := products[Id]
		// fmt.Printf("Id: %d image: %s \n", Id, ImageURL.String)

		if ok {
			if ImageURL.Valid {
				product.Image = append(product.Image, ImageURL.String)
			}
			products[Id] = product

		} else {
			product := domain.Product{
				Id:        Id,
				ProductId: ProductID,
				Name:      Name,
				Price:     Price,
				Quantity:  Quantity,
				Category: domain.Category{
					Id:         Id,
					CategoryId: CategoryId,
					Name:       CategoryName,
					CreatedAt:  CreatedAt,
				},
				Description: Description,
				CreatedAt:   CreatedAt,
			}
			if ImageURL.Valid {
				product.Image = append(product.Image, ImageURL.String)
			}
			products[Id] = product
		}

	}

	productModel := domain.Product{}
	for _, product := range products {
		x := domain.Product{
			Id:          product.Id,
			ProductId:   product.ProductId,
			Name:        product.Name,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Image:       product.Image,
			Category:    product.Category,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
		}

		productModel = x
		// fmt.Println(productModel)
	}

	return productModel, nil
}

func (repository *ProductRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `SELECT REPLACE (BIN_TO_UUID(p.product_id),'-','') AS id, p.name, p.price, p.quantity, p.description, 
	REPLACE (BIN_TO_UUID(c.category_id), '-', '') AS category_id, c.name, c.created_at, c.updated_at, p.created_at, p.updated_at
	FROM products AS p 
	INNER JOIN categories AS c ON (c.category_id = p.category_id)
	WHERE p.id = ?`

	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.ProductId, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Category.CategoryId, &product.Category.Name, &product.Category.CreatedAt, &product.Category.UpdatedAt, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `SELECT p.id, REPLACE (BIN_TO_UUID(p.product_id),'-','') AS product_id, p.name, p.price, p.quantity, i.image_url, p.description,
					REPLACE (BIN_TO_UUID(c.category_id), '-', '') AS category_id, c.name, c.created_at, p.created_at
					FROM products AS p
					INNER JOIN categories AS c ON (c.category_id = p.category_id)
					LEFT JOIN product_images as i ON (p.id = i.product_id)
					ORDER BY p.created_at DESC`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	// products := []domain.Product{}
	products := make(map[int]domain.Product)

	for rows.Next() {

		var (
			Id              int
			ProductID       string
			Name            string
			Price           int
			Quantity        int
			ImageURL        sql.NullString
			Description     string
			CategoryId      string
			CategoryName    string
			CategoryCreated string
			// CategoryUpdate  string
			CreatedAt int64
			// UpdatedAt       *mysql.NullInt
		)

		// product := domain.Product{}
		// err := rows.Scan(&product.Id, &product.ProductId, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Category.CategoryId, &product.Category.Name, &product.Category.CreatedAt, &product.Category.UpdatedAt, &product.CreatedAt, &product.UpdatedAt)
		err := rows.Scan(&Id, &ProductID, &Name, &Price, &Quantity, &ImageURL, &Description, &CategoryId, &CategoryName, &CategoryCreated, &CreatedAt)
		helper.PanicIfError(err)

		product, ok := products[Id]
		// fmt.Printf("Id: %d image: %s \n", Id, ImageURL.String)

		if ok {
			if ImageURL.Valid {
				product.Image = append(product.Image, ImageURL.String)
			}
			products[Id] = product

		} else {
			product := domain.Product{
				Id:        Id,
				ProductId: ProductID,
				Name:      Name,
				Price:     Price,
				Quantity:  Quantity,
				Category: domain.Category{
					Id:         Id,
					CategoryId: CategoryId,
					Name:       CategoryName,
					CreatedAt:  CreatedAt,
				},
				Description: Description,
				CreatedAt:   CreatedAt,
			}
			if ImageURL.Valid {
				product.Image = append(product.Image, ImageURL.String)
			}
			products[Id] = product
		}

	}

	result := make([]domain.Product, 0, len(products))

	for _, product := range products {
		// fmt.Printf("Id: %d image: %s", product.Id, product.Images)
		result = append(result, product)
	}

	return result
}

func (repository *ProductRepositoryImpl) FindProductImages(ctx context.Context, tx *sql.Tx, productId int) []domain.ProductImage {
	SQL := "SELECT product_id, image_url FROM product_images where product_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	images := []domain.ProductImage{}

	for rows.Next() {
		image := domain.ProductImage{}
		err := rows.Scan(&image.ProductId, &image.ImageUrl)
		helper.PanicIfError(err)

		images = append(images, image)
	}

	return images
}

func (repository *ProductRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, products []domain.Product) []domain.Product {
	SQL := "INSERT INTO products(product_id, name, price, quantity, description, image, category_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,?,UUID_TO_BIN(?),?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, product := range products {
		result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Quantity, product.Description, product.ImageSingle, product.Category.CategoryId, product.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		product.Id = int(id)
	}

	return products
}

func (repository *ProductRepositoryImpl) SaveImage(ctx context.Context, tx *sql.Tx, products []domain.ProductImage) []domain.ProductImage {

	SQL := "INSERT INTO product_images(product_id, image_url) VALUES (?,?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, image := range products {

		_, err := stmt.ExecContext(ctx, image.ProductId, image.ImageUrl)
		helper.PanicIfError(err)
	}

	return products

}

func (repository *ProductRepositoryImpl) DeleteImage(ctx context.Context, tx *sql.Tx, products domain.Product, product domain.Product) {

	SQL := "DELETE FROM product_images"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)

}

func (repository *ProductRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "TRUNCATE products"

	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
