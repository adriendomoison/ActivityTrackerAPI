package DAO

import (
	"time"
)

type Program struct {
	Name               string `gorm:"column:name"`
	Description        string `gorm:"column:description"`
	StartDate          string `gorm:"column:start_date"`
	EndDate            string `gorm:"column:end_date"`
	NbrHoursToComplete int    `gorm:"column:nbr_hours_to_complete"`
	CreatedAt          time.Time `gorm:"column:created_date"`
}