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
	reservations map[int32]Reservation
	nextID       func() int32
	screening    api.Screening_Service
	user         api.User_Service
}

func (service *Service) CreateReservation(ctx context.Context, req *api.CreateReservationReq, resp *api.CreateReservationResp) error {
	reservationID := service.nextID()
	_, err := service.screening.GetScreening(ctx, &api.GetScreeningReq{ScreeningID: req.GetScreeningID()})
	if err != nil {
		return err
	}
	_, err = service.user.GetUser(ctx, &api.GetUserReq{UserID: req.GetUserID()})
	if err != nil {
		return err
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
	reservation, ok := service.reservations[req.GetReservationID()]
	if !ok {
		return errors.NotFound("ERR-NO-RESERVATION", "Reservation (ID: %d) not found!", req.GetReservationID())
	}
	_, err := service.screening.ChangeFreeSeats(context.TODO(), &api.ChangeFreeSeatsReq{
		ScreeningID: reservation.screeningID,
		Change:      -reservation.seats,
	})
	if err != nil {
		return err
	}
	_, err = service.user.AddUserReservation(ctx, &api.AddUserReservationReq{UserID: reservation.userID})
	if err != nil {
		_, changeErr := service.screening.ChangeFreeSeats(context.TODO(), &api.ChangeFreeSeatsReq{
			ScreeningID: reservation.screeningID,
			Change:      reservation.seats,
		})
		if changeErr != nil {
			return errors.InternalServerError("ERR-ROLLBACK", "Error rolling back screening after user service failure, data is now inconsistent!\nPlease contact admin.")
		}
		return err
	}
	reservation.isActive = true
	service.reservations[req.GetReservationID()] = reservation
	return nil
}

func (service *Service) DeleteReservation(ctx context.Context, req *api.DeleteReservationReq, resp *api.DeleteReservationResp) error {
	reservation, ok := service.reservations[req.GetReservationID()]
	if !ok {
		return errors.NotFound("ERR-NO-RESERVATION", "Reservation (ID: %d) not found!", req.GetReservationID())
	}
	_, err := service.user.DeleteUserReservation(context.TODO(), &api.DeleteUserReservationReq{ReservationID: req.GetReservationID(), UserID: reservation.userID})
	if err != nil {
		return err
	}
	delete(service.reservations, req.GetReservationID())
	return nil
}

func (service *Service) DeleteReservationsWithScreening(ctx context.Context, req *api.DeleteReservationsWithScreeningReq, resp *api.DeleteReservationsWithScreeningResp) error {
	ids := make([]int32, 0)
	for id, reservation := range service.reservations {
		if reservation.screeningID == req.GetScreeningID() {
			ids = append(ids, id)
		}
	}
	for _, id := range ids {
		err := service.DeleteReservation(ctx, &api.DeleteReservationReq{ReservationID: id}, &api.DeleteReservationResp{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *Service) GetReservation(ctx context.Context, req *api.GetReservationReq, resp *api.GetReservationResp) error {
	reservation, ok := service.reservations[req.ReservationID]
	if !ok {
		return errors.NotFound("ERR-NO-RESERVATION", "Reservation (ID: %d) not found!", req.GetReservationID())
	}
	resp.ScreeningID = reservation.screeningID
	resp.UserID = reservation.userID
	resp.Active = reservation.isActive
	resp.NrOfSeats = reservation.seats
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
		reservations: make(map[int32]Reservation),
		nextID:       helpers.IDGenerator(),
		screening:    api.NewScreening_Service("screening", screening.Client()),
		user:         api.NewUser_Service("user", user.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
