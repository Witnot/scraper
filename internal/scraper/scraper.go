package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Witnot/scraper/internal/db"
	"github.com/Witnot/scraper/internal/models"
)

type FakeStoreProduct struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func ScrapeFakeStoreProduct(ctx context.Context, url string) error {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var product FakeStoreProduct
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return err
	}

	// upsert product
	var p models.Product
	db.DB.Where(models.Product{Source: "fakestoreapi", ExternalID: fmt.Sprintf("%d", product.ID)}).
		FirstOrCreate(&p, models.Product{
			Source:     "fakestoreapi",
			ExternalID: fmt.Sprintf("%d", product.ID),
			Name:       product.Title,
			URL:        url,
		})

	// create price record
	pr := models.PriceRecord{
		ProductID:  p.ID,
		Price:      product.Price,
		Currency:   "USD",
		RecordedAt: time.Now(),
	}
	db.DB.Create(&pr)
	return nil
}
