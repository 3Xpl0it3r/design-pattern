package interpreter

import "testing"

func TestInterpret(t *testing.T){
	ctx := Context{text: "dsadsadasdaas"}
	list := []AbstratrExpression{}

	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonterminalExpression))

	for _,v := range list{
		v.Interpret(&ctx)
	}
}