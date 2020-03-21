package command

import (
	"fmt"
)

// Command is the interface of all concrete command
type Command interface {
	Execute()
}


// command A
type ConcreteCommandA struct {
	receiver *ReceiverA
}
func(c *ConcreteCommandA)Execute(){
	c.receiver.Execute()
}
func (c *ConcreteCommandA)SetReceiver(r *ReceiverA)  {
	if c == nil {return}
	c.receiver = r
}
func NewConcreteCommandA()*ConcreteCommandA{
	return &ConcreteCommandA{}
}

type ReceiverA struct {
}
func(r *ReceiverA)Execute(){
	fmt.Println("针对ConcreteCommandA，如何处理该命令 ")
}
func NewReceiverA()*ReceiverA{
	return &ReceiverA{}
}


//

type ConcreteCommandB struct {
	receiver *ReceiverB
}

func (c *ConcreteCommandB) SetReceiver(r *ReceiverB) {
	if c == nil {
		return
	}
	c.receiver = r
}

// 具体命令的执行体
func (c *ConcreteCommandB) Execute() {
	if c == nil {
		return
	}
	c.receiver.Execute()
}

func NewConcreteCommandB() *ConcreteCommandB {
	return &ConcreteCommandB{}
}

// 针对ConcreteCommandB，如何处理该命令
type ReceiverB struct {
}

func (r *ReceiverB) Execute() {
	if r == nil {
		return
	}
	fmt.Println("针对ConcreteCommandB，如何处理该命令")
}

func NewReceiverB() *ReceiverB {
	return &ReceiverB{}
}
