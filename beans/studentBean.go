package beans

import (
	"time"
)

type Student struct {
	ID        int    `json:"id gorm:"primary_key""`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	// YearOfJoining time.    `json:"year_of_joining"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	CreatedAt  time.Time `json:"created_at"`
	ApprovedAt time.Time `json:"approved_at"`
	RejectedAt time.Time `json:"rejected_at"`
}
