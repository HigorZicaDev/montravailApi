package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uint32
	Student     string
	Teacher     string
	Description string
	DateEvent   time.Time
	Color       string
	Status      bool
}
