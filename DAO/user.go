package DAO

type User struct {
	ID        uint      `gorm:"column:iduser"`
	FirstName string    `gorm:"column:forst_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
}
