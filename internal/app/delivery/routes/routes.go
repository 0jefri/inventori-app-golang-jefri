package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inventori-app-jeff/config"
	"github.com/inventori-app-jeff/internal/app/delivery/controller"
	"github.com/inventori-app-jeff/internal/app/delivery/middleware"
	"github.com/inventori-app-jeff/internal/app/manager"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// User Controller
	userController := controller.NewUserController(repoManager.UserService(), repoManager.AuthService())
	//product controller
	productController := controller.NewProductController(repoManager.ProductService())
	// Transaction Controller
	transactionController := controller.NewTransactionController(repoManager.TransactionService())

	// Category Controller
	categoryController := controller.NewCategoryController(repoManager.CategoryService())

	v1 := router.Group("/api/v1")
	{
		inventori := v1.Group("/inventori")
		{
			auth := inventori.Group("/auth")
			{
				auth.POST("/register", userController.Registration)
				auth.POST("/login", userController.Login)
			}

			users := inventori.Group("/users", middleware.AuthMiddleware())
			{
				users.GET("/", userController.FindAllUsers)
				users.GET("/:id", userController.FindUser)
				users.PUT("/:id", userController.UpdateUser)
				users.DELETE("/:id", userController.DeleteUser)

				// Product
				users.POST("/product", productController.AddProduct)
				users.GET("/product/:id", productController.FindProduct)
				users.GET("/products/", productController.FindAllProducts)
				users.DELETE("/product/:id", productController.DeleteProduct)
				users.PUT("/product/:id", productController.UpdateProduct)
				users.GET("/products", productController.FindProductByName)

				// Transaction
				users.POST("/product/:id/transactions/receive", transactionController.ReceiveProduct)
				users.POST("/product/:id/transactions/send", transactionController.SendProduct)
				users.GET("/product/transactions", transactionController.ListTransactions)

				//category
				users.POST("/product/:id/category", categoryController.AddCategory)
				// users.GET("/product/:id/categorys", categoryController.FindAllCategory)
			}
		}
	}

	return router.Run()

}
