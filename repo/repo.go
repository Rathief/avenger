package repo

import (
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}
