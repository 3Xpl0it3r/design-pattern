package visitor

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}


type ConcreteElementA struct {
	name string
}
func(e *ConcreteElementA)Accept(visitor Visitor){
	visitor.VisitConcreteElementA(*e)
}
func(e *ConcreteElementA)OperatorA(){
	fmt.Println("OperatorA")
}



type ConcreteElementB struct {
	name string
}

func(e *ConcreteElementB)Accept(visitor Visitor){
	visitor.VisitConcreteElementB(*e)
}
func(e *ConcreteElementB)OPeratorB(){
	fmt.Println("OperatorB")
}

