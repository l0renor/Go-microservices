package screening_service

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
)

type Screening struct {
	movieID int32
	roomID  int32
}

type Service struct {
	screenings map[int32]Screening
	nextID     func() int32
}

func (service *Service) CreateScreening(ctx context.Context, req *api.CreateScreeningReq, resp *api.CreateScreeningResp) {
	screeningID := service.nextID()
	service.screenings[screeningID] = Screening{
		movieID: req.GetMovieID(),
		roomID:  req.GetRoomID(),
	}
	resp.ScreeningID = screeningID
}

func (service *Service) DeleteScreening(ctx context.Context, req *api.DeleteScreeningReq, resp *api.DeleteScreeningResp) {
	// TODO: Check reservations
	if true {
		delete(service.screenings, req.ScreeningID)
		resp.Success = true
	} else {
		resp.Success = false
	}
}

func (service *Service) GetScreening(ctx context.Context, req *api.GetScreeningReq, resp *api.GetScreeningResp) {
	screening, ok := service.screenings[req.ScreeningID]
	if ok {
		resp.MovieID = screening.movieID
		resp.RoomID = screening.roomID
	}
}

func (service *Service) GetScreenings(ctx context.Context, req *api.GetScreeningsReq, resp *api.GetScreeningsResp) {
	var screenings []*api.GetScreeningResp
	for _, screening := range service.screenings {
		screenings = append(screenings, &api.GetScreeningResp{
			MovieID: screening.movieID,
			RoomID:  screening.roomID,
		})
	}
	resp.Screenings = screenings
}
