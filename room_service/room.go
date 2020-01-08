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

func (m *Service) CreateRoom(ctx context.Context, req *api.CreateRoomReq, resp *api.CreateRoomResp) error {
	id := m.nextID()
	m.rooms[id] = Room{
		nrOfSeats: req.NrOfSeats,
		name:      req.Name,
	}
	resp.RoomID = id
	return nil
}

func (m *Service) DeleteRoom(ctx context.Context, req *api.DeleteRoomReq, resp *api.DeleteRoomResp) error {
	_, ok := m.rooms[req.GetRoomID()]
	if !ok {
		return errors.NotFound("ERR-NO-ROOM", "Room (ID: %d) not found!", req.GetRoomID())
	}
	_, err := m.screening.DeleteScreeningsWithRoom(context.TODO(), &api.DeleteScreeningsWithRoomReq{RoomID: req.GetRoomID()})
	if err != nil {
		return err
	}
	delete(m.rooms, req.GetRoomID())
	return nil
}

func (m *Service) GetRoom(ctx context.Context, req *api.GetRoomReq, resp *api.GetRoomResp) error {
	room, ok := m.rooms[req.GetRoomID()]
	if !ok {
		return errors.NotFound("ERR-NO-ROOM", "Room (ID: %d) not found!", req.GetRoomID())
	}
	resp.Room = &api.Room{
		Name:      room.name,
		RoomID:    req.GetRoomID(),
		NrOfSeats: room.nrOfSeats,
	}
	return nil
}

func (m *Service) GetRooms(ctx context.Context, req *api.GetRoomsReq, resp *api.GetRoomsResp) error {
	var rooms []*api.Room
	for id, room := range m.rooms {
		rooms = append(rooms, &api.Room{
			Name:      room.name,
			RoomID:    id,
			NrOfSeats: room.nrOfSeats,
		})
	}
	resp.Rooms = rooms
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
		nextID:    helpers.IDGenerator(0),
		screening: api.NewScreening_Service("screening", screening.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
