package services

import (
	"encoding/json"
	"geolocalizacion/api/models"
	"io/ioutil"
	"log"
	"os"
)

// Estructura para cargar el JSON
type DiscountData struct {
	Usuario    models.Usuario     `json:"usuario"`
	Descuentos []models.Descuento `json:"productos"`
}

func LoadDiscountsFromJSON(filepath string) ([]models.Descuento, error) {
	// Abre el archivo JSON
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Println("Error abriendo el archivo JSON:", err)
		return nil, err
	}
	defer jsonFile.Close()

	// Lee el contenido del archivo
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Define una variable para almacenar los descuentos y el usuario
	var discountData DiscountData

	// Convierte el JSON a la estructura de Go
	err = json.Unmarshal(byteValue, &discountData)
	if err != nil {
		log.Println("Error al convertir el archivo JSON:", err)
		return nil, err
	}

	// Retorna solo la lista de descuentos
	return discountData.Descuentos, nil
}
