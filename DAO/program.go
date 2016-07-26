package DAO


type ProgramCreation struct {
	Name                 string
	EnrolledStudents     []int
	NbrOfHoursToComplete int
}

type Program struct {
	Id                    int    `sql:"idprogram"`
	Name                  string `sql:"name"`
	Description           string `sql:"description"`
	Start_Date            string `sql:"start_date"`
	End_Date              string `sql:"end_date"`
	Nbr_hours_to_complete int    `sql:"nbr_hours_to_complete"`
	Enrolled_student      []StudentBasic  `sql:"-"`
}