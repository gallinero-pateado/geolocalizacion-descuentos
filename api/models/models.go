package models

// Usuario representa la ubicación del usuario
type Usuario struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Descuento representa cada producto
type Descuento struct {
	ID              int     `json:"id"`
	Titulo          string  `json:"titulo"`
	Descripcion     string  `json:"descripcion"`
	Precio          int     `json:"precio"`
	Precio_Anterior int     `json:"precio_anterior"`
	Descuento       int     `json:"descuento"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
}
