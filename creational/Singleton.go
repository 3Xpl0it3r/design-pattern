package creational

import "sync"

// singleton pattern
// standard singleton pattern

type mazeFactorySingleton struct {
	Factory MazeAbstractFactory
}

var instance *mazeFactorySingleton
var once sync.Once

func GetMazeFactory()*mazeFactorySingleton{
	once.Do(func() {
		instance = &mazeFactorySingleton{}
		instance.Factory = new(MazeStandardAbstractFactory)
	})
	return instance
}

// register of singleton
type registerSingle struct {
	factory map[string]MazeAbstractFactory
}

var registerInstance *registerSingle

func init() {
	once.Do(func() {
		registerInstance = &registerSingle{}
		registerInstance.factory = make(map[string]MazeAbstractFactory, 10)
	})
}

func GetFactoryByName(name string)MazeAbstractFactory{
	return registerInstance.factory[name]
}
func RegisterFactoryToSingleton(name string, facotry MazeAbstractFactory){
	registerInstance.factory[name] = facotry
}