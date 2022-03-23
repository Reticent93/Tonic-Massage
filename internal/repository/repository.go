package repository

import "github.com/Reticent93/Tonic-Massage/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
