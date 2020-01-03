package user_service

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
)

type Service struct {
	users              map[int32]string
	nextID             func() int32
	reservationService api.Reservation_Service
}

func (service *Service) CreateUser(ctx context.Context, req *api.CreateUserReq, resp *api.CreateUserResp) error {
	userID := service.nextID()
	service.users[userID] = req.GetName()
	resp.UserID = userID
	return nil
}

func (service *Service) DeleteUser(ctx context.Context, req *api.DeleteUserReq, resp *api.DeleteUserResp) error {
	_, ok := service.users[req.GetUserID()]
	if ok {
		// TODO: Check reservations
		delete(service.users, req.GetUserID())
		resp.Success = true
	} else {
		resp.Success = false
	}
	return nil
}

func (service *Service) GetUser(ctx context.Context, req *api.GetUserReq, resp *api.GetUserResp) error {
	name, ok := service.users[req.GetUserID()]
	if ok {
		resp.Name = name
	} else {
		re
	}
	return nil
}

func (service *Service) GetUsers(ctx context.Context, req *api.GetUsersReq, resp *api.GetUsersResp) error {
	users := make([]*api.GetUsersResp_User, 0)
	for userID, name := range service.users {
		users = append(users, &api.GetUsersResp_User{
			UserID: userID,
			Name:   name,
		})
	}
	resp.Users = users
}
