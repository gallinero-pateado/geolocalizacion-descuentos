// services.go
package services

import (
	"math"
)

// CalculateDistance formula para calcular la distancia entre dos puntos
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const EarthRadius = 6371 // Radio de la Tierra en kilómetros

	latDiff := toRadians(lat2 - lat1)
	lonDiff := toRadians(lon2 - lon1)

	a := math.Sin(latDiff/2)*math.Sin(latDiff/2) +
		math.Cos(toRadians(lat1))*math.Cos(toRadians(lat2))*
			math.Sin(lonDiff/2)*math.Sin(lonDiff/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c
}

// Convertir grados a radianes
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
