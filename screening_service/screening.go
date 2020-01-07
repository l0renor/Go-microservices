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
	_, err := service.movieService.GetMovie(ctx, &api.GetMovieMsg{Id: req.GetMovieID()})
	if err != nil {
		return err
	}
	roomResp, err := service.roomService.GetRoom(ctx, &api.GetRoomMsg{Id: req.GetRoomID()})
	if err != nil {
		return err
	}
	service.screenings[screeningID] = Screening{
		movieID:   req.GetMovieID(),
		roomID:    req.GetRoomID(),
		freeSeats: roomResp.Room.GetNrOfSeats(),
	}
	resp.ScreeningID = screeningID
	return nil
}

func (service *Service) ChangeFreeSeats(ctx context.Context, req *api.ChangeFreeSeatsReq, resp *api.ChangeFreeSeatsResp) error {
	screening, ok := service.screenings[req.GetScreeningID()]
	if !ok {
		return errors.NotFound("ERR-NO-SCREENING", "Screening (ID: %d) not found!", req.GetScreeningID())
	} else if screening.freeSeats+req.GetChange() < 0 {
		return errors.Conflict("ERR-FULL", "Screening (ID: %v) already has too many reservations!", req.GetScreeningID())
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

func (service *Service) DeleteRoom(ctx context.Context, req *api.DeleteRoomReq, resp *api.DeleteRoomResp) error {
	ids := make([]int32, 0)
	for id, screening := range service.screenings {
		if screening.roomID == req.GetRoomID() {
			ids = append(ids, id)
		}
	}
	for _, id := range ids {
		delete(service.screenings, id)
	}
	return nil
}

func (service *Service) DeleteMovie(ctx context.Context, req *api.DeleteMovieReq, resp *api.DeleteMovieResp) error {
	ids := make([]int32, 0)
	for id, screening := range service.screenings {
		if screening.movieID == req.GetMovieID() {
			ids = append(ids, id)
		}
	}
	for _, id := range ids {
		delete(service.screenings, id)
	}
	return nil
}

func (service *Service) GetScreening(ctx context.Context, req *api.GetScreeningReq, resp *api.GetScreeningResp) error {
	screening, ok := service.screenings[req.ScreeningID]
	if !ok {
		return errors.NotFound("ERR-NO-SCREENING", "Screening (ID: %d) not found!", req.GetScreeningID())
	}
	resp.MovieID = screening.movieID
	resp.RoomID = screening.roomID
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
