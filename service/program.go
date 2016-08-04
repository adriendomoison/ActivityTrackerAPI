package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
)

func CreateProgram(p *DTO.Program) utils.ReturnMsg {
	return repository.CreateProgram(
		DAO.Program{
			Name:p.Name,
			Description:p.Description,
			StartDate:p.StartDate,
			EndDate:p.EndDate,
			NbrHoursToComplete:p.NbrOfHoursToComplete,
		})
}

func RetrieveAllPrograms() (programs []DTO.Program, rMsg utils.ReturnMsg) {
	//var programsData map[string]DAO.Program
	//var keys []string
	//
	//programsData, rMsg = repository.RetrieveAllPrograms()
	//for k := range programsData {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//for _, k := range keys {
	//	startDate, endDate := findDate(programsData[k])
	//	programs = append(programs, DTO.Program{programsData[k].Id, programsData[k].Name, programsData[k].Description,
	//		time.Time(startDate).Format("2006-01-02"), time.Time(endDate).Format("2006-01-02"),
	//		programsData[k].Nbr_hours_to_complete})
	//}
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