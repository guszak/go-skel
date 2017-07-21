package services

import (
	"errors"

	"gitlab.com/guszak/test/conn"
	"gitlab.com/guszak/test/models"
)

// CreateProduct add product
func CreateProduct(p models.Product) (*models.Product, error) {
	db := conn.InitDb()
	defer db.Close()

	if err := db.Create(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// QueryProducts list product
func QueryProducts(offset int64, limit int64) ([]*models.Product, error) {
	db := conn.InitDb()
	defer db.Close()

	var p []*models.Product
	if limit == 0 {
		limit = 10
	}

	err := db.Offset(offset).Limit(limit).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetProduct show product
func GetProduct(id int64) (*models.Product, error) {
	db := conn.InitDb()
	defer db.Close()

	var p models.Product

	if err := db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// UpdateProduct update a product
func UpdateProduct(p models.Product, id int64) (*models.Product, error) {
	db := conn.InitDb()
	defer db.Close()

	var product models.Product

	db.First(&product, id)
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	if p.Description == "" || p.Unity == "" {
		return nil, errors.New("Not implemented")
	}
	product.Description = p.Description
	product.Unity = p.Unity

	if err := db.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// DeleteProduct delete a product
func DeleteProduct(id int64) error {
	db := conn.InitDb()
	defer db.Close()

	var p models.Product
	if err := db.First(&p, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
