package models

type UserDb struct {
	DNI            int64 `gorm:"primary_key"`
	Name           string
	LastName       string
	QuantityMovies int8
}

func (UserDb) TableName() string {
	return "User"
}
