package visitor

import "fmt"

type Visitor interface {
	VisitConcreteElementA(a ConcreteElementA)
	VisitConcreteElementB(b ConcreteElementB)
}

type ConcreteVisitorA struct {
	name string
}
func(v *ConcreteVisitorA)VisitConcreteElementA(a ConcreteElementA){
	fmt.Println(v.name, "----",a.name)
	a.OperatorA()
}
func(v *ConcreteVisitorA)VisitConcreteElementB(b ConcreteElementB){
	fmt.Println(v.name, "----",b.name)
	b.OPeratorB()
}



type ConcreteVisitorB struct {
	name string
}

func(v *ConcreteVisitorB)VisitConcreteElementA(a ConcreteElementA){
	fmt.Println(v.name, "----",a.name)
	a.OperatorA()
}
func(v *ConcreteVisitorB)VisitConcreteElementB(b ConcreteElementB){
	fmt.Println(v.name, "----",b.name)
	b.OPeratorB()
}