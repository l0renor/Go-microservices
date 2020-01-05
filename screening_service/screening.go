package screening_service

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
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
	// TODO: Check if Movie and Romm ID are valid
	_, err := service.movieService.GetMovie(ctx, &api.GetMovieMsg{Id: req.MovieID})
	if err != nil {
		//TODO error handling
	}
	_, err = service.roomService.GetRoom(ctx, &api.GetRoomMsg{Id: req.RoomID})
	if err != nil {
		//TODO error handling
	}
	service.screenings[screeningID] = Screening{
		movieID:   req.GetMovieID(),
		roomID:    req.GetRoomID(),
		freeSeats: 0, // TODO: Get nr of Seats in Room
	}
	resp.ScreeningID = screeningID
	return nil
}

func (service *Service) DeleteScreening(ctx context.Context, req *api.DeleteScreeningReq, resp *api.DeleteScreeningResp) error {
	// TODO: Check reservations
	if true {
		delete(service.screenings, req.ScreeningID)
		resp.Success = true
	} else {
		resp.Success = false
	}
	return nil
}

func (service *Service) GetScreening(ctx context.Context, req *api.GetScreeningReq, resp *api.GetScreeningResp) error {
	screening, ok := service.screenings[req.ScreeningID]
	if ok {
		resp.MovieID = screening.movieID
		resp.RoomID = screening.roomID
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
