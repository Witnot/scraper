package api

import (
	"net/http"

	"github.com/Witnot/scraper/internal/db"
	"github.com/Witnot/scraper/internal/models"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		var ps []models.Product
		db.DB.Find(&ps)
		c.JSON(http.StatusOK, ps)
	})

	r.GET("/products/:id/prices", func(c *gin.Context) {
		id := c.Param("id")
		var prs []models.PriceRecord
		db.DB.Where("product_id = ?", id).Order("recorded_at desc").Find(&prs)
		c.JSON(http.StatusOK, prs)
	})

	r.GET("/reports/price-trend", func(c *gin.Context) {
		// Example SQL query for Grafana / API aggregation
		productID := c.Query("product_id")
		rows, _ := db.DB.Raw(`SELECT date_trunc('day', recorded_at) as day, AVG(price) as avg_price FROM price_records WHERE product_id = ? GROUP BY day ORDER BY day`, productID).Rows()
		defer rows.Close()
		// map rows to JSON response (left as exercise)
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.Run(":8080")
}
