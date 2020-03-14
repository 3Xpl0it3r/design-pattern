package facade

import "fmt"

type Facade struct {
	Sa SystemA
	Sb SystemB
	Sc SystemC
}

func NewFacede()*Facade{
	return &Facade{
		Sa: SystemA{},
		Sb: SystemB{},
		Sc: SystemC{},
	}
}

func(f *Facade)ExecuteA(){
	f.Sa.Execute()
	f.Sb.Execute()
}

func(f *Facade)ExecuteB(){
	f.Sc.Execute()
	f.Sb.Execute()
}
func(f *Facade)ExecuteC(){
	f.Sa.Execute()
	f.Sc.Execute()
}

type SystemA struct {
}

func(s *SystemA)Execute(){
	fmt.Println("system execute")
}


type SystemB struct {
}
func(s *SystemB)Execute(){
	fmt.Println("system execute")
}

type SystemC struct {
}
func(s *SystemC)Execute(){
	fmt.Println("system execute")
}