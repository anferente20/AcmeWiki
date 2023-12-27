package main

import (
	"github.com/anferente20/AcmeWiki/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	eng := gin.Default()

	router := routes.NewRouter(eng)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
