package repository

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
	"github.com/go-sql-driver/mysql"
)

func CreateProgram(p DAO.Program) (DTO.ReturnMsg) {
	var err error
	res, err := db.Exec("INSERT INTO program (name, nbr_event, nbr_student, nbr_hours_available_per_student, " +
		"nbr_student_danger, fk_creator, nbr_hours_to_complete) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?)", p.Name, 0, 0, 0, 0, 0, p.NbrOfHoursToComplete)
	if err != nil {
		log.Println(err);
		if pErr, ok := err.(*mysql.MySQLError); ok && pErr.Number == ER_DUP_ENTRY {
			return DTO.ReturnMsg{-1, ErDupEntry(pErr)}
		} else {
			return DTO.ReturnMsg{-1, translate.T("errorDatabase")}
		}
	}
	n, err := res.RowsAffected()
	if err != nil {
		log.Println(err);
		return DTO.ReturnMsg{-1, translate.T("errorDatabase")}
	}
	if n > 0 {
		return DTO.ReturnMsg{0, translate.T("programCreated")}
	}
	return DTO.ReturnMsg{1, translate.T("noEntryCreated")}
}
