package api

import (
	"ecommerce/internal/application/order"
	"ecommerce/internal/application/order/dtos"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderController interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Cancel(c *fiber.Ctx) error
	Initialize (app *fiber.App, authorizeFunc func(*fiber.Ctx) error)
}

type orderController struct {
	appService order.AppService
}

func (controller orderController) Initialize(app *fiber.App,  authorizeFunc func(*fiber.Ctx) error){
	group := app.Group("/api/order")
	group.Get("/", authorizeFunc, controller.GetAll)
	group.Get("/:id", authorizeFunc, controller.GetById)
	group.Post("/",  authorizeFunc, controller.Create)
	group.Put("/:id/cancel", controller.Cancel)
}

func NewOrderController (appService order.AppService) OrderController{
	return orderController{appService: appService}
}

func (controller orderController) GetAll (c *fiber.Ctx) error {
	orders := controller.appService.GetAll()
	data,_ := json.Marshal(orders)
	c.Write(data)

	return nil
}

func (controller orderController) GetById (c *fiber.Ctx) error {
	id := c.Params("id")
	order := controller.appService.GetById(uuid.MustParse(id))
	data,_ := json.Marshal(order)
	c.Write(data)

	return nil
}

func (controller orderController) Create (c *fiber.Ctx) error {
	data := []byte(c.Body())
	var createOrderDto dtos.CreateOrderDto
	json.Unmarshal(data, &createOrderDto)
	controller.appService.Create(createOrderDto)

	return nil
}

func (controller orderController) Cancel (c *fiber.Ctx) error {
	id := c.Params("id")
	controller.appService.Cancel(uuid.MustParse(id))
	c.SendStatus(fiber.StatusOK)

	return nil
}