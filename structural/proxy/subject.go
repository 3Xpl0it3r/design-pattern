package proxy

// Graphic is reference of subject
type Graphic interface {
	Draw()
	HandleMouse()
	Load()
	Save()
}
