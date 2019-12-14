package tests

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/room_service"
	"testing"
)

func TestInitRoomService(t *testing.T) {
	service := micro.NewService(
		micro.Name("Room_Service"),
		micro.Version("latest"),
	)

	service.Init()
	api.RegisterRoom_ServiceHandler(service.Server(), &room_service.Service{})

	if err := service.Run(); err != nil {
		t.Error("Error run room Service")
	}

}
