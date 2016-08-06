package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
	"time"
)

func CreateProgram(p *DTO.Program) utils.ReturnMsg {
	return repository.CreateProgram(
		DAO.Program{
			Name:p.Name,
			Description:p.Description,
			StartDate:utils.ChangeStringToDate(p.StartDate),
			EndDate:utils.ChangeStringToDate(p.EndDate),
			NbrHoursToComplete:p.NbrOfHoursToComplete,
			CreatedAt: time.Now(),
		})
}

func RetrieveAllPrograms() (programs []DTO.Program, rMsg utils.ReturnMsg) {
	p, rMsg := repository.RetrieveAllPrograms()
	
	for i := 0; i < len(p); i++ {
		programs = append(programs, DTO.Program{
			p[i].ID,
			p[i].Name,
			p[i].Description,
			p[i].StartDate.Format("02-01-2006"),
			p[i].EndDate.Format("02-01-2006"),
			p[i].NbrHoursToComplete,
			p[i].GetUsersIdOnly(),
		})
	}
	return
}

func RetrieveProgram(id int) (DTO.Program, utils.ReturnMsg) {
	p, rMsg := repository.RetrieveProgram(id)
	return DTO.Program{
		p.ID,
		p.Name,
		p.Description,
		p.StartDate.Format("2006-01-02"),
		p.EndDate.Format("2006-01-02"),
		p.NbrHoursToComplete,
		p.GetUsersIdOnly(),
	}, rMsg
}

func DeleteProgram(id int) utils.ReturnMsg {
	return repository.DeleteProgram(id)
}