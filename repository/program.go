package repository

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
	"github.com/go-sql-driver/mysql"
	"github.com/kisielk/sqlstruct"
	"fmt"
)

func CreateProgram(p DAO.ProgramCreation) DTO.ReturnMsg {
	res, err := db.Exec("INSERT INTO program (name, nbr_event, nbr_student, nbr_hours_available, " +
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

func RetrieveProgram() (map[string]DAO.Program, DTO.ReturnMsg) {
	
	var programs = make(map[string]DAO.Program)
	
	sql := "SELECT %s, %s FROM program AS p " +
		"INNER JOIN program_subscription AS ps ON p.idprogram = ps.fk_program " +
		"INNER JOIN user AS s ON s.iduser = ps.fk_user"
	sql = fmt.Sprintf(sql, sqlstruct.ColumnsAliased(DAO.Program{}, "p"), sqlstruct.ColumnsAliased(DAO.StudentBasic{}, "s"))
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var program DAO.Program
		var student DAO.StudentBasic
		
		err = sqlstruct.ScanAliased(&program, rows, "p")
		if err != nil {
			log.Fatal(err)
		}
		err = sqlstruct.ScanAliased(&student, rows, "s")
		// TODO; fix error with Timestamp
		if err != nil {
			log.Println(err)
		}
		if _, ok := programs[program.Name]; !ok {
			programs[program.Name] = program
		}
		appendStudentToProgram(programs, program.Name, student)
	}
	return programs, DTO.ReturnMsg{}
}

func appendStudentToProgram(p map[string]DAO.Program, n string, s DAO.StudentBasic) {
	tmp := p[n]
	tmp.Enrolled_student = append(p[n].Enrolled_student, s)
	p[n] = tmp
}