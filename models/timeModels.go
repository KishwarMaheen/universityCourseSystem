package models

import "gorm.io/gorm"

/*
	Period model - can be adjusted using bi-semester or tri-semester system.

For example, Fall, Summer and Spring for tri-semester system.
*/
type Period struct {
	gorm.Model
	Name string
}

// Semester model. For example, Spring 2015.
type Semester struct {
	gorm.Model
	PeriodID int
	Period   Period
	Year     string
}

// Slot model. Defines a definite set of class times that a particular course can have.
type Slot struct {
	gorm.Model

	// Using string because we're not doing calculations with this.
	SlotTime string
}
