package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"log"
)

type Reservation struct {
	screeningID int32
	userID      int32
	seats       int32
	isActive    bool
}

type Service struct {
	reservations     map[int32]Reservation
	nextID           func() int32
	screeningService api.Screening_Service
	userService      api.User_Service
}

func (service *Service) CreateReservation(ctx context.Context, req *api.CreateReservationReq, resp *api.CreateReservationResp) error {
	reservationID := service.nextID()
	_, err := service.screeningService.GetScreening(ctx, &api.GetScreeningReq{ScreeningID: req.ScreeningID})
	if err != nil {
		return errors.NotFound("screening_not_found", "screening(ID: %v not found", req.ScreeningID)
	}
	_, err = service.userService.GetUser(ctx, &api.GetUserReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.NotFound("user_not_found", "user(ID: %v not found", req.UserID)
	}
	service.reservations[reservationID] = Reservation{
		screeningID: req.GetScreeningID(),
		userID:      req.GetUserID(),
		isActive:    false,
	}
	resp.ReservationID = reservationID
	return nil
}

func (service *Service) ActivateReservation(ctx context.Context, req *api.ActivateReservationReq, resp *api.ActivateReservationResp) error {
	reservation, ok := service.reservations[req.ReservationID]
	if ok {
		screeningRsp, err := service.screeningService.GetScreening(ctx, &api.GetScreeningReq{ScreeningID: reservation.screeningID})
		if err != nil {
			return errors.NotFound("screening_not_found", "screening(ID: %v not found", reservation.screeningID)
		}
		if screeningRsp.NrOfFreeSeats < reservation.seats {
			return errors.Conflict("Not_enough_Seats", "Not_enough_Seats needed %v; free %v", screeningRsp.NrOfFreeSeats, reservation.seats)
		}
		reservation.isActive = true
		service.reservations[req.ReservationID] = reservation
		_, err = service.userService.AddReservation(ctx, &api.AddReservationReq{
			UserID: reservation.userID,
		})
		if err != nil {
			return errors.NotFound("Error reserving in user profile", "error reserving in user profile")
		}
	} else {
		return errors.InternalServerError("Reservation activation failed", "Reservation activation failed id %v", req.ReservationID)
	}
	return nil
}

func (service *Service) DeleteReservation(ctx context.Context, req *api.DeleteReservationReq, resp *api.DeleteReservationResp) error {
	_, ok := service.reservations[req.ReservationID]
	if ok {
		delete(service.reservations, req.ReservationID)

	} else {
		return errors.NotFound("Reservation_not_found", "Reservation(%v)not_found", req.ReservationID)
	}
	return nil
}

func (service *Service) GetReservation(ctx context.Context, req *api.GetReservationReq, resp *api.GetReservationResp) error {
	reservation, ok := service.reservations[req.ReservationID]
	if ok {
		resp.ScreeningID = reservation.screeningID
		resp.UserID = reservation.userID
		resp.Active = reservation.isActive
		resp.NrOfSeats = reservation.seats
		return errors.NotFound("Reservation_not_found", "Reservation(%v)not_found", req.ReservationID)
	}
	return nil
}

func (service *Service) GetReservations(ctx context.Context, req *api.GetReservationsReq, resp *api.GetReservationsResp) error {
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
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("reservation"),
		micro.Version("latest"),
	)

	screening := micro.NewService()
	screening.Init()

	user := micro.NewService()
	user.Init()

	service.Init()

	if err := api.RegisterReservation_ServiceHandler(service.Server(), &Service{
		reservations:     make(map[int32]Reservation),
		nextID:           helpers.IDGenerator(),
		screeningService: api.NewScreening_Service("screening", screening.Client()),
		userService:      api.NewUser_Service("user", user.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
