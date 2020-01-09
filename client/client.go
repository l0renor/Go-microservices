package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"log"
	"time"
)

type Client struct {
	roomService        api.Room_Service
	userService        api.User_Service
	reservationService api.Reservation_Service
	screeningService   api.Screening_Service
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

	log.Print("Users created")

	roomrsp, err := c.roomService.CreateRoom(context.TODO(), &api.CreateRoomReq{
		Name:      "Mordor",
		NrOfSeats: 4,
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mordor"] = roomrsp.RoomID

	roomrsp, err = c.roomService.CreateRoom(context.TODO(), &api.CreateRoomReq{
		Name:      "Isengard",
		NrOfSeats: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Isengard"] = roomrsp.RoomID

	log.Print("Rooms created")

	moviersp, err := c.movieService.CreateMovie(context.TODO(), &api.CreateMovieReq{Name: "Leon der Profi"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Leon der Profi"] = moviersp.MovieID
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieReq{Name: "Mogli"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mogli"] = moviersp.MovieID
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieReq{Name: "Die Zwei Türme"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Die Zwei Türme"] = moviersp.MovieID
	moviersp, err = c.movieService.CreateMovie(context.TODO(), &api.CreateMovieReq{Name: "Mitten im Leben der Film"})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["Mitten im Leben der Film"] = moviersp.MovieID

	log.Print("Movies created")

	time.Sleep(1 * time.Second)

	scrreningrsp, err := c.screeningService.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Mitten im Leben der Film"],
		RoomID:  c.ids["Isengard"],
	})
	if err != nil {
		log.Fatal(err)
	}

	c.ids["1"] = scrreningrsp.ScreeningID
	log.Print("Screening 1 created")
	log.Print(c.ids)
	scrreningrsp, err = c.screeningService.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Mogli"],
		RoomID:  c.ids["Isengard"],
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Screening 2 created")
	c.ids["2"] = scrreningrsp.ScreeningID

	scrreningrsp, err = c.screeningService.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Leon der Profi"],
		RoomID:  c.ids["Mordor"],
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Screening 3 created")
	c.ids["3"] = scrreningrsp.ScreeningID

	scrreningrsp, err = c.screeningService.CreateScreening(context.TODO(), &api.CreateScreeningReq{
		MovieID: c.ids["Die Zwei Türme"],
		RoomID:  c.ids["Mordor"],
	})
	if err != nil {
		log.Fatal(err)
	}
	c.ids["4"] = scrreningrsp.ScreeningID
	time.Sleep(1 * time.Second)
	log.Print("-------------- Setup done -----------------")
}

//Call after setup()
func (c Client) deletedRoom() {
	log.Print("---------------- Deleting Room -----------------")
	createReservationResp, err := c.reservationService.CreateReservation(context.TODO(), &api.CreateReservationReq{
		UserID:      c.ids["Oleg"],
		ScreeningID: c.ids["3"],
		NrOfSeats:   2,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Reservation (ID: %v)", createReservationResp.GetReservationID())
	_, err = c.reservationService.ActivateReservation(context.TODO(), &api.ActivateReservationReq{ReservationID: createReservationResp.GetReservationID()})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Activated Reservation (ID: %v)", createReservationResp.GetReservationID())
	_, err = c.roomService.DeleteRoom(context.TODO(), &api.DeleteRoomReq{RoomID: c.ids["Mordor"]})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted Room (ID: %v), checking screening and reservation service for consistency", c.ids["Mordor"])
	_, err = c.screeningService.GetScreening(context.TODO(), &api.GetScreeningReq{ScreeningID: c.ids["3"]})
	if err == nil {
		log.Fatal("Screening still present although room was deleted!")
	}
	_, err = c.reservationService.GetReservation(context.TODO(), &api.GetReservationReq{ReservationID: createReservationResp.GetReservationID()})
	if err == nil {
		log.Fatal("Reservation still present although screening was deleted!")
	}
	log.Print("Success, all services are consistent")
}

// Call after setup()
func (c Client) conflictReservation() {
	log.Print("---------------- Conflicting Reservations -----------------")
	createReservationResp, err := c.reservationService.CreateReservation(context.TODO(), &api.CreateReservationReq{
		UserID:      c.ids["Oleg"],
		ScreeningID: c.ids["1"],
		NrOfSeats:   2,
	})
	if err != nil {
		log.Fatal(err)
	}

	createReservationResp2, err := c.reservationService.CreateReservation(context.TODO(), &api.CreateReservationReq{
		UserID:      c.ids["Fabi"],
		ScreeningID: c.ids["1"],
		NrOfSeats:   2,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("2 Reservations created")
	_, err = c.reservationService.ActivateReservation(context.TODO(), &api.ActivateReservationReq{ReservationID: createReservationResp.ReservationID})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("1st Reservation activated")
	_, err = c.reservationService.ActivateReservation(context.TODO(), &api.ActivateReservationReq{ReservationID: createReservationResp2.ReservationID})
	if err == nil {
		log.Fatal("Error: Was able to activate 2nd reservations although there are no more free seats!")
	}
	log.Print("Got expected error from activating 2nd reservation")
	reservationsrsp, err := c.reservationService.GetReservations(context.TODO(), &api.GetReservationsReq{})
	for i := 0; i < len(reservationsrsp.Reservations); i++ {
		log.Printf("Reservation| ID: %v, screeningID: %v, nrSeats: %v , active %v", i, reservationsrsp.Reservations[i].ScreeningID, reservationsrsp.Reservations[i].NrOfSeats, reservationsrsp.Reservations[i].Active)
	}

}

func main() {

	serviceMovie := micro.NewService()
	serviceMovie.Init()

	serviceUser := micro.NewService()
	serviceUser.Init()

	serviceRoom := micro.NewService()
	serviceRoom.Init()

	serviceReservation := micro.NewService()
	serviceReservation.Init()

	serviceScreening := micro.NewService()
	serviceScreening.Init()

	// create the greeter client using the service name and client
	client := Client{
		roomService:        api.NewRoom_Service("room", serviceRoom.Client()),
		userService:        api.NewUser_Service("user", serviceUser.Client()),
		reservationService: api.NewReservation_Service("reservation", serviceReservation.Client()),
		screeningService:   api.NewScreening_Service("screening", serviceScreening.Client()),
		movieService:       api.NewMovie_Service("movie", serviceMovie.Client()),
		ids:                make(map[string]int32),
	}

	client.setup()
	client.conflictReservation()
	client.deletedRoom()
}
