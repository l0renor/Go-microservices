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
	rooms     map[int32]Room
	nextID    func() int32
	screening api.Screening_Service
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
	_, ok := m.rooms[req.GetId()]
	if !ok {
		return errors.NotFound("room_not_found", "room(ID: %v not found", req.Id)
	}
	_, err := m.screening.DeleteRoom(context.TODO(), &api.DeleteRoomReq{RoomID: req.GetId()})
	if err != nil {
		return err
	}
	delete(m.rooms, req.GetId())
	return nil
}

func (m *Service) GetRoom(ctx context.Context, req *api.GetRoomMsg, rsp *api.GetRoomResponseMsg) error {
	room, ok := m.rooms[req.GetId()]
	if ok {
		rsp.Room = &api.RoomData{
			Name:      room.name,
			Id:        req.GetId(),
			NrOfSeats: room.nrOfSeats,
		}
	} else {
		return errors.NotFound("room_not_found", "room(ID: %v not found", req.Id)
	}
	return nil
}

func (m *Service) GetRooms(ctx context.Context, req *api.GetRoomsMsg, rsp *api.GetRoomsResponseMsg) error {
	var rooms []*api.RoomData
	for k, v := range m.rooms {
		rooms = append(rooms, &api.RoomData{
			Name:      v.name,
			Id:        k,
			NrOfSeats: v.nrOfSeats,
		})
	}
	rsp.Rooms = rooms
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("room"),
		micro.Version("latest"),
	)

	screening := micro.NewService()
	screening.Init()

	service.Init()

	if err := api.RegisterRoom_ServiceHandler(service.Server(), &Service{
		rooms:     make(map[int32]Room),
		nextID:    helpers.IDGenerator(),
		screening: api.NewScreening_Service("screening", screening.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
