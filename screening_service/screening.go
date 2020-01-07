package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"log"
)

type Screening struct {
	movieID   int32
	roomID    int32
	freeSeats int32
}

type Service struct {
	screenings   map[int32]Screening
	nextID       func() int32
	roomService  api.Room_Service
	movieService api.Movie_Service
}

func (service *Service) CreateScreening(ctx context.Context, req *api.CreateScreeningReq, resp *api.CreateScreeningResp) error {
	screeningID := service.nextID()
	_, err := service.movieService.GetMovie(ctx, &api.GetMovieMsg{Id: req.MovieID})
	if err != nil {
		return errors.NotFound("movie_not_found", "movie(ID: %v not found", req.MovieID)
	}
	_, err = service.roomService.GetRoom(ctx, &api.GetRoomMsg{Id: req.RoomID})
	if err != nil {
		return errors.NotFound("room_not_found", "room(ID: %v not found", req.RoomID)
	}
	service.screenings[screeningID] = Screening{
		movieID:   req.GetMovieID(),
		roomID:    req.GetRoomID(),
		freeSeats: 0, // TODO: Get nr of Seats in Room
	}
	resp.ScreeningID = screeningID
	return nil
}

func (service *Service) ChangeFreeSeats(ctx context.Context, req *api.ChangeFreeSeatsReq, resp *api.ChangeFreeSeatsResp) error {
	screening, ok := service.screenings[req.GetScreeningID()]
	if !ok {
		return errors.NotFound("screening_not_found", "screening (ID: %v) not found", req.GetScreeningID())
	} else if screening.freeSeats+req.GetChange() < 0 {
		return errors.Conflict("already_reserved", "Screening (ID: %v) has already too much reservations", req.GetScreeningID())
	} else {
		screening.freeSeats += req.GetChange()
		return nil
	}
}

func (service *Service) DeleteScreening(ctx context.Context, req *api.DeleteScreeningReq, resp *api.DeleteScreeningResp) error {
	delete(service.screenings, req.ScreeningID)
	// TODO: Notify reservations
	return nil
}

func (service *Service) GetScreening(ctx context.Context, req *api.GetScreeningReq, resp *api.GetScreeningResp) error {
	screening, ok := service.screenings[req.ScreeningID]
	if ok {
		resp.MovieID = screening.movieID
		resp.RoomID = screening.roomID
	} else {
		return errors.NotFound("screening_not_found", "screening(ID: %v not found", req.ScreeningID)
	}
	return nil
}

func (service *Service) GetScreenings(ctx context.Context, req *api.GetScreeningsReq, resp *api.GetScreeningsResp) error {
	screenings := make([]*api.GetScreeningResp, 0)
	for _, screening := range service.screenings {
		screenings = append(screenings, &api.GetScreeningResp{
			MovieID: screening.movieID,
			RoomID:  screening.roomID,
		})
	}
	resp.Screenings = screenings
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("screening"),
		micro.Version("latest"),
	)

	room := micro.NewService()
	room.Init()

	movie := micro.NewService()
	movie.Init()

	service.Init()

	if err := api.RegisterScreening_ServiceHandler(service.Server(), &Service{
		screenings:   make(map[int32]Screening),
		nextID:       helpers.IDGenerator(),
		roomService:  api.NewRoom_Service("room", room.Client()),
		movieService: api.NewMovie_Service("movie", movie.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
