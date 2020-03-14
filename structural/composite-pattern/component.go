package composite_pattern


// component
// 在整体层次为所有的设备定义一个接口
type Equipment interface {
	Name()string
	Price()float32
	Add(equipment Equipment)
	Remove(equipment Equipment)
	Display()
}

type BaseEquipment struct {
	name string
	price float32
}