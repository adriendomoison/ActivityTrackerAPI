package DTO

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
)

type Program struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	StartDate            string `json:"startDate"`
	EndDate              string `json:"endDate"`
	NbrOfHoursToComplete int    `json:"nbrOfHoursToComplete"`
	EnrolledStudents     []DAO.StudentBasic `json:"enrolledStudent"`
}

func (p Program) CheckField() (err string) {
	if p.Name == "" || p.Description == "" || p.StartDate == "" || p.EndDate == "" || p.NbrOfHoursToComplete == 0 {
		if p.Name == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("fieldNameMissing"))
		}
		if p.Description == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("fieldDescriptionMissing"))
		}
		if p.StartDate == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("fieldStartDateMissing"))
		}
		if p.EndDate == "" {
			err = utils.AppendStringHTMLNewLine(err, translate.T("fieldEndDateMissing"))
		}
		if p.NbrOfHoursToComplete == 0 {
			err = utils.AppendStringHTMLNewLine(err, translate.T("fieldNbrOfHoursToCompleteMissing"))
		}
	}
	return
}