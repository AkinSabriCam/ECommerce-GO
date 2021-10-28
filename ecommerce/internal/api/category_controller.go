package api

import (
	"ecommerce/internal/application/category"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAll(c *fiber.Ctx) error
	Initialize (app *fiber.App, authorizeFunc func(*fiber.Ctx) error)
}

type categoryController struct {
	appService category.AppService
}

func NewCategoryController(appService category.AppService) CategoryController {
	return categoryController{appService: appService}
}

func (controller categoryController) Initialize(app *fiber.App, authorizeFunc func(*fiber.Ctx) error) {
	group := app.Group("/api/category")
	group.Get("/", authorizeFunc, controller.GetAll)
}

func (controller categoryController) GetAll(c *fiber.Ctx) error {
	categories := controller.appService.GetAll()
	data,_ := json.Marshal(categories)
	c.Write(data)

	return nil
}