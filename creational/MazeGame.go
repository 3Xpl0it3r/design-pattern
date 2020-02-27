package creational

type MazeGame struct {
}

func(mg *MazeGame)CreateMazeFromAbstractFactory(factory MazeAbstractFactory)Maze{
	aMaze := factory.MakeMaze(1)

	roomFrom := factory.MakeRoom(1)
	roomTo := factory.MakeRoom(2)

	door := factory.MakeDoor(1, roomFrom, roomTo)

	aMaze.AddRoom(roomFrom)
	aMaze.AddRoom(roomTo)

	roomFrom.SetSide(North, factory.MakeWall(1))
	roomTo.SetSide(South, door)
	return aMaze
}
func (mg *MazeGame)CreateMazeFromBuilder(builder MazeBuilder)Maze{
	builder.BuildMaze(1)

	return builder.GetMaze()
}
