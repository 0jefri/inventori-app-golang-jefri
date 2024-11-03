package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/inventori-app-jeff/internal/app/delivery/routes"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := routes.SetupRouter(app.engine); err != nil {
		panic("Application error")
	}
}

func Server() *application {
	router := gin.Default()

	return &application{
		engine: router,
	}

}
