package composite_pattern

import "fmt"

// CompositeEquipment 为包含其他设备的基类
type CompositeEquipment struct {
	BaseEquipment
	List []Equipment
}
func NewCompositeEquipment(name string ,price float32)Equipment{
	return &CompositeEquipment{
		BaseEquipment: BaseEquipment{name: name, price: price},
		List:          []Equipment{},
	}
}
func(c *CompositeEquipment)Display(){
	for _, v := range c.List{
		v.Display()
	}
}
func(c *CompositeEquipment)Name()string{ return c.name}
func(c *CompositeEquipment)Price()float32{return c.price}
func(c *CompositeEquipment)Add(equipment Equipment){ c.List = append(c.List, equipment)}
func(c *CompositeEquipment)Remove(equipment Equipment){
	fmt.Printf("%s cost %f\n", c.name, c.price)
	for i, v := range c.List{
		if v == equipment{
			c.List = append(c.List[:i], c.List[i+1:]...)
		}
	}
}

// FloppyDisk 为软盘设备
type FloppyDisk struct {
	BaseEquipment
}
func NewFloppyDisk(name string, price float32)Equipment{
	return &FloppyDisk{BaseEquipment{
		name:  name,
		price: price,
	}}
}
func(f *FloppyDisk)Display(){fmt.Printf("%s cost %f\n", f.name, f.price)}
func(f *FloppyDisk)Name()string{return f.name}
func(f *FloppyDisk)Price()float32{return f.price}
func(f *FloppyDisk)Add(equipment Equipment){}
func(f *FloppyDisk)Remove(equipment Equipment){}


// Chassis为计算机底盘
type Chassis struct {
	BaseEquipment
	List []Equipment
}
func NewChassis(name string, price float32)Equipment{
	return &Chassis{
		BaseEquipment: BaseEquipment{name: name, price: price},
		List:          []Equipment{},
	}
}
func(c *Chassis)Display(){
	fmt.Printf("%s cost %f\n", c.name, c.price)
	for _, v := range c.List{
		v.Display()
	}
}
func(c *Chassis)Name()string{return c.name}
func(c *Chassis)Price()float32{return c.price}
func(c *Chassis)Add(equipment Equipment){ c.List = append(c.List, equipment)}
func(c *Chassis)Remove(equipment Equipment){
	for i, v := range c.List{
		if v == equipment{
			c.List	= append(c.List[:i], c.List[i+1:]...)
		}
	}
}