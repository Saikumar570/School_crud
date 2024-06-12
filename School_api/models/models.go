package models

import "github.com/jinzhu/gorm"

type School struct {
	gorm.Model
	Name          string
	SchoolId      string
	SchoolAddress Address `gorm:"foreignkey:SchoolID"`
	Classes       Classes `gorm:"foreignkey:SchoolID"`
}

type Classes struct {
	gorm.Model
	SchoolID uint
	Class6   Class6 `gorm:"foreignkey:ClassesID"`
	Class7   Class7 `gorm:"foreignkey:ClassesID"`
	Class8   Class8 `gorm:"foreignkey:ClassesID"`
}

type Class6 struct {
	gorm.Model
	ClassesID uint
	ClassName string
	Students  []Student `gorm:"foreignkey:Class6ID"`
}

type Class7 struct {
	gorm.Model
	ClassesID uint
	ClassName string
	Students  []Student `gorm:"foreignkey:Class7ID"`
}

type Class8 struct {
	gorm.Model
	ClassesID uint
	ClassName string
	Students  []Student `gorm:"foreignkey:Class8ID"`
}

type Student struct {
	gorm.Model
	Name       string
	Percentage int
	Address    Address `gorm:"foreignkey:StudentID"`
	Class6ID   uint
	Class7ID   uint
	Class8ID   uint
}

type Address struct {
	gorm.Model
	Street    string
	City      string
	State     string
	SchoolID  uint
	StudentID uint
}
