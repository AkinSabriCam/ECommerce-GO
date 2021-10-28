package main

import (
	"ecommerce/internal/api"
	basket2 "ecommerce/internal/application/basket"
	category2 "ecommerce/internal/application/category"
	merchant2 "ecommerce/internal/application/merchant"
	order2 "ecommerce/internal/application/order"
	product2 "ecommerce/internal/application/product"
	"ecommerce/internal/domain/basket"
	"ecommerce/internal/domain/category"
	"ecommerce/internal/domain/merchant"
	"ecommerce/internal/domain/order"
	"ecommerce/internal/domain/product"
	"ecommerce/internal/infrastructure/repositories"
	"ecommerce/pkg/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func JWTAuthentication() func(*fiber.Ctx) error {

	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey: []byte(util.JwtSecret),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
		SigningMethod:"RS256",
		AuthScheme: "Bearer",
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {

	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

func main(){
	dsn := "host=localhost user=dbadmin password=dbadmin dbname=StockDb port=5433"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.OrderDetail{})
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&basket.Basket{})
	db.AutoMigrate(&basket.BasketDetail{})
	db.AutoMigrate(&merchant.Merchant{})
	db.AutoMigrate(&merchant.MerchantContact{})

	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	basketRepository := repositories.NewBasketRepository(db)
	merchantRepository := repositories.NewMerchantRepository(db)

	orderDomainService := order.NewDomainService(orderRepository)
	basketDomainService := basket.NewBasketDomainService(basketRepository, productRepository)

	productAppService := product2.NewProductAppService(productRepository)
	orderAppService := order2.NewOrderAppService(orderRepository, orderDomainService)
	categoryAppService := category2.NewAppService(categoryRepository)
	basketAppService := basket2.NewBasketAppService(basketRepository, basketDomainService)
	merchantAppService := merchant2.NewMerchantAppService(merchantRepository)

	productController := api.NewProductController(productAppService)
	orderController := api.NewOrderController(orderAppService)
	categoryController := api.NewCategoryController(categoryAppService)
	basketController := api.NewBasketController(basketAppService)
	merchantController := api.NewMerchantController(merchantAppService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		 c.SendString("Hello World!")
		return nil
	})

	productController.Initialize(app, JWTAuthentication())
	orderController.Initialize(app, JWTAuthentication())
	basketController.Initialize(app, JWTAuthentication())
	merchantController.Initialize(app, JWTAuthentication())
	categoryController.Initialize(app, JWTAuthentication())

	app.Listen(":5316")
}
