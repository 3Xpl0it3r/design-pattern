package adapter_object

import "fmt"

type Target interface {
	Request()
}

type SpecialAdaptee struct {
}
func(s *SpecialAdaptee)SpeicalRequest(){
	fmt.Println("this is special request")
}
func NewSpecialRequest()*SpecialAdaptee{
	return &SpecialAdaptee{}
}

type Adapter struct {
	Target
	SpecialAdaptee
}

func (a *Adapter)Request(){
	a.SpeicalRequest()
}

func NewAdapter()Target{
	return &Adapter{
		SpecialAdaptee: SpecialAdaptee{},
	}
}