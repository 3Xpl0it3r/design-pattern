package visitor
type ObjectStructure struct {
	list []Element
}

func(o *ObjectStructure)Attach(element Element){
	o.list = append(o.list, element)
}

func(o *ObjectStructure)Accept(visitor Visitor){
	for _,v := range o.list{
		v.Accept(visitor)
	}
}
