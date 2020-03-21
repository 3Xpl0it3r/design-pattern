package observer

import "fmt"

type update func(int)
type Subject interface {
	Notify()
	State()int
	SetState(int2 int)
	AddCallFunc(*update)
	RemoveCallFunc(*update)
}

type SubjectA struct {
	state int
	call []*update
}
func(s *SubjectA)Notify(){
	for _,c := range s.call{
		(*c)(s.state)
	}
}
func(s *SubjectA)State()int{
	return s.state
}
func(s *SubjectA)SetState(i int){
	s.state = i
}
func(s *SubjectA)AddCallFunc(f *update){
	s.call = append(s.call, f)
}
func(s *SubjectA)RemoveCallFunc(f *update){
	for i,c := range s.call{
		if c == f {
			s.call = append(s.call[:i], s.call[i+1:]...)
		}
	}
}
func NewSubjectA(s int)*SubjectA{
	return &SubjectA{s,[]*update{}}
}


//
type Observer interface {
	Update(int)
}

type ObserverA struct {
	s Subject
	state int
}
func(o *ObserverA)Update(s int){
	fmt.Println("ObserverA")
	fmt.Println(s)
	fmt.Println(o)
}
func NewObserverA (sa Subject, s int)*ObserverA{
	return &ObserverA{
		s:     sa,
		state: s,
	}
}

type ObserverB struct {
	s Subject
	state int
}

func(o *ObserverB)Update(s int){
	fmt.Println("ObserverB")
	fmt.Println(s)
	fmt.Println(o)
	
}

func NewObserverB(sa Subject,s int)*ObserverB{
	return &ObserverB{
		s:     sa,
		state: s,
	}
}