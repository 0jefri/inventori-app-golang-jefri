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
	// transactionController := controller.NewTransactionController(repoManager.TransactionService())
	// Bill Controller
	// billController := controller.NewBillController(repoManager.BillService())
	// Contact Controller
	// contactController := controller.NewContactController(repoManager.ContactService())
	// Card Controller
	// cardController := controller.NewCardController(repoManager.CardService())

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
				// users.POST("/:id/upload", userController.UploadPicture)
				// users.GET("/:id/download", userController.DownloadPicture)
				users.DELETE("/:id", userController.DeleteUser)

				// Product
				users.POST("/product", productController.AddProduct)
				users.GET("/product/:id", productController.FindProduct)
				users.GET("/products", productController.FindAllProducts)
				users.DELETE("/product/:id", productController.DeleteProduct)
				users.PUT("/product/:id", productController.UpdateProduct)
				users.GET("/product/name", productController.FindProductByName)

				// Contact
				// users.POST("/:id/contacts", contactController.AddContact)
				// users.GET("/:id/contacts", contactController.FindAllContacts)
				// users.GET("/:id/contacts/:contactID", contactController.FindContact)
				// users.DELETE("/:id/contacts/:contactID", contactController.DeleteContact)
				// // Transaction
				// users.POST("/:id/transactions/deposit/:cardID", transactionController.Deposit)
				// users.POST("/:id/transactions/send/:userID", transactionController.SendMoney)
				// users.POST("/:id/transactions/withdraw/:cardID", transactionController.WithdrawMoney)
				// users.GET("/:id/transactions", transactionController.FindAllTransactions)
				// users.GET("/:id/transactions/:transactionID", transactionController.FindTransaction)
				// Bill
				// users.POST("/:id/bills", billController.CreateBill)
				// users.GET("/:id/bills", billController.FindAllBills)
				// // Card
				// users.GET("/:id/cards", cardController.FindAllCards)
				// users.POST("/:id/cards", cardController.AddCard)
				// users.DELETE("/:id/cards/:cardID", cardController.DeleteCard)
			}
		}
	}

	return router.Run()

}
