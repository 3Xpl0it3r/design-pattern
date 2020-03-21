package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T){
	w := NewWork()
	fmt.Println("=====================================")
	w.writePrograme()
	fmt.Println("==========================")
	w.SetHour(12)
	w.writePrograme()
	fmt.Println("++++++++++==================")
	w.SetHour(14)
	w.writePrograme()
	fmt.Println("===========================")
	w.SetHour(21)
	w.writePrograme()
	fmt.Println("================================")
}