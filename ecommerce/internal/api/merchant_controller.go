package api

import (
	"ecommerce/internal/application/merchant"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MerchantController interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Initialize (app *fiber.App, authorizeFunc func(*fiber.Ctx) error)
}

type merchantController struct {
	appService merchant.AppService
}

func NewMerchantController (appService merchant.AppService) MerchantController {
	return merchantController{appService: appService}
}

func (controller merchantController) Initialize(app *fiber.App, authorizeFunc func(*fiber.Ctx) error){
	group := app.Group("/api/merchant")
	group.Get("/", authorizeFunc, controller.GetAll)
	group.Get("/id:", authorizeFunc, controller.GetById)
}

func (controller merchantController) GetAll(c *fiber.Ctx) error {
	merchants := controller.appService.GetAll()
	data,_ := json.Marshal(merchants)
	c.Write(data)

	return nil
}

func (controller merchantController) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	merchant := controller.appService.GetById(uuid.MustParse(id))
	data,_ := json.Marshal(merchant)
	c.Write(data)

	return nil
}
