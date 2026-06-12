package main

import (
	"log"

	"github.com/IvanDarma-S/apps_marketplace_integration_backend/config"
	"github.com/IvanDarma-S/apps_marketplace_integration_backend/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{Name: "wofl tech gundam barbatos lupus rex", Price: 1000000, Category: "third party", Stock: 50,
			Description: "gundam barbatos lupus rex dengan detail yang sangat baik skala 1/100, include standbase",
			ImageURL:    "https://fuwafuwaland.ca/cdn/shop/files/rn-image_picker_lib_temp_5afc2cdf-e3e5-4206-9bf0-2441925fa424.jpg?v=1753107807"},
		{Name: "Gundam PG NU", Price: 2500000, Category: "toys",
			Stock:       100,
			Description: "gundam PG NU dengan detail yang sangat baik, include standbase",
			ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTpniKrR7P5seYnWYsUYlNKSnVrHy-3OEW0Tw&s"},
		{Name: "SNAA Titan Greatsword Titan", Price: 295000, Category: "third party",
			Stock:       200,
			Description: "model kit SNAA Titan Greatsword Titan dengan detail yang sangat baik skala 1/144, include decal",
			ImageURL:    "https://kyoucdn.id/items/418447-model-kit-1144-snaa-sc-02-titan-greatsword-tristan-knights-of-the-round-table-series.jpg"},
		{Name: "SNAA nether emperor", Price: 425000, Category: "third party",
			Stock:       150,
			Description: "model kit SNAA nether emperor dengan detail yang sangat baik skala 1/100, include standbase",
			ImageURL:    "https://kyoucdn.id/items/491842-model-kit-1100-snaa-yr-05-nether-emperor-ap-type.jpg"},
		{Name: "gundam zero PG", Price: 30000, Category: "bandai", Stock: 30,
			Description: "gundam zero dengan detail yang sangat baik skala 1/100, include standbase",
			ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSsHUaJdAi96Hmd6v6jJamdPXqrhsGmAvaVKg&s"},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
