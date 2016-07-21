package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
)

func CreateProgram(program *DTO.Program) (DTO.ReturnMsg) {
	p := DAO.Program{}
	p.Name = program.Name
	p.EnrolledStudents = program.EnrolledStudents
	p.NbrOfHoursToComplete = program.NbrOfHoursToComplete
	return repository.CreateProgram(p)
}
