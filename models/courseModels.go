package models

import "gorm.io/gorm"

// Department model
type Department struct {
	gorm.Model
	Name string
}

// Course model
type Course struct {
	gorm.Model
	Name           string
	DepartmentID   int
	Department     Department
	PrerequisiteID int
	Prerequisite   *Course
	Credits        int
}

// Program model - for example Masters in something
type Program struct {
	gorm.Model
	Name         string
	Courses      []Course `gorm:"many2many:program_courses;"`
	DepartmentID int
	Department   Department
	TotalCredits int
}
