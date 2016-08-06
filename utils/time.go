package utils

import (
	"time"
	"log"
)

func ChangeStringToDate(s string) (date time.Time) {
	date, err := time.Parse("02-01-2006", s)
	if err != nil {
		log.Println(err)
	}
	return
}

func ChangeStringToDateWithHours(s string) (date time.Time) {
	date, err := time.Parse("02-01-2006 15:04:05", s)
	if err != nil {
		log.Println(err)
	}
	return
}