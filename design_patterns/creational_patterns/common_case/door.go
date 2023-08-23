package common_case

type Door struct {
	room1, room2 *Room
	isOpen       bool
}

func NewDoor(room1, room2 *Room) *Door {
	return &Door{
		room1: room1,
		room2: room2,
	}
}

func (d *Door) Enter() {

}

func (d *Door) OtherSideFrom(room *Room) *Room {
	return nil
}