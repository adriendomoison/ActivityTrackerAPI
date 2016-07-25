package DAO

type ProgramCreation struct {
	Name                 string
	EnrolledStudents     []int
	NbrOfHoursToComplete int
}

type Program struct {
	Id                    int    `sql:"idprogram"`
	Name                  string `sql:"name"`
	Nbr_event             int    `sql:"nbr_event"`
	Nbr_hours_available   int    `sql:"nbr_hours_available"`
	Nbr_student           int    `sql:"nbr_student"`
	Nbr_student_danger    int    `sql:"nbr_student_danger"`
	Nbr_hours_to_complete int    `sql:"nbr_hours_to_complete"`
	Enrolled_student      []StudentBasic `sql:"-"`
}

