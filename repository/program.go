package repository

import (
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
)

func CreateProgram(p DAO.Program) utils.ReturnMsg {
	if err := db.Create(&p).Error; err != nil {
		log.Println(err)
		return utils.ReturnMsg{1, translate.T("programAlreadyExist")}
	}
	return utils.ReturnMsg{0, translate.T("programCreated")}
}

func RetrieveAllPrograms() (programs []DAO.Program, rMsg utils.ReturnMsg) {
	db.Find(&programs)
	for i := 0; i < len(programs); i++ {
		db.Where(DAO.ProgramsSubscription{FkProgram: programs[i].ID}).Find(&programs[i].Users)
	}
	return programs, utils.ReturnMsg{0, ""}
}

func RetrieveProgram(id int) (program DAO.Program, rMsg utils.ReturnMsg) {
	db.Find(&program, id)
	db.Where(DAO.ProgramsSubscription{FkProgram: program.ID}).Find(&program.Users)
	return program, utils.ReturnMsg{0, ""}
}

func DeleteProgram(id int) (utils.ReturnMsg) {
	db.Delete(DAO.Program{}, id)
	return utils.ReturnMsg{0, translate.T("programDeleted")}
}
