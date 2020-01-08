package main

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"testing"
)

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

}
