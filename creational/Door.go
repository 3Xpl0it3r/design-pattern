package creational

import "fmt"

type Door interface {
	MapSite
	// for prototype
	Clone()Door
}
// standdoor product 1
type StandardDoor struct {
	roomFrom Room
	roomTo Room
	isOpen bool
	id int
}

func NewStandardDoor(id int, roomFrom Room, roomTo Room,)*StandardDoor{
	return &StandardDoor{
		roomFrom: roomFrom,
		roomTo:   roomTo,
		isOpen:   false,
		id:       id,
	}
}
func NewStandardDoorDeepCopy(door *StandardDoor)*StandardDoor{
	return &StandardDoor{
		roomFrom: door.roomFrom,
		roomTo:   door.roomTo,
		isOpen:   door.isOpen,
		id:       door.id,
	}
}
func(d *StandardDoor)Type()string{
	return "StandardDoor"
}
func(d *StandardDoor)Id()int{
	return d.id
}
func(d *StandardDoor)Enter(){
	fmt.Printf("Door->%s:%d Enter\n", d.Type(), d.Id())
}
func(d *StandardDoor)Clone()Door{
	return NewStandardDoorDeepCopy(d)
}



// enchanted door : product2
type EnchantedDoor struct {
	*StandardDoor
}

func NewEnchantedDoor(id int, roomFrom Room, roomTo Room)*EnchantedDoor{
	return &EnchantedDoor{NewStandardDoor(id, roomFrom, roomTo)}
}
func NewEnchantedDoorDeepCopy(door *EnchantedDoor)*EnchantedDoor{
	return &EnchantedDoor{NewStandardDoor(door.id, door.roomFrom, door.roomTo)}
}
func(door *EnchantedDoor)Type()string{
	return "EnchantedDoor"
}
func (door *EnchantedDoor)Clone()Door{
	return NewEnchantedDoorDeepCopy(door)
}
