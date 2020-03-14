package flyweight

// flyweight factory

type FlyweightFactory struct {
	flyweights map[string]IFlyweight
}

func(f *FlyweightFactory)Flyweight(name string)IFlyweight{
	if f == nil {return nil}
	if name == "u"{
		return &UnsharedConcreteFlyweiht{name: name}
	} else if _,ok := f.flyweights[name];!ok {
		f.flyweights[name] = &ConcreteFlyweight{name: name}
	}
	return f.flyweights[name]
}

func NewFlyWeightFactory()*FlyweightFactory{
	ff := FlyweightFactory{flyweights: make(map[string]IFlyweight)}
	ff.flyweights["a"] = &ConcreteFlyweight{"a"}
	ff.flyweights["b"] = &ConcreteFlyweight{"b"}
	return &ff
}