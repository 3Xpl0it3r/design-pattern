package bridge

import "testing"

func TestBridge(t *testing.T) {
	var window Window
	window = NewIconWindow(NewLinuxWindowImp())
	window.Draw()
	window = NewIconWindow(NewMacWindowImp())
	window.Draw()
}