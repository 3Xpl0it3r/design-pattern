package creational
const (
	North = iota
	South
	East
	West
)

type MapSite interface {
	Enter()
	Type()string
	Id()int
}
