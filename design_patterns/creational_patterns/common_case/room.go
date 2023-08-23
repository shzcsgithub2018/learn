package common_case

type Room struct {
	mapSite [4]*MapSite
	roomNo  int32
}

func NewRoom(roomNo int32) *Room {
	return &Room{
		roomNo: roomNo,
	}
}

func (r *Room) Enter() {

}

func (r *Room) GetSide(direction Direction) {

}

func (r *Room) SetSide(direction Direction, site MapSite) {

}
