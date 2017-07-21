package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/guszak/test/models"
	"gitlab.com/guszak/test/services"
)

// CreateProduct add product
func CreateProduct(c *gin.Context) {

	var company = c.MustGet("company").(*models.Company)
	var p models.Product
	c.Bind(&p)

	p.CompanyID = company.ID
	data, err := services.CreateProduct(p)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, data)
	}
}

// QueryProducts list product
func QueryProducts(c *gin.Context) {

	//var company = c.MustGet("company").(*models.Company)

	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := strconv.ParseInt(c.Param("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := services.QueryProducts(offset, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, products)
	}
}

// GetProduct show product
func GetProduct(c *gin.Context) {

	var company = c.MustGet("company").(*models.Company)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := services.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if p.CompanyID != company.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso Negado"})
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// UpdateProduct update a product
func UpdateProduct(c *gin.Context) {

	var company = c.MustGet("company").(*models.Company)
	var product models.Product
	c.Bind(&product)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p, err := services.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if p.CompanyID != company.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso Negado"})
		return
	}

	product.CompanyID = company.ID
	data, err := services.UpdateProduct(product, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// DeleteProduct delete a product
func DeleteProduct(c *gin.Context) {

	var company = c.MustGet("company").(*models.Company)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p, err := services.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if p.CompanyID != company.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso Negado"})
		return
	}

	err = services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	} else {
		c.JSON(http.StatusNoContent, "Product deleted")
	}
}
