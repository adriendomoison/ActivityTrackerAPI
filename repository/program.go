package repository

import (
	"github.com/adriendomoison/ActivityTrackerAPI/DAO"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
)

func CreateProgram(p DAO.Program) utils.ReturnMsg {
	if err := db.Create(&p).Error; err != nil {
		log.Println(err)
		return utils.ReturnMsg{1, translate.T("programAlreadyExist")}
	}
	return utils.ReturnMsg{0, translate.T("programCreated")}
}

func RetrieveAllPrograms() (programs []DAO.Program, rMsg utils.ReturnMsg) {
	
	//var programs = make(map[string]DAO.Program)
	//// TODO: patch the request so it do a concatenation and get programs even when nobody subscribed to it.
	//sql := "SELECT %s, %s FROM program AS p " +
	//	"INNER JOIN program_subscription AS ps ON p.idprogram = ps.fk_program " +
	//	"INNER JOIN user AS s ON s.iduser = ps.fk_user"
	//sql = fmt.Sprintf(sql, sqlstruct.ColumnsAliased(DAO.Program{}, "p"), sqlstruct.ColumnsAliased(DAO.StudentBasic{}, "s"))
	//rows, err := db.Query(sql)
	//if err != nil {
	//	log.Println(err)
	//	return programs, DTO.ReturnMsg{-1, err.Error()}
	//}
	//defer rows.Close()
	//i := 0
	//for ; rows.Next(); i++ {
	//
	//	var program DAO.Program
	//	var student DAO.StudentBasic
	//
	//	err = sqlstruct.ScanAliased(&program, rows, "p")
	//	if err != nil {
	//		log.Println(err)
	//		return programs, DTO.ReturnMsg{-1, err.Error()}
	//	}
	//	err = sqlstruct.ScanAliased(&student, rows, "s")
	//	if err != nil {
	//		log.Println(err)
	//		return programs, DTO.ReturnMsg{-1, err.Error()}
	//	}
	//	if _, ok := programs[program.Name]; !ok {
	//		programs[program.Name] = program
	//	}
	//	appendStudentToProgram(programs, program.Name, student)
	//}
	//if i == 0 {
	//	return programs, DTO.ReturnMsg{1, translate.T("noProgram")}
	//}
	return programs, utils.ReturnMsg{0, ""}
}

func RetrieveProgram(id int) (program DAO.Program, rMsg utils.ReturnMsg) {
	//var program DAO.Program
	//
	//sql := "SELECT %s, %s FROM program AS p " +
	//	"INNER JOIN program_subscription AS ps ON p.idprogram = ps.fk_program " +
	//	"INNER JOIN user AS s ON s.iduser = ps.fk_user WHERE p.idprogram = ?"
	//sql = fmt.Sprintf(sql, sqlstruct.ColumnsAliased(DAO.Program{}, "p"), sqlstruct.ColumnsAliased(DAO.StudentBasic{}, "s"))
	//rows, err := db.Query(sql, id)
	//if err != nil {
	//	log.Println(err)
	//	return program, DTO.ReturnMsg{-1, err.Error()}
	//}
	//defer rows.Close()
	//i := 0
	//for ; rows.Next(); i++ {
	//	var student DAO.StudentBasic
	//	if i == 0 {
	//		err = sqlstruct.ScanAliased(&program, rows, "p")
	//		if err != nil {
	//			log.Println(err)
	//			return program, DTO.ReturnMsg{-1, err.Error()}
	//		}
	//	}
	//	err = sqlstruct.ScanAliased(&student, rows, "s")
	//	if err != nil {
	//		log.Println(err)
	//		return program, DTO.ReturnMsg{-1, err.Error()}
	//	}
	//	program.Enrolled_student = append(program.Enrolled_student, student)
	//}
	//if i == 0 {
	//	return program, DTO.ReturnMsg{1, translate.T("noProgramFor", map[string]interface{}{
	//		"Value": id,
	//		"Field": "id",
	//	})}
	//}
	return program, utils.ReturnMsg{0, ""}
}

func DeleteProgram(id int) (rMsg utils.ReturnMsg) {
	//stmt, err := db.Prepare("DELETE FROM program WHERE idprogram = ?")
	//if err != nil {
	//	log.Println(err)
	//	return DTO.ReturnMsg{-1, err.Error()}
	//}
	//res, err := stmt.Exec(id)
	//if err != nil {
	//	log.Println(err)
	//	return DTO.ReturnMsg{-1, err.Error()}
	//}
	//affect, err := res.RowsAffected()
	//if err != nil {
	//	log.Println(err)
	//	return DTO.ReturnMsg{-1, err.Error()}
	//}
	//if affect == 0 {
	//	return DTO.ReturnMsg{1, translate.T("noProgramFor", map[string]interface{}{
	//		"Value": id,
	//		"Field": "id",
	//	})}
	//}
	return utils.ReturnMsg{0, translate.T("programSuccessfullyDeleted")}
}
//
//func appendStudentToProgram(p map[string]DAO.Program, n string, s DAO.StudentBasic) {
//	tmp := p[n]
//	tmp.Enrolled_student = append(p[n].Enrolled_student, s)
//	p[n] = tmp
//}