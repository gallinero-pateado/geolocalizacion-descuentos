package main

import (
	"geolocalizacion/api/handlers"
	"geolocalizacion/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar los descuentos desde el archivo JSON al iniciar la aplicación
	err := services.InitDiscounts()
	if err != nil {
		log.Fatal("Error al cargar los descuentos:", err)
	}

	router := gin.Default()
	// Rutas para los descuentos
	router.GET("/descuentos/cercanos", handlers.GetDiscounts)

	router.Run(":8080")
}
