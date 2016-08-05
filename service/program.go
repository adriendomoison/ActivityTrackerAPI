package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
	"time"
	"log"
)

func changeStringToDate(s string) (date time.Time) {
	date, err := time.Parse("02-01-2006", s)
	if err != nil {
		log.Println(err)
	}
	return
}

func CreateProgram(p *DTO.Program) utils.ReturnMsg {
	return repository.CreateProgram(
		DAO.Program{
			Name:p.Name,
			Description:p.Description,
			StartDate:changeStringToDate(p.StartDate),
			EndDate:changeStringToDate(p.EndDate),
			NbrHoursToComplete:p.NbrOfHoursToComplete,
			CreatedAt: time.Now(),
		})
}

func transformToArrayOfId(users []DAO.ProgramsSubscription) (userArray []uint) {
	for i := 0; i < len(users); i++ {
		userArray = append(userArray, users[i].FkUser)
	}
	return
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
			transformToArrayOfId(p[i].Users),
		})
	}
	return
}

func RetrieveProgram(id int) (program DTO.Program, rMsg utils.ReturnMsg) {
	
	//programData, rMsg := repository.RetrieveProgram(id)
	//
	//startDate, endDate := findDate(programData)
	//program := DTO.Program{programData.Id, programData.Name, programData.Description,
	//	time.Time(startDate).Format("2006-01-02"), time.Time(endDate).Format("2006-01-02"),
	//	programData.Nbr_hours_to_complete}
	return program, rMsg
}

func DeleteProgram(id int) utils.ReturnMsg {
	return repository.DeleteProgram(id)
}

//func findDate(programData DAO.Program) (timeStart time.Time, timeEnd time.Time) {
//	var err error
//
//	if timeStart, err = time.Parse("2006-01-02 15:04:05", programData.Start_date); err != nil {
//		log.Println(err)
//	}
//	if timeEnd, err = time.Parse("2006-01-02 15:04:05", programData.End_date); err != nil {
//		log.Println(err)
//	}
//	return
//}