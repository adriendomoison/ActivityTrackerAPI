package repository

import (
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
)

func CreateEvent(e DAO.Event) (utils.ReturnMsg) {
	if err := db.Create(&e).Error; err != nil {
		log.Println(err)
		/*
		 * TODO find a way to say that the event can possibly be a duplication for a certain program,
		 * TODO this doesn't work and should be replace by "error unknown" or the name of the error thrown by mysql
		 */
		return utils.ReturnMsg{1, translate.T("eventAlreadyExist")}
	}
	return utils.ReturnMsg{0, translate.T("eventCreated")}
}

func RetrieveEvent(id int) (event DAO.Event, rMsg utils.ReturnMsg) {
	db.First(&event, id)
	if event.ID != 0 {
		return event, utils.ReturnMsg{0, ""}
	}
	return event, utils.ReturnMsg{1, translate.T("eventNotFound")}
}

func RetrieveAllEvent(id int) (events []DAO.Event, rMsg utils.ReturnMsg) {
	db.Where("fk_program = ?", id).Find(&events)
	if 0 < len(events) {
		return events, utils.ReturnMsg{0, ""}
	}
	return events, utils.ReturnMsg{1, translate.T("eventNoResult")}
}
// TODO implement types and groups
//func RetrieveEventTypes() (types []DAO.Type) {
//	db.Find(&types)
//	return
//}