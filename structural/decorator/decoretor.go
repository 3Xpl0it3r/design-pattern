package decorator

import "fmt"

// component
type AbstractPerson interface {
	Show()
}
// ConcreteComponent
type Persion struct {
	Name string
}
func(p *Persion)Show(){
	fmt.Println("姓名：", p.Name)
}

//Decorator
type Decorator struct {
	AbstractPerson
}

func(d *Decorator)SetDecorator(person AbstractPerson){
	if d == nil{
		return
	}
	d.AbstractPerson = person
}
func(d *Decorator)Show(){
	if d == nil {return}
	if d.AbstractPerson == nil{return}
	d.AbstractPerson.Show()
}

// ConcreteDecorator
type TShirts struct {
	Decorator
}
func(t *TShirts)Show(){
	fmt.Println("Ready ware T恤")
	t.Decorator.Show()
}

type BigTrouser struct {
	Decorator
}
func(b *BigTrouser)Show(){
	fmt.Println("大裤衩")
	b.Decorator.Show()
}