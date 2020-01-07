package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"log"
)

type Client struct {
	roomService        api.Room_Service
	userService        api.User_Service
	reservationService api.Reservation_Service
	screening_service  api.Screening_Service
	movieService       api.Movie_Service
	ids                map[string]int32
}

/*
Sets up:
4 users
4 movies
2 rooms
4 screenings


*/
func (c Client) setup() {
	rsp, err := c.userService.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Oleg"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Oleg"] = rsp.UserID
	rsp, err = c.userService.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Brandl"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Brandl"] = rsp.UserID
	rsp, err = c.userService.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Fabi"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Fabi"] = rsp.UserID
	rsp, err = c.userService.CreateUser(context.TODO(), &api.CreateUserReq{Name: "SvenShulz"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["SvenShulz"] = rsp.UserID

	roomrsp, err := c.roomService.CreateRoom(context.TODO(), &api.CreateRoomMsg{
		Name:      "Mordor",
		NrOfSeats: 4,
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mordor"] = roomrsp.Id

	roomrsp, err = c.roomService.CreateRoom(context.TODO(), &api.CreateRoomMsg{
		Name:      "Isengard",
		NrOfSeats: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Isengard"] = roomrsp.Id

	moviersp, err := c.movieService.CreateMovie(context.TODO(), &api.CreateMovieMsg{Name: "Leon der Profi"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Leon der Profi"] = moviersp.Id
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieMsg{Name: "Mogli"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mogli"] = moviersp.Id
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieMsg{Name: "Die Zwei Türme"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Die Zwei Türme"] = moviersp.Id
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieMsg{Name: "Mitten im Leben der Film"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mitten im Leben der Film"] = moviersp.Id

	scrreningrsp, err := c.screening_service.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Mitten im Leben der Film"],
		RoomID:  c.ids["Isengard"],
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["1"] = scrreningrsp.ScreeningID

	scrreningrsp, err = c.screening_service.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Mogli"],
		RoomID:  c.ids["Isengard"],
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["2"] = scrreningrsp.ScreeningID

	scrreningrsp, err = c.screening_service.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Leon der Profi "],
		RoomID:  c.ids["Mordor"],
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["3"] = scrreningrsp.ScreeningID

	scrreningrsp, err = c.screening_service.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Die Zwei Türme"],
		RoomID:  c.ids["Mordor"],
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["4"] = scrreningrsp.ScreeningID

}

//Call after setup()
func (c Client) deletedRoom() {
	reservationrsp, err := c.reservationService.CreateReservation(context.TODO(), &api.CreateReservationReq{
		UserID:      c.ids["Oleg"],
		ScreeningID: c.ids["3"],
		NrOfSeats:   2,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.reservationService.ActivateReservation(context.TODO(), &api.ActivateReservationReq{ReservationID: reservationrsp.ReservationID})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.roomService.DeleteRoom(context.TODO(), &api.DeleteRoomMsg{Id: c.ids["Mordor"]})
	if err != nil {
		log.Fatal(err)
	}

	reservationsrsp, err := c.reservationService.GetReservations(context.TODO(), &api.GetReservationsReq{})

	for i := 0; i < len(reservationsrsp.Reservations); i++ {
		print(reservationsrsp.Reservations[i])
	}

}

func main() {

	service := micro.NewService()
	service.Init()

	// create the greeter client using the service name and client
	client := Client{
		roomService:        api.NewRoom_Service("room", service.Client()),
		userService:        api.NewUser_Service("user", service.Client()),
		reservationService: api.NewReservation_Service("reservation", service.Client()),
		screening_service:  api.NewScreening_Service("screening", service.Client()),
		movieService:       api.NewMovie_Service("movie", service.Client()),
		ids:                make(map[string]int32),
	}
	client.interact()

}
