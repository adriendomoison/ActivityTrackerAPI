package DTO

type Program struct {
	Name                 string `json:"name"`
	EnrolledStudents     []int  `json:"entrolledStudents"`
	NbrOfHoursToComplete int    `json:"nbrOfHoursToComplete"`
}
