package creational

type MazeFactoryMethod interface {
	MakeMaze(int)Maze
	MakeWall(int)Maze
	MakeRoom(int)Room
	MakeDoor(int, Room, Room)Door

	CreateMaze()Maze
}

type StandardMazeFactoryMethod struct {
}
func (factorymethod *StandardMazeFactoryMethod)MakeMaze(id int)Maze	{
	return NewStandardMaze(id)
}
func (factorymethod *StandardMazeFactoryMethod)MakeWall(id int)Wall{
	return NewStandardWall(id)
}
func (factorymethod *StandardMazeFactoryMethod)MakeRoom(id int)Room{
	return NewStandardRoom(id)
}
func (factorymethod *StandardMazeFactoryMethod)MakeDoor(id int, roomFrom Room, roomTo Room)Door{
	return NewStandardDoor(id, roomFrom, roomTo)
}
func (factorymethod *StandardMazeFactoryMethod)CreateMaze()Maze{
	aMaze := factorymethod.MakeMaze(1)
	factorymethod.MakeWall(1)
	//roomFrom := factorymethod.MakeRoom(1)
	//roomTo := factorymethod.MakeRoom(2)
	//door := factorymethod.MakeDoor(1, roomFrom, roomTo)
	//
	//aMaze.AddRoom(roomTo)
	//aMaze.AddRoom(roomFrom)
	//
	//roomTo.SetSide(North, door)
	//roomFrom.SetSide(South, factorymethod.MakeWall(1))
	return aMaze
}





