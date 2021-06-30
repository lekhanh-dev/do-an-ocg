package controller

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []model.Category

	result := repo.Db.Find(&categories)

	categoriesTree := make([]map[string]interface{}, 0)

	for _, category := range categories {
		if category.ParentID == 0 {
			var obj = make(map[string]interface{})
			obj["id"] = category.ID
			obj["name"] = category.Name
			categoriesTree = append(categoriesTree, obj)
		}
	}

	for _, categoryParent := range categoriesTree {
		listCategoryChild := make([]map[string]interface{}, 0)
		for _, category := range categories {
			if categoryParent["id"] == category.ParentID {
				obj := make(map[string]interface{})
				obj["id"] = category.ID
				obj["name"] = category.Name
				listCategoryChild = append(listCategoryChild, obj)
			}
		}
		categoryParent["list_child"] = listCategoryChild
	}

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(categoriesTree)
	}

}

func GetChildCategoriesByParent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	categoriesChild := repo.GetCategoriesById(id)
	categoriesChild = append(categoriesChild, repo.GetCategoryById(id))
	return c.JSON(categoriesChild)
}
