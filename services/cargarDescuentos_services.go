package services

import (
	"fmt"
	"geolocalizacion/api/models"
)

// Variable global que contiene los descuentos cargados
var discounts []models.Descuento

// Inicializar los descuentos al arrancar el servicio
func InitDiscounts() error {
	var err error
	discounts, err = LoadDiscountsFromJSON("./descuentos.json")
	if err != nil {
		return err
	}
	return nil
}

// GetDiscountsNearby devuelve los descuentos cercanos a la ubicación proporcionada
func GetDiscountsNearby(latUser, lonUser float64) []models.Descuento {
	var nearbyDiscounts []models.Descuento

	// Usar la función CalculateDistance que está en geolocalizacion_services.go
	for _, discount := range discounts {
		fmt.Printf("Discount ID: %d, Latitude: %f, Longitude: %f\n", discount.ID, discount.Latitude, discount.Longitude)
		distance := CalculateDistance(latUser, lonUser, discount.Latitude, discount.Longitude) // Llama directamente a CalculateDistance
		if distance <= 5 {                                                                     // Filtra los descuentos a menos de 5 km
			nearbyDiscounts = append(nearbyDiscounts, discount)
		}
	}

	return nearbyDiscounts
}
