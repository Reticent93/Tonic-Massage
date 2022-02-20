package models

import (
	"time"
)

//Booking holds booking data
type Booking struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Date      string
	Time      string
}

//Users is the users model
type Users struct {
	ID          int
	Firstname   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Therapists is the therapist model
type Therapists struct {
	ID            int
	TherapistName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//Restrictions is the restriction model
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Reservations is the reservation model
type Reservations struct {
	ID          int
	Firstname   string
	LastName    string
	Email       string
	Phone       string
	StartDate   time.Time
	EndDate     time.Time
	TherapistId int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Therapist   Therapists
}

//RoomRestrictions is a room restriction model
type RoomRestrictions struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	TherapistID   int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Therapist     Therapists
	Reservation   Reservations
	Restriction   Restrictions
}
