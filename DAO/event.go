package DAO

import "time"

type Event struct {
	ID              uint      `gorm:"column:idevent;primary_key"`
	Name            string      `gorm:"column:name"`
	StartDate       time.Time   `gorm:"column:start_date"`
	EndDate         time.Time   `gorm:"column:end_date"`
	Address         string      `gorm:"column:address"`
	Credit          float64     `gorm:"column:credit"`
	// TODO: Constraints     []string    `gorm:"column:constraints"`
	Details         string      `gorm:"column:details"`
	MaxParticipants uint        `gorm:"column:max_participants"`
	Status          string      `gorm:"column:fk_status"`
	// TODO Type            string      `gorm:"column:type"`
	// TODO Group           string      `gorm:"column:group"`
	AttachedProgram uint        `gorm:"column:fk_program"`
	CreatedAt       time.Time   `gorm:"column:created_date"`
}

type Type struct {
	Name string
}

//func (e Event) CheckType(types []string) bool {
//	for i := 0; i < len(types); i++ {
//		if e.Type == types[i] {
//			return true
//		}
//	}
//	return false
//}