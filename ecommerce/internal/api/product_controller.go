package api

import (
	"ecommerce/internal/application/product"
	"ecommerce/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController interface {
	GetAll(c *fiber.Ctx) error
	GetAllByMerchantId(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Initialize (app *fiber.App, jwtAuth func(*fiber.Ctx) error)
}

type productController struct {
	appService product.AppService
}

func NewProductController (appService product.AppService) ProductController{
	return productController{appService: appService}
}

func (controller productController) Initialize (app *fiber.App, jwtAuth func(*fiber.Ctx) error){
	group := app.Group("/api/product")
	group.Get("/", jwtAuth, controller.GetAll)
	group.Get("/:merchantId/by-merchant-id", controller.GetAllByMerchantId)
	group.Get("/:id", controller.GetById)
}

func (controller productController) GetAll(c *fiber.Ctx) error {

	fmt.Println("here!")
	claims, err := util.ExtractTokenMetadata(c)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(claims)
	product := controller.appService.GetAll()
	data, _ := json.Marshal(product)
	c.Write(data)
	return nil
}

func (controller productController) GetAllByMerchantId(c *fiber.Ctx) error{
	id := c.Params("merchantId")
	products := controller.appService.GetAllByMerchantId(uuid.MustParse(id))
	data,_ := json.Marshal(products)
	c.Write(data)
	return nil
}

func (controller productController) GetById(c *fiber.Ctx) error{
	id := c.Params("id")
	product := controller.appService.GetById(uuid.MustParse(id))
	data,_ := json.Marshal(product)
	c.Write(data)
	return nil
}