package main

import (
	"github.com/cfindlayisme/whatcanimake/requesthandlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	listenAddress := "0.0.0.0:8080"

	router.GET(requesthandlers.ApiPath+"/getRecipes", requesthandlers.GetRecipes)

	router.Run(listenAddress)
}
