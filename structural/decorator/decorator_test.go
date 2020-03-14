package decorator

import "testing"

func TestDecorator(t *testing.T) {
	person := &Persion{Name: "xiaowang"}
	person.Show()

	ts := new(TShirts)
	ts.SetDecorator(person)

	bt := new(BigTrouser)
	bt.SetDecorator(person)
	bt.Show()
}
