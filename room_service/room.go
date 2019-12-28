package room_service

import (
	"context"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"

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
	rsp.Success = !ok
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
		rsp.Room = &api.RoomData{}
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

func idGenerator() func() int32 {
	i := 0
	return func() int32 {
		i++
		return int32(i)
	}
}
