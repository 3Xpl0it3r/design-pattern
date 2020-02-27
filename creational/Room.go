package creational

import "fmt"

type Room interface {
	MapSite
	SetSide(int, MapSite)
}
type StandardRoom struct {
	id int
}
func NewStandardRoom(id int)*StandardRoom{
	return &StandardRoom{id:id}
}
func(r *StandardRoom)Type()string{
	return "StandardRoom"
}
func(r *StandardRoom)Id()int{
	return r.id
}
func(r *StandardRoom)Enter(){
	fmt.Printf("Room->%s:%d Enter\n", r.Type(), r.Id())
}
func(r *StandardRoom)SetSide(direction int, mapsite MapSite){
	fmt.Printf("Room->%s:%d Set %s:%d\n", r.Type(), r.Id(), mapsite.Type(), mapsite.Id())
}