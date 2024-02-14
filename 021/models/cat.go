package models

type Cat struct {
	Id       uint `gorm:"primaryKey"`
	Cat_Name string
	Cat_Type string
}
