package common_case

type Maze struct {
	m map[int32]*Room
}

func (m *Maze) AddRoom(room *Room) {
	m.m[room.roomNo] = room
}

func (m *Maze) RoomNo(roomNo int32) *Room {
	return m.m[roomNo]
}
