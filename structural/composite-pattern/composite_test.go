package composite_pattern

import "testing"

func TestComposite(t *testing.T) {
	root := NewCompositeEquipment("base", 12.2)

	chassis := NewChassis("地盘",1100.2)
	chassis.Add(NewFloppyDisk("hp disk", 34.9))
	chassis.Add(NewFloppyDisk("dell disk", 125.6))

	root.Add(chassis)

	root.Display()
}

