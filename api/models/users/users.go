package users

// Users - slice of Usr
type Users []User

// User - model of user from the db
type User struct {
	ID       int64  `xorm:"id"`
	Email    string `xorm:"email"`
	Password string `xorm:"password"`
}

// TableName - sets the corect table name
func (u *User) TableName() string {
	return "users"
}
