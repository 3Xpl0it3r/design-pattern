package interpreter

import "fmt"

type Context struct {
	text string
}

type AbstratrExpression interface {
	Interpret(ctx *Context)
}

//
type TerminalExpression struct {
}
func(t *TerminalExpression)Interpret(ctx *Context){
	if t == nil {
		return
	}
	ctx.text = ctx.text[:len(ctx.text)-1]
	fmt.Println(ctx)
}

type NonterminalExpression struct {
}

func(n *NonterminalExpression)Interpret(ctx *Context){
	if n == nil {return }
	ctx.text = ctx.text[:len(ctx.text)-1]
	fmt.Println(ctx)
}