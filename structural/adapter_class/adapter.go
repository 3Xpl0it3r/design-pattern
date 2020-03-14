package adapter_class

import "fmt"

type Target interface {
	Request()
}

type Adaptee interface {
	SpecialRequest()
}

type Adapter interface {
	Target
	Adaptee
}
//
type SpecialAdapteeA struct {
}

func(s *SpecialAdapteeA)Request(){
	s.SpecialRequest()
}
func(s *SpecialAdapteeA)SpecialRequest(){
	fmt.Println("Special request from adaptee a")
}
func NewAdapterForSpecialAdapteeA()Target{
	return &SpecialAdapteeA{}
}

//
type SpecialAdapteeB struct {
}
func (s *SpecialAdapteeB)Request(){
	s.SpecialRequest()
}
func (s *SpecialAdapteeB)SpecialRequest(){
	fmt.Println("Special request from adaptee b")
}

func NewApapterForSpecialAdapteeB()Target{
	return &SpecialAdapteeB{}
}
