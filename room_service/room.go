package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"log"
)

type Room struct {
	nrOfSeats int32
	name      string
}

type Service struct {
	rooms  map[int32]Room
	nextID func() int32
}

func (m *Service) CreateRoom(ctx context.Context, req *api.CreateRoomMsg, rsp *api.CreateRoomResponseMsg) error {
	id := m.nextID()
	m.rooms[id] = Room{
		nrOfSeats: req.NrOfSeats,
		name:      req.Name,
	}
	rsp.Id = id
	return nil
}

func (m *Service) DeleteRoom(ctx context.Context, req *api.DeleteRoomMsg, rsp *api.DeleteRoomResponseMsg) error {
	id := req.Id
	delete(m.rooms, id)
	_, ok := m.rooms[id]
	if !ok {
		return errors.NotFound("room_not_found", "room(ID: %v not found", req.Id)
	}
	return nil
}

func (m *Service) GetRoom(ctx context.Context, req *api.GetRoomMsg, rsp *api.GetRoomResponseMsg) error {
	id := req.Id
	res, ok := m.rooms[id]
	if ok {
		rsp.Room = &api.RoomData{
			Name:      res.name,
			Id:        id,
			NrOfSeats: res.nrOfSeats,
		}
	} else {
		return errors.NotFound("room_not_found", "room(ID: %v not found", req.Id)
	}
	return nil
}

func (m *Service) GetRooms(ctx context.Context, req *api.GetRoomsMsg, rsp *api.GetRoomsResponseMsg) error {
	var res []*api.RoomData
	for k, v := range m.rooms {
		res = append(res, &api.RoomData{
			Name:      v.name,
			Id:        k,
			NrOfSeats: v.nrOfSeats,
		})
	}
	rsp.Rooms = res
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("room"),
		micro.Version("latest"),
	)

	service.Init()

	if err := api.RegisterRoom_ServiceHandler(service.Server(), &Service{
		rooms:  make(map[int32]Room),
		nextID: helpers.IDGenerator(),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
