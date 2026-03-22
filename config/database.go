package config

import (
	"fmt"
	"log"
	"os"

	"github.com/agrahafiz13/gin-firebase-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB adalah instance GORM global yang dipakai di seluruh aplikasi
var DB *gorm.DB

func InitDatabase() {
	// Ambil konfigurasi dari environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	}