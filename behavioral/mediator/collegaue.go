package mediator

import "fmt"

type IColleague interface {
	Send(string2 string)
	Notify(string2 string)
}
type Colleague struct {
	mediator IMediator
}

type ConcreteColleagueA struct {
	Colleague
}

func(c *ConcreteColleagueA)Notify(message string){
	if c == nil {
		return
	}
	fmt.Println("Concrete get message: ", message)
}
func(c *ConcreteColleagueA)Send(message string){
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

func NewConcreteColleageA(mediator IMediator)IColleague{
	return &ConcreteColleagueA{Colleague{mediator: mediator}}
}

type ConcreteColleagueB struct {
	Colleague
}
func NewConcreteColleagueB(mediator IMediator)IColleague{
	return &ConcreteColleagueB{Colleague{mediator: mediator}}
}
func(c *ConcreteColleagueB)Notify(string2 string){
	if c == nil{
		return}
	fmt.Println("concere colleagus B message ", string2)
}
func(c *ConcreteColleagueB)Send(string2 string){
	c.mediator.Send(string2, c)
}