package main

import (
	"management/api"
	"management/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})
}
