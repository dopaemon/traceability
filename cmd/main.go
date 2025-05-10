package main

import (
    "github.com/gin-gonic/gin"
    "traceability/routes"
    "traceability/services"
)

func main() {
    services.InitMongo()

    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8081")
}
