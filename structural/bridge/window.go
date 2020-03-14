package bridge
// abstraction

type Window interface {
	Draw()
}

// icon window

type IconWindow struct {
	platform WindowImp
}


func(w *IconWindow)Draw(){
	w.platform.DeviceDraw()
}

func NewIconWindow(imp WindowImp)Window{
	return &IconWindow{platform: imp}
}

type ButtonWindow struct {
	platform WindowImp
}

func(b *ButtonWindow)Draw(){
	b.platform.DeviceDraw()
}

func NewButtonWindow(imp WindowImp)Window{
	return &ButtonWindow{platform: imp}
}