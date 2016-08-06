package service

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"time"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
)

func CreateEvent(e *DTO.Event) (utils.ReturnMsg) {
	return repository.CreateEvent(DAO.Event{
		Address:e.Address,
		//Constraints:e.Constraints,
		CreatedAt:time.Now(),
		Credit:e.Credit,
		Details:e.Details,
		EndDate:utils.ChangeStringToDateWithHours(e.EndDate),
		Name:e.Name,
		//Group:e.Group,
		MaxParticipants:e.MaxParticipants,
		AttachedProgram:e.AttachedProgram,
		StartDate:utils.ChangeStringToDateWithHours(e.StartDate),
		Status:e.Status,
		//Type:e.Type,
	})
}

func RetrieveEvent(id int) (DTO.Event, utils.ReturnMsg) {
	e, rMsg := repository.RetrieveEvent(id)
	return DTO.Event{
		Id:e.ID,
		Status:e.Status,
		//Type:e.Type,
		Address:e.Address,
		//Constraints:e.,
		Credit:e.Credit,
		Details:e.Details,
		EndDate:e.EndDate.Format("02-01-2006 15:04:05"),
		StartDate:e.StartDate.Format("02-01-2006 15:04:05"),
		//Group:e.
		MaxParticipants:e.MaxParticipants,
		Name:e.Name,
	}, rMsg
}

func RetrieveAllEvent(id int) (events []DTO.Event, rMsg utils.ReturnMsg) {
	es, rMsg := repository.RetrieveAllEvent(id)
	for i := 0; i < len(es); i++ {
		events = append(events, DTO.Event{
			Status:es[i].Status,
			//Type:es[i].Type,
			Address:es[i].Address,
			//Constraints:es[i].,
			Credit:es[i].Credit,
			Details:es[i].Details,
			EndDate:es[i].EndDate.Format("02-01-2006 15:04:05"),
			StartDate:es[i].StartDate.Format("02-01-2006 15:04:05"),
			//Group:es[i].
			MaxParticipants:es[i].MaxParticipants,
			Name:es[i].Name,
			Id:es[i].ID,
		})
	}
	return
}