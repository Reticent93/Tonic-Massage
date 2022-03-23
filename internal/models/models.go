package models

import (
	"time"
)

//User is the user model
type User struct {
	ID          int
	Firstname   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Therapist is the therapist model
type Therapist struct {
	ID            int
	TherapistName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Reservation is the reservation model
type Reservation struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	StartDate   time.Time
	EndDate     time.Time
	TherapistId int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Therapist   Therapist
}

//RoomRestriction is a room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	TherapistID   int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Therapist     Therapist
	Reservation   Reservation
	Restriction   Restriction
}
