package routes

import (
	"github.com/TechMaster/golang/15GoMySQL/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {
	manufacturerAPI := app.Group("/api/manufacturer")
	routeManufacturer(&manufacturerAPI) //http://localhost:3000/api/book

	productAPI := app.Group("/api/product")
	routeProduct(&productAPI)

	categoryAPI := app.Group("/api/category")
	routeCategory(&categoryAPI)
}

func routeManufacturer(router *fiber.Router) {
	//Return all manufacturers
	(*router).Get("/", controller.GetManufacturers) //Liệt kê
}

func routeProduct(router *fiber.Router) {
	(*router).Get("/", controller.GetProducts) //Liệt kê
	(*router).Get("/query", controller.GetProductsByPaging)
	(*router).Get("/query/v2", controller.GetProductsByQuery)
	(*router).Get("/count", controller.CountProductsByQuery)
	(*router).Get("/countAll", controller.CountAllProduct)
	(*router).Get("/getone/:id", controller.GetProductById)
}

func routeCategory(router *fiber.Router) {
	(*router).Get("/", controller.GetCategories) //Liệt kê
	(*router).Get("/:id", controller.GetChildCategoriesByParent)
}
