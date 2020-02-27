package creational

import "fmt"

type Maze interface {
	MapSite
	AddRoom(Room)
}

type StandardMaze struct {
	id int
}
func NewStandardMaze(id int)*StandardMaze{
	return &StandardMaze{id:id}
}
func(m *StandardMaze)Enter(){
	fmt.Printf("Maze: %s:%d Enter\n", m.Type(), m.Id())
}
func(m *StandardMaze)Type()string{
	return "StandardMaze"
}
func(m *StandardMaze)Id()int{
	return m.id
}
func(m *StandardMaze)AddRoom(room Room){
	fmt.Printf("Maze->%s:%d Add %s:%d\n",m.Type(), m.Id(), room.Type(), room.Id())
}