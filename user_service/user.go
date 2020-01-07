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
		name: req.GetName(),
	}
	resp.UserID = userID
	return nil
}

func (service *Service) DeleteUser(ctx context.Context, req *api.DeleteUserReq, resp *api.DeleteUserResp) error {
	user, ok := service.users[req.GetUserID()]
	if ok {
		if len(user.reservations) > 0 {
			return errors.Conflict("usr_has_res", "The user still has reservations;can't be deleted")
		}
		delete(service.users, req.GetUserID())
	} else {
		return errors.NotFound("usr_not_found", "User can't be deleted not found")
	}
	return nil
}

func (service *Service) GetUser(ctx context.Context, req *api.GetUserReq, resp *api.GetUserResp) error {
	user, ok := service.users[req.GetUserID()]
	if ok {
		resp.Name = user.name
	} else {
		return errors.NotFound("usr_not_found", "User not found")
	}
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

func (service *Service) AddReservation(ctx context.Context, req *api.AddReservationReq, resp *api.AddReservationResp) error {
	// check user exists
	user, ok := service.users[req.UserID]
	if !ok {
		return errors.NotFound("user_not_found", "User not found")
	}
	// check reservation not exists
	existsAlready := contains(user.reservations, req.ReservationID)
	if existsAlready {
		return errors.Conflict("reservation_exists", "This reservation already exists")
	}
	user.reservations = append(user.reservations, req.ReservationID)
	return nil
}

func (service *Service) DeleteReservation(ctx context.Context, req *api.DeleteReservationReq, resp *api.DeleteReservationResp) error {
	user, ok := service.users[req.UserID]
	if !ok {
		return errors.NotFound("user_not_found", "User not found")
	}
	exists := remove(user.reservations, req.ReservationID)
	if !exists {
		return errors.NotFound("reservation_not_found", "Reservation not found")
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
		nextID: helpers.IDGenerator(),
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
