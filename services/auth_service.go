package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/agrahafiz13/gin-firebase-backend/config"
	"github.com/agrahafiz13/gin-firebase-backend/models"
	"github.com/agrahafiz13/gin-firebase-backend/repositories"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{userRepo: repositories.NewUserRepository()}
}