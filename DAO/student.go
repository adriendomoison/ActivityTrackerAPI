package DAO

type StudentBasic struct {
	Id        int    `sql:"iduser"`
	FirstName string `sql:"first_name"`
	LastName  string `sql:"last_name"`
}
