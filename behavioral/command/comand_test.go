package command

import (
	"testing"
)

func TestCommand(t *testing.T){

	invoke := NewInvoke()
	cmdA := NewConcreteCommandA()
	cmdA.SetReceiver(NewReceiverA())

	cmdB := NewConcreteCommandB()
	cmdB.SetReceiver(NewReceiverB())

	invoke.AddCommand(cmdA)
	invoke.AddCommand(cmdB)

	invoke.ExecuteCommand()


}
