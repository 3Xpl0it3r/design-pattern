package flyweight

import "fmt"


// flyweight interface
type IFlyweight interface {
	Operation(int)
}
// concrete flyweight
type ConcreteFlyweight struct {
	name string
}
func(c *ConcreteFlyweight)Operation(outState int){
	if c == nil{return}
	fmt.Println("共享对象相应外部状态", outState)
}

//unshared concrete flyweight
type UnsharedConcreteFlyweiht struct {
	name string
}
func(uc *UnsharedConcreteFlyweiht)Operation(outState int){
	if uc == nil{return}
	fmt.Println("不共享对象相应外部状态", outState)
}

