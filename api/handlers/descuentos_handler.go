package handlers

import (
	"geolocalizacion/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDiscounts(c *gin.Context) {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	// Verificar si los parámetros se proporcionaron
	if latitudeStr == "" || longitudeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ubicación no proporcionada"})
		return
	}

	// Convertir latitude y longitude de string a float64
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitud inválida"})
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Longitud inválida"})
		return
	}

	// Llamar al servicio que busca descuentos cercanos con los valores convertidos
	discounts := services.GetDiscountsNearby(latitude, longitude)

	// Devolver la respuesta con los descuentos cercanos
	c.JSON(http.StatusOK, gin.H{
		"discounts": discounts,
	})
}
