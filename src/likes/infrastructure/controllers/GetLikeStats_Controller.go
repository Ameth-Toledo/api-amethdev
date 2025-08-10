package controllers

import (
	"AmethToledo/src/likes/application"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GetLikeStatsController struct {
	getLikeStats *application.GetLikeStats
}

func NewGetLikeStatsController(getLikeStats *application.GetLikeStats) *GetLikeStatsController {
	return &GetLikeStatsController{
		getLikeStats: getLikeStats,
	}
}

type GetLikeStatsRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}

func (glsc *GetLikeStatsController) Execute(c *gin.Context) {
	moduloIDStr := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	var req GetLikeStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros de fecha requeridos (start_date, end_date)"})
		return
	}

	startDate, err := parseFlexibleDate(req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             fmt.Sprintf("failed to parse start_date: %v", err),
			"received":          req.StartDate,
			"supported_formats": []string{"2006-01-02", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05-07:00"},
		})
		return
	}

	endDate, err := parseFlexibleDate(req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             fmt.Sprintf("failed to parse end_date: %v", err),
			"received":          req.EndDate,
			"supported_formats": []string{"2006-01-02", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05-07:00"},
		})
		return
	}

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, time.UTC)

	response, err := glsc.getLikeStats.Execute(moduloID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Estadísticas de likes obtenidas exitosamente",
		"data":    response,
	})
}

func parseFlexibleDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	formats := []string{
		"2006-01-02",                // YYYY-MM-DD (most common for your use case)
		time.RFC3339,                // Standard RFC3339: "2006-01-02T15:04:05Z07:00"
		"2006-01-02T15:04:05Z",      // ISO format with Z (UTC)
		"2006-01-02T15:04:05-07:00", // ISO format with timezone
		"2006-01-02T15:04:05",       // ISO format without timezone
		"2006-01-02 15:04:05",       // YYYY-MM-DD HH:MM:SS
		"01/02/2006",                // MM/DD/YYYY
		"02/01/2006",                // DD/MM/YYYY
	}

	var lastErr error
	for _, format := range formats {
		parsed, err := time.Parse(format, dateStr)
		if err == nil {
			if format == "2006-01-02" || format == "01/02/2006" || format == "02/01/2006" {
				result := time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, time.UTC)
				return result, nil
			}
			return parsed, nil
		} else {
			lastErr = err
		}
	}

	return time.Time{}, fmt.Errorf("unsupported date format '%s': %v", dateStr, lastErr)
}
