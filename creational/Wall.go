package creational

import "fmt"

type Wall interface {
	MapSite
}
// stand product : wall
type StandardWall struct {
	id int
}
func NewStandardWall(id int)*StandardWall{
	return &StandardWall{id:id}
}
func(w *StandardWall)Type()string{
	return "StandardWall"
}
func(w *StandardWall)Id()int{
	return w.id
}
func(w *StandardWall)Enter(){
	fmt.Printf("Wall->%s:%d Enter\n", w.Type(), w.Id())
}

// enchanrted wall
type EnchantedWall struct {
	*StandardWall
}
func NewEnchantedWall(id int)*EnchantedWall{
	fmt.Printf("this is nw enchaneted wall \n")
	return &EnchantedWall{NewStandardWall(id)}
}
func (wall *EnchantedWall)Type()string{
	return "EnchantedWall"
}
func(wall *EnchantedWall)Enter(){
	fmt.Println("this is enchanted wall enter")
}