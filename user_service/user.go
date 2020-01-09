package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"log"
)

type User struct {
	name         string
	reservations []int32
}

type Service struct {
	users  map[int32]User
	nextID func() int32
}

func (service *Service) CreateUser(ctx context.Context, req *api.CreateUserReq, resp *api.CreateUserResp) error {
	userID := service.nextID()
	service.users[userID] = User{
		name:         req.GetName(),
		reservations: make([]int32, 0),
	}
	resp.UserID = userID
	return nil
}

func (service *Service) DeleteUser(ctx context.Context, req *api.DeleteUserReq, resp *api.DeleteUserResp) error {
	user, ok := service.users[req.GetUserID()]
	if !ok {
		return errors.NotFound("ERR-NO-USER", "User (ID: %d) was not found!", req.GetUserID())
	}
	if len(user.reservations) > 0 {
		return errors.Conflict("ERR-RESERVATION-LEFT", "The user still has %d reservation(s) and can't be deleted!", len(user.reservations))
	}
	delete(service.users, req.GetUserID())
	return nil
}

func (service *Service) GetUser(ctx context.Context, req *api.GetUserReq, resp *api.GetUserResp) error {
	user, ok := service.users[req.GetUserID()]
	if !ok {
		return errors.NotFound("ERR-NO-USER", "User (ID: %d) not found!", req.GetUserID())
	}
	resp.Name = user.name
	return nil
}

func (service *Service) GetUsers(ctx context.Context, req *api.GetUsersReq, resp *api.GetUsersResp) error {
	users := make([]*api.GetUsersResp_User, 0)
	for userID, user := range service.users {
		users = append(users, &api.GetUsersResp_User{
			UserID: userID,
			Name:   user.name,
		})
	}
	resp.Users = users
	return nil
}

func (service *Service) AddUserReservation(ctx context.Context, req *api.AddUserReservationReq, resp *api.AddUserReservationResp) error {
	log.Print(req.GoString())
	user, ok := service.users[req.GetUserID()]
	if !ok {
		return errors.NotFound("ERR-NO-USER", "User (ID: %d) not found!", req.GetUserID())
	}
	existsAlready := contains(user.reservations, req.GetReservationID())
	if existsAlready {
		return errors.Conflict("ERR-RESERVATION-EXISTS", "Reservation (ID: %d) already exists!", req.GetReservationID())
	}
	user.reservations = append(user.reservations, req.GetReservationID())
	return nil
}

func (service *Service) DeleteUserReservation(ctx context.Context, req *api.DeleteUserReservationReq, resp *api.DeleteUserReservationResp) error {
	log.Print(service.users)
	user, ok := service.users[req.GetUserID()]
	if !ok {
		return errors.NotFound("ERR-NO-USER", "User (ID: %d) not found!", req.GetUserID())
	}
	log.Print(user.reservations)
	exists := remove(user.reservations, req.GetReservationID())
	if !exists {
		return errors.NotFound("ERR-NO-RESERVATION", "Reservation (ID: %d) not found!", req.GetReservationID())
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	service.Init()

	if err := api.RegisterUser_ServiceHandler(service.Server(), &Service{
		users:  make(map[int32]User),
		nextID: helpers.IDGenerator(0),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//true on success false if not present
func remove(s []int32, e int32) bool {
	for i, a := range s {
		if a == e {
			// Remove the element at index i from a.
			s[i] = s[len(s)-1] // Copy last element to index i.
			s[len(s)-1] = 0    // Erase last element (write zero value).
			s = s[:len(s)-1]   // Truncate slice.
			return true
		}
	}
	return false
}
