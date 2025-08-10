package controllers

import (
	"AmethToledo/src/donaciones/application"
	"AmethToledo/src/donaciones/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateDonacionController struct {
	updateDonacion *application.UpdateDonacion
}

func NewUpdateDonacionController(updateDonacion *application.UpdateDonacion) *UpdateDonacionController {
	return &UpdateDonacionController{
		updateDonacion: updateDonacion,
	}
}

type UpdateDonacionRequest struct {
	UsuarioID     int     `json:"usuario_id" binding:"required"`
	ModuloID      int     `json:"modulo_id" binding:"required"`
	Monto         float64 `json:"monto" binding:"required,gt=0"`
	Moneda        string  `json:"moneda"`
	Estado        string  `json:"estado" binding:"required"`
	MetodoPago    string  `json:"metodo_pago"`
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
}

func (udc *UpdateDonacionController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateDonacionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer moneda por defecto si no se especifica
	if req.Moneda == "" {
		req.Moneda = "MXN"
	}

	donacion := entities.Donacion{
		ID:            id,
		UsuarioID:     req.UsuarioID,
		ModuloID:      req.ModuloID,
		Monto:         req.Monto,
		Moneda:        req.Moneda,
		Estado:        req.Estado,
		MetodoPago:    req.MetodoPago,
		TransactionID: req.TransactionID,
		PaymentID:     req.PaymentID,
	}

	err = udc.updateDonacion.Execute(donacion)
	if err != nil {
		if err.Error() == "donación no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donación actualizada exitosamente",
	})
}
