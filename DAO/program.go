package DAO

type Program struct {
	Id                    int    `sql:"idprogram"`
	Name                  string `sql:"name"`
	Description           string `sql:"description"`
	Start_date            string `sql:"start_date"`
	End_date              string `sql:"end_date"`
	Nbr_hours_to_complete int    `sql:"nbr_hours_to_complete"`
	Enrolled_student      []StudentBasic  `sql:"-"`
}