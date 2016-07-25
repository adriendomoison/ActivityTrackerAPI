package DTO

import "github.com/adriendomoison/ActivityTrackerAPI/DAO"

type ProgramCreation struct {
	Name                 string `json:"name"`
	EnrolledStudents     []int  `json:"entrolledStudents"`
	NbrOfHoursToComplete int    `json:"nbrOfHoursToComplete"`
}

type Program struct {
	Id                   int
	Name                 string
	NbrEvent             int
	NbrHoursAvailable    int
	NbrStudent           int
	NbrStudentDanger     int
	EnrolledStudents     []DAO.StudentBasic
	NbrOfHoursToComplete int
}