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

func CreateProgram(p DAO.Program) DTO.ReturnMsg {
	res, err := db.Exec("INSERT INTO program (name, description, start_date, end_date, nbr_hours_to_complete, fk_creator) " +
		"VALUES (?, ?, ?, ?, ?, ?)", p.Name, p.Description, p.Start_date, p.End_date, p.Nbr_hours_to_complete, 0)
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

func RetrieveAllPrograms() (map[string]DAO.Program, DTO.ReturnMsg) {
	
	var programs = make(map[string]DAO.Program)
	// TODO: patch the request so it do a concatenation and get programs even when nobody subscribed to it.
	sql := "SELECT %s, %s FROM program AS p " +
		"INNER JOIN program_subscription AS ps ON p.idprogram = ps.fk_program " +
		"INNER JOIN user AS s ON s.iduser = ps.fk_user"
	sql = fmt.Sprintf(sql, sqlstruct.ColumnsAliased(DAO.Program{}, "p"), sqlstruct.ColumnsAliased(DAO.StudentBasic{}, "s"))
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return programs, DTO.ReturnMsg{-1, err.Error()}
	}
	defer rows.Close()
	i := 0
	for ; rows.Next(); i++ {
		
		var program DAO.Program
		var student DAO.StudentBasic
		
		err = sqlstruct.ScanAliased(&program, rows, "p")
		if err != nil {
			log.Println(err)
			return programs, DTO.ReturnMsg{-1, err.Error()}
		}
		err = sqlstruct.ScanAliased(&student, rows, "s")
		if err != nil {
			log.Println(err)
			return programs, DTO.ReturnMsg{-1, err.Error()}
		}
		if _, ok := programs[program.Name]; !ok {
			programs[program.Name] = program
		}
		appendStudentToProgram(programs, program.Name, student)
	}
	if i == 0 {
		return programs, DTO.ReturnMsg{1, translate.T("noProgram")}
	}
	return programs, DTO.ReturnMsg{0, ""}
}

func RetrieveProgram(id int) (DAO.Program, DTO.ReturnMsg) {
	var program DAO.Program
	
	sql := "SELECT %s, %s FROM program AS p " +
		"INNER JOIN program_subscription AS ps ON p.idprogram = ps.fk_program " +
		"INNER JOIN user AS s ON s.iduser = ps.fk_user WHERE p.idprogram = ?"
	sql = fmt.Sprintf(sql, sqlstruct.ColumnsAliased(DAO.Program{}, "p"), sqlstruct.ColumnsAliased(DAO.StudentBasic{}, "s"))
	rows, err := db.Query(sql, id)
	if err != nil {
		log.Println(err)
		return program, DTO.ReturnMsg{-1, err.Error()}
	}
	defer rows.Close()
	i := 0
	for ; rows.Next(); i++ {
		var student DAO.StudentBasic
		if i == 0 {
			err = sqlstruct.ScanAliased(&program, rows, "p")
			if err != nil {
				log.Println(err)
				return program, DTO.ReturnMsg{-1, err.Error()}
			}
		}
		err = sqlstruct.ScanAliased(&student, rows, "s")
		if err != nil {
			log.Println(err)
			return program, DTO.ReturnMsg{-1, err.Error()}
		}
		program.Enrolled_student = append(program.Enrolled_student, student)
	}
	if i == 0 {
		return program, DTO.ReturnMsg{1, translate.T("noProgramFor", map[string]interface{}{
			"Value": id,
			"Field": "id",
		})}
	}
	return program, DTO.ReturnMsg{0, ""}
}

func DeleteProgram(id int) (DTO.ReturnMsg) {
	stmt, err := db.Prepare("DELETE FROM program WHERE idprogram = ?")
	if err != nil {
		log.Println(err)
		return DTO.ReturnMsg{-1, err.Error()}
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return DTO.ReturnMsg{-1, err.Error()}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return DTO.ReturnMsg{-1, err.Error()}
	}
	if affect == 0 {
		return DTO.ReturnMsg{1, translate.T("noProgramFor", map[string]interface{}{
			"Value": id,
			"Field": "id",
		})}
	}
	return DTO.ReturnMsg{0, translate.T("programSuccessfullyDeleted")}
}

func appendStudentToProgram(p map[string]DAO.Program, n string, s DAO.StudentBasic) {
	tmp := p[n]
	tmp.Enrolled_student = append(p[n].Enrolled_student, s)
	p[n] = tmp
}