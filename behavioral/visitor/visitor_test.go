package visitor

import "testing"

func TestVisitor(t *testing.T) {
	object := ObjectStructure{list: []Element{}}
	object.Attach(&ConcreteElementA{"aaaa"})
	object.Attach(&ConcreteElementB{"bbbb"})

	cva := &ConcreteVisitorA{"VA"}
	cvb := &ConcreteVisitorB{"VB"}


	object.Accept(cva)
	object.Accept(cvb)
}
