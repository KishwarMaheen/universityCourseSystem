package models

import "gorm.io/gorm"

type Department struct {
	// Department model
	gorm.Model
	Name string
}

type Course struct {
	// Course model
	gorm.Model
	Name           string
	DepartmentID   int
	Department     Department
	PrerequisiteID int
	Prerequisite   *Course
	Credits        int
}

type Program struct {
	// Program model - for example Masters in something
	gorm.Model
	Name         string
	Courses      []Course `gorm:"many2many:program_courses;"`
	DepartmentID int
	Department   Department
	TotalCredits int
}
