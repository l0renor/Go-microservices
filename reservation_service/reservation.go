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
		reservation.isActive = true
		service.reservations[req.ReservationID] = reservation
		resp.Success = true
	} else {
		resp.Success = false
	}
}

func (service *Service) DeleteReservation(ctx context.Context, req *api.DeleteReservationReq, resp *api.DeleteReservationResp) {
	_, ok := service.reservations[req.ReservationID]
	if ok {
		delete(service.reservations, req.ReservationID)
		resp.Success = true
	} else {
		resp.Success = false
	}
}

func (service *Service) GetReservation(ctx context.Context, req *api.GetReservationReq, resp *api.GetReservationResp) {
	reservation, ok := service.reservations[req.ReservationID]
	if ok {
		resp.ScreeningID = reservation.screeningID
		resp.UserID = reservation.userID
		resp.Active = reservation.isActive
		resp.NrOfSeats = reservation.seats
	}
}

func (service *Service) GetReservations(ctx context.Context, req *api.GetReservationsReq, resp *api.GetReservationsResp) {
	reservations := make([]*api.GetReservationResp, 0)
	for _, reservation := range service.reservations {
		reservations = append(reservations, &api.GetReservationResp{
			ScreeningID: reservation.screeningID,
			UserID:      reservation.userID,
			Active:      reservation.isActive,
			NrOfSeats:   reservation.seats,
		})
	}
	resp.Reservations = reservations
}
