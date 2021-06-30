package controller

import (
	"strconv"

	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product

	result := repo.Db.
		Preload("Country").
		Preload("Manufacturer").
		Preload("Category").
		Preload("ProductMedia").
		Find(&products)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func GetProductsByPaging(c *fiber.Ctx) error {
	limit, err1 := strconv.Atoi(c.Query("limit"))
	offset, err2 := strconv.Atoi(c.Query("offset"))
	if err1 != nil {
		return c.JSON(err1)
	}
	if err2 != nil {
		return c.JSON(err2)
	}

	var products []model.Product

	result := repo.Db.
		Preload("Country").
		Preload("Manufacturer").
		Preload("Category").
		Preload("ProductMedia").
		Limit(limit).Offset(offset).
		Find(&products)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func GetProductsByQuery(c *fiber.Ctx) error {
	limit, err1 := strconv.Atoi(c.Query("limit"))
	offset, err2 := strconv.Atoi(c.Query("offset"))
	categoryId, err3 := strconv.Atoi(c.Query("categoryId"))
	priceMin, err4 := strconv.Atoi(c.Query("priceMin"))
	priceMax, err5 := strconv.Atoi(c.Query("priceMax"))

	if err1 != nil {
		return c.JSON(err1)
	}
	if err2 != nil {
		return c.JSON(err2)
	}
	if err3 != nil {
		return c.JSON(err3)
	}
	if err4 != nil {
		return c.JSON(err4)
	}
	if err5 != nil {
		return c.JSON(err5)
	}

	var products []model.Product
	result := repo.Db.
		Preload("Country").
		Preload("Manufacturer").
		Preload("Category").
		Preload("ProductMedia").
		Where("category_id = ? AND price >= ? AND price <= ?", categoryId, priceMin, priceMax).
		Limit(limit).Offset(offset).
		Find(&products)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func CountProductsByQuery(c *fiber.Ctx) error {
	categoryId, err3 := strconv.Atoi(c.Query("categoryId"))
	priceMin, err4 := strconv.Atoi(c.Query("priceMin"))
	priceMax, err5 := strconv.Atoi(c.Query("priceMax"))
	if err3 != nil {
		return c.JSON(err3)
	}
	if err4 != nil {
		return c.JSON(err4)
	}
	if err5 != nil {
		return c.JSON(err5)
	}

	var products []model.Product
	var count int64
	result := repo.Db.
		Where("category_id = ? AND price >= ? AND price <= ?", categoryId, priceMin, priceMax).
		Find(&products).Count(&count)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(count)
	}
}

func CountAllProduct(c *fiber.Ctx) error {
	var count int64
	var products []model.Product
	result := repo.Db.Find(&products).Count(&count)
	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(count)
	}
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(err)
	}

	var product model.Product
	result := repo.Db.
		Preload("Country").
		Preload("Manufacturer").
		Preload("Category").
		Preload("ProductMedia").
		Preload("ProductProperty").
		Where("ID = ? ", id).
		First(&product)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(product)
	}
}
