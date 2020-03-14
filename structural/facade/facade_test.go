package facade

import "testing"

func TestFacade(t *testing.T) {
	f := NewFacede()
	f.ExecuteA()
	f.ExecuteB()
	f.ExecuteC()
}
