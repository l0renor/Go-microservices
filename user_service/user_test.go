package main

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	service := Service{users: map[int32]User{}, nextID: helpers.IDGenerator(0)}

	resp := &api.CreateUserResp{}
	if err := service.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Alf"}, resp); err != nil {
		t.Errorf("CreateUser() error = %v, wantErr %v", err, false)
	}
	if resp.GetUserID() != 0 {
		t.Errorf("CreateUser() returned ID = %v, wantID %v", resp.GetUserID(), 0)
	}

	resp = &api.CreateUserResp{}
	if err := service.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Bert"}, resp); err != nil {
		t.Errorf("CreateUser() error = %v, wantErr %v", err, false)
	}
	if resp.GetUserID() != 1 {
		t.Errorf("CreateUser() returned ID = %v, wantID %v", resp.GetUserID(), 1)
	}

	resp = &api.CreateUserResp{}
	if err := service.CreateUser(context.TODO(), &api.CreateUserReq{Name: "Alf"}, resp); err != nil {
		t.Errorf("CreateUser() error = %v, wantErr %v", err, false)
	}
	if resp.GetUserID() != 2 {
		t.Errorf("CreateUser() returned ID = %v, wantID %v", resp.GetUserID(), 2)
	}
}

func TestService_DeleteUser(t *testing.T) {
	service := Service{users: map[int32]User{0: {name: "Alf", reservations: []int32{9}}, 1: {name: "Bert", reservations: []int32{}}}, nextID: helpers.IDGenerator(2)}
	resp := &api.DeleteUserResp{}

	if err := service.DeleteUser(context.TODO(), &api.DeleteUserReq{UserID: 1}, resp); err != nil {
		t.Errorf("DeleteUser() error = %v, wantErr %v", err, false)
	}

	if err := service.DeleteUser(context.TODO(), &api.DeleteUserReq{UserID: 1}, resp); err == nil {
		t.Errorf("DeleteUser() error = %v, wantErr %v", err, true)
	}

	if err := service.DeleteUser(context.TODO(), &api.DeleteUserReq{UserID: 0}, resp); err == nil {
		t.Errorf("DeleteUser() error = %v, wantErr %v", err, true)
	}
}

func TestService_GetUser(t *testing.T) {
	service := Service{users: map[int32]User{0: {name: "Alf", reservations: []int32{9}}, 1: {name: "Bert", reservations: []int32{}}}, nextID: helpers.IDGenerator(2)}

	resp := &api.GetUserResp{}
	if err := service.GetUser(context.TODO(), &api.GetUserReq{UserID: 0}, resp); err != nil {
		t.Errorf("GetUser() error = %v, wantErr %v", err, false)
	}
	if resp.GetName() != service.users[0].name {
		t.Errorf("GetUser(%v) returned name = %v, wantName %v", 0, resp.GetName(), service.users[0].name)
	}

	resp = &api.GetUserResp{}
	if err := service.GetUser(context.TODO(), &api.GetUserReq{UserID: 2}, resp); err == nil {
		t.Errorf("GetUser() error = %v, wantErr %v", err, true)
	}
}

func TestService_GetUsers(t *testing.T) {
	service := Service{users: map[int32]User{0: {name: "Alf", reservations: []int32{9}}, 1: {name: "Bert", reservations: []int32{}}}, nextID: helpers.IDGenerator(2)}

	resp := &api.GetUsersResp{}
	if err := service.GetUsers(context.TODO(), &api.GetUsersReq{}, resp); err != nil {
		t.Errorf("GetUsers() error = %v, wantErr %v", err, false)
	}
	if len(resp.Users) != 2 {
		t.Errorf("GetUsers() returned %v users, wantedUser %v", len(resp.Users), 2)
	}
}

func TestService_AddUserReservation(t *testing.T) {
	service := Service{users: map[int32]User{0: {name: "Alf", reservations: []int32{9}}, 1: {name: "Bert", reservations: []int32{}}}, nextID: helpers.IDGenerator(2)}
	resp := &api.AddUserReservationResp{}

	if err := service.AddUserReservation(context.TODO(), &api.AddUserReservationReq{UserID: 1, ReservationID: 8}, resp); err != nil {
		t.Errorf("AddUserReservation() error = %v, wantErr %v", err, false)
	}

	if err := service.AddUserReservation(context.TODO(), &api.AddUserReservationReq{UserID: 0, ReservationID: 9}, resp); err == nil {
		t.Errorf("AddUserReservation() error = %v, wantErr %v", err, true)
	}
}

func TestService_DeleteUserReservation(t *testing.T) {
	service := Service{users: map[int32]User{0: {name: "Alf", reservations: []int32{9}}, 1: {name: "Bert", reservations: []int32{}}}, nextID: helpers.IDGenerator(2)}
	resp := &api.DeleteUserReservationResp{}

	if err := service.DeleteUserReservation(context.TODO(), &api.DeleteUserReservationReq{UserID: 0, ReservationID: 9}, resp); err != nil {
		t.Errorf("DeleteUserReservation() error = %v, wantErr %v", err, false)
	}

	if err := service.DeleteUserReservation(context.TODO(), &api.DeleteUserReservationReq{UserID: 2, ReservationID: 9}, resp); err == nil {
		t.Errorf("DeleteUserReservation() error = %v, wantErr %v", err, true)
	}

	if err := service.DeleteUserReservation(context.TODO(), &api.DeleteUserReservationReq{UserID: 0, ReservationID: 9}, resp); err == nil {
		t.Errorf("DeleteUserReservation() error = %v, wantErr %v", err, true)
	}
}
