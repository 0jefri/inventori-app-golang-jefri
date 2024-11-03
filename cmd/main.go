package main

import (
	"github.com/inventori-app-jeff/config"
	"github.com/inventori-app-jeff/internal/app/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {
	delivery.Server().Run()
}
