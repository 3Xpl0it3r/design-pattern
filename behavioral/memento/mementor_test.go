package memento

import "testing"

func TestMementor(t *testing.T){
	gr := GameRole{}
	gr.GetInitState()
	gr.StateDisplay()

	careTaker := RoleStateGareTaker{}
	careTaker.memento = gr.SaveState()
	gr.Fight()
	gr.StateDisplay()
	gr.RecoveryState(careTaker.memento)
	gr.StateDisplay()
}
