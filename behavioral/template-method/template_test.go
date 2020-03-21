package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	b := NewBingA()
	b.g = b
	b.getSomeFood()

	g := NewGuo()
	g.g = g
	g.getSomeFood()
}
