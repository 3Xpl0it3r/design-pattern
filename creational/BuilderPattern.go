package creational

type MazeBuilder interface {
	BuildMaze(int)
	BuildDoor(int,Room, Room)
	BuildRoom(int)
	GetMaze()Maze
}

type MazeStandardBuilder struct {
	maze Maze
}
func (builder *MazeStandardBuilder)BuildMaze(id int){
	builder.maze = NewStandardMaze(id)
}
func(builder *MazeStandardBuilder)BuildRoom(id int){
	room := NewStandardRoom(id)
	room.SetSide(North, NewStandardWall(1))
	room.SetSide(South, NewStandardWall(2))
	builder.maze.AddRoom(room)
}
func(builder *MazeStandardBuilder)BuildDoor(id int,roomFrom Room, roomTo Room){
	door := NewStandardDoor(id, roomFrom, roomTo)
	roomFrom.SetSide(North, door)
}
func (builder *MazeStandardBuilder)GetMaze()Maze{
	return builder.maze
}