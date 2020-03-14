package proxy

import "fmt"

type Image struct {
}

func(i *Image)Draw(){
	fmt.Println("Image draw")
}
func (i *Image)HandleMouse()  {
	fmt.Println("Image handlemouse")
}
func(i *Image)Load(){
	fmt.Println("Image Load")
}
func(i *Image)Save(){
	fmt.Println("Image save")
}