package api

import (
	"ecommerce/internal/application/basket"
	"ecommerce/internal/application/basket/dtos"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type BasketController interface {
	addNewProductToBasket(c *fiber.Ctx) error
	getMyBasket(c *fiber.Ctx) error
	Initialize (app *fiber.App,  authorizeFunc func(*fiber.Ctx) error)
}

type basketController struct {
	appService basket.AppService
}

func NewBasketController(appService basket.AppService) BasketController {
	return basketController{appService: appService}
}

func (controller basketController) Initialize(app *fiber.App, authorizeFunc func(*fiber.Ctx) error) {
	v1 := app.Group("api/basket")
	v1.Post("/new-product", authorizeFunc, controller.addNewProductToBasket)
	v1.Get("/my-basket", authorizeFunc, controller.getMyBasket)
}

func (controller basketController) addNewProductToBasket(c *fiber.Ctx) error {
	var dto dtos.AddProductToBasketDto
	request := []byte(c.Body()) // todo: do not convert to  byte slice,use as string and than serialize to dto structure
	json.Unmarshal(request, &dto)
	controller.appService.AddNewProductToBasket(dto)

	return nil
}

func (controller basketController) getMyBasket(c *fiber.Ctx) error {
	basket := controller.appService.GetMyBasket()
	data,_ := json.Marshal(basket)
	c.Write(data)

	return nil
}