package DTO

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
)

type ProgramCreation struct {
	Name                 string `json:"name"`
	EnrolledStudents     []int  `json:"entrolledStudents"`
	NbrOfHoursToComplete int    `json:"nbrOfHoursToComplete"`
}

type Program struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	NbrOfHoursToComplete int    `json:"nbrOfHoursToComplete"`
	EnrolledStudents     []DAO.StudentBasic `json:"enrolledStudent"`
}