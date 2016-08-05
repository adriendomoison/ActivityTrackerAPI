package DAO

import (
	"time"
)

type Program struct {
	ID                 uint `gorm:"column:idprogram"`
	Name               string `gorm:"column:name"`
	Description        string `gorm:"column:description"`
	StartDate          string `gorm:"column:start_date"`
	EndDate            string `gorm:"column:end_date"`
	NbrHoursToComplete int    `gorm:"column:nbr_hours_to_complete"`
	CreatedAt          time.Time `gorm:"column:created_date"`
	Users              []ProgramsSubscription `gorm:"many2many:programs_subscriptions;"`
}

type ProgramsSubscription struct {
	FkUser    uint `gorm:"column:fk_user"`
	FkProgram uint `gorm:"column:fk_program"`
	FkStatus  string `gorm:"column:fk_status"`
}