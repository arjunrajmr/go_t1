package main

import (
	"log"

	routes "github.com/arjunrajmr/go_t1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
