package repositories

import (
	"github.com/agrahafiz13/gin-firebase-backend/config"
	"github.com/agrahafiz13/gin-firebase-backend/models"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
