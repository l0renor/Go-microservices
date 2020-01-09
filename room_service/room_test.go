package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"strconv"
	"testing"
)

var roomservice api.Room_Service
var id1 int32

func TestInitRoom(t *testing.T) {
	service := micro.NewService(
		micro.Name("Room_Service"),
		micro.Version("latest"),
	)

	service.Init()

	if err := api.RegisterRoom_ServiceHandler(service.Server(), &Service{}); err != nil {
		t.Error("Error registering room service.")
	}

	if err := service.Run(); err != nil {
		t.Error("Error running room service")
	}

	roomservice = api.NewRoom_Service("Room_Service", service.Client())
}

func TestCreateRoom(t *testing.T) {
	rsp, err := roomservice.CreateRoom(context.TODO(), &api.CreateRoomReq{
		Name:      "Raum1",
		NrOfSeats: 5,
	})
	if err != nil {
		t.Error(err)
	}
	id1 = rsp.RoomID

	getRsp, err := roomservice.GetRoom(context.TODO(), &api.GetRoomReq{RoomID: id1})

	if err != nil {
		t.Error(err)
	}
	if getRsp.Room.RoomID != id1 || getRsp.Room.NrOfSeats != 5 || getRsp.Room.Name != "Raum1" {
		t.Error("Wrong values of room")
	}
}

func TestDeleteRoom(t *testing.T) {
	_, err := roomservice.DeleteRoom(context.TODO(), &api.DeleteRoomReq{RoomID: id1})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateRooms(t *testing.T) {
	ids := make([]int32, 10)
	for i := 0; i < 10; i++ {
		rsp, err := roomservice.CreateRoom(context.TODO(), &api.CreateRoomReq{
			Name:      "Raum" + strconv.Itoa(i),
			NrOfSeats: int32(i),
		})
		if err != nil {
			t.Error(err)
		}
		ids[i] = rsp.RoomID
	}

	rsp, err := roomservice.GetRooms(context.TODO(), &api.GetRoomsReq{})
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		if rsp.Rooms[i].Name != "Raum"+strconv.Itoa(i) || rsp.Rooms[i].NrOfSeats != int32(i) {
			t.Error("Wrong information in rooms")
		}

	}

}
