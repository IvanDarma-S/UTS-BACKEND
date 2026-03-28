package repositories

import (
	"github.com/agrahafiz13/gin-firebase-backend/config"
	"github.com/agrahafiz13/gin-firebase-backend/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}