package reservation_service

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
)

type Reservation struct {
	screeningID int32
	userID      int32
	seats       int32
	isActive    bool
}

type Service struct {
	reservations map[int32]Reservation
	nextID       func() int32
}

func (service *Service) CreateReservation(ctx context.Context, req *api.CreateReservationReq, resp *api.CreateReservationResp) {
	reservationID := service.nextID()
	// TODO: Check if Screening and User ID are valid
	service.reservations[reservationID] = Reservation{
		screeningID: req.GetScreeningID(),
		userID:      req.GetUserID(),
		isActive:    false,
	}
	resp.ReservationID = reservationID
}

func (service *Service) ActivateReservation(ctx context.Context, req *api.ActivateReservationReq, resp *api.ActivateReservationResp) {
	reservation, ok := service.reservations[req.ReservationID]
	if ok {
		// TODO: Check if screening still has free seats and book them

	}
}
