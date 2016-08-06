package DTO

import (
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
)

type Event struct {
	Id              uint                            `json:"id"`
	Name            string                          `json:"name"`
	StartDate       string                          `json:"startDate"`
	EndDate         string                          `json:"endDate"`
	Address         string                          `json:"address"`
	Credit          float64                         `json:"credit"`
	Constraints     []string                        `json:"constraints"`
	Details         string                          `json:"details"`
	MaxParticipants uint                            `json:"maxParticipants"`
	Status          string                          `json:"status"`
	Type            string                          `json:"type"`
	Group           string                          `json:"group"`
	AttachedProgram uint                            `json:"attachedProgram"`
}

// TODO: Need to get the status from the DB
var eventStatus = []string{"DELETED", "ONLINE", "OFFLINE", "REFUSED", "WAITING_FOR_APPROVAL"}

func (e *Event) CheckFields() (err string) {
	if (e.Name == "" || e.StartDate == "" || e.EndDate == "" || e.Address == "" || e.AttachedProgram == 0) {
		if e.Name == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventFieldNameMissing"))
		}
		if e.StartDate == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventFieldStartDateMissing"))
		}
		if e.EndDate == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventFieldEndDateMissing"))
		}
		if e.Address == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventFieldAddressMissing"))
		}
		if e.AttachedProgram == 0 {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventFieldAttachedProgramMissing"))
		}
		if e.checkStatus() == false {
			err = utils.AppendStringHTMLNewLine(err, translate.T("eventStatusUnknwon"))
		}
	}
	return
}

func (e *Event) checkStatus() bool {
	for i := 0; i < len(eventStatus); i++ {
		if eventStatus[i] == e.Status {
			return true;
		}
	}
	return false
}

//// TODO implement enum constraint
//type EventRegistrationConstraints int
//
//const (
//	CANT_BE_DROPPED EventRegistrationConstraints = iota
//	ONE_BY_GROUPE
//	ONE_BY_TYPE
//)