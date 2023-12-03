package models

/* Period model - can be adjusted using bi-semester or tri-semester system.
For example, Fall, Summer and Spring for tri-semester system. */
type Period struct {
	Name string
}

// Semester model. For example, Spring 2015.
type Semester struct {
	// Semester model
	PeriodID int
	Period   Period
	Year     string
}

// Slot model. Defines a definite set of class times that a particular course can have.
type Slot struct {
	// Using string because we're not doing calculations with this.
	SlotTime string
}
