package state

type Work struct {
	hour int
	state State
}

func(w *Work)Hour()int{
	return w.hour
}

func(w *Work)State()State{
	return w.state
}

func(w *Work)SetHour(h int){
	w.hour = h
}
func(w *Work)SetState(s State){
	w.state = s
}

func(w *Work)writePrograme(){
	if w == nil {return }
	w.state.writePrograme(w)
}

func NewWork()*Work{
	return &Work{
		hour:  0,
		state: new(MoringState),
	}
}