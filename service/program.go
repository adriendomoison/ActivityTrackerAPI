package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
	"sort"
)

func CreateProgram(program *DTO.ProgramCreation) DTO.ReturnMsg {
	p := DAO.ProgramCreation{}
	p.Name = program.Name
	p.EnrolledStudents = program.EnrolledStudents
	p.NbrOfHoursToComplete = program.NbrOfHoursToComplete
	return repository.CreateProgram(p)
}

func RetrieveAllPrograms() (programs []DTO.Program, rMsg DTO.ReturnMsg) {
	programsData, rMsg := repository.RetrieveProgram()
	var keys []string
	for k := range programsData {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		np := DTO.Program{
			programsData[k].Id,
			programsData[k].Name,
			programsData[k].Nbr_event,
			programsData[k].Nbr_hours_available,
			programsData[k].Nbr_student,
			programsData[k].Nbr_student_danger,
			programsData[k].Enrolled_student,
			programsData[k].Nbr_hours_to_complete,
		}
		programs = append(programs, np)
	}
	return programs, rMsg
}