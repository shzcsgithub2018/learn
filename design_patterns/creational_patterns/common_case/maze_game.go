package common_case

type MazeGame struct {
}

func (m *MazeGame) CreateMaze() {
	aMaze := new(Maze)
	r1 := NewRoom(1)
	r2 := NewRoom(2)
	theDoor := NewDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, new(Wall))
	r1.SetSide(South, theDoor)
	r1.SetSide(East, new(Wall))
	r1.SetSide(West, new(Wall))

	r2.SetSide(North, new(Wall))
	r2.SetSide(South, new(Wall))
	r2.SetSide(East, new(Wall))
	r2.SetSide(West, theDoor)
}
