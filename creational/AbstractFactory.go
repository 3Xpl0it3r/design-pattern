package creational

type MazeAbstractFactory interface {
	MakeMaze(int)Maze
	MakeRoom(int)Room
	MakeDoor(int, Room, Room)Door
	MakeWall(int)Wall
}
// standard factory
type MazeStandardAbstractFactory struct {
}
func (factory *MazeStandardAbstractFactory)MakeMaze(id int)Maze{
	return NewStandardMaze(id)
}
func(factory *MazeStandardAbstractFactory)MakeRoom(id int)Room{
	return NewStandardRoom(id)
}
func(factory *MazeStandardAbstractFactory)MakeDoor(id int, roomFrom Room, roomTo Room)Door{
	return NewStandardDoor(id, roomFrom, roomTo )
}
func(factory *MazeStandardAbstractFactory)MakeWall(id int)Wall{
	return NewStandardWall(id)
}
// enchanted factory
type MazeEnchantedAbstractFactory struct {
	MazeStandardAbstractFactory
}
func (factory *MazeEnchantedAbstractFactory)MakeWall(id int)Wall{
	return NewEnchantedWall(id)
}
func (factory *MazeEnchantedAbstractFactory)MakeDoor(id int, roomFrom Room, roomTo Room)Door{
	return NewEnchantedDoor(id, roomFrom, roomTo)
}