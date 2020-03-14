package bridge

import "fmt"

// implementor
type WindowImp interface {
	DeviceDraw()
}

// concreteImplementor
type LinuxWindowImp struct {
}

func(w *LinuxWindowImp)DeviceDraw(){
	fmt.Println("Draw a window on linux")
}

func NewLinuxWindowImp()WindowImp{
	return &LinuxWindowImp{}
}

type MacWindowImp struct {
}

func(w *MacWindowImp)DeviceDraw(){
	fmt.Println("Draw a window on mac")
}

func NewMacWindowImp()WindowImp{
	return &MacWindowImp{}
}