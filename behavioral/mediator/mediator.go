package mediator

type IMediator interface {
	Send(string,IColleague)
}


type Mediator struct {
}

type ConcreteMedia struct {
	Mediator
	colleagues []IColleague
}

func(m *ConcreteMedia)AddColleague(c IColleague){
	if m == nil{return }
	m.colleagues = append(m.colleagues, c)
}

func(m *ConcreteMedia)Send(message string, c IColleague){
	if m == nil {
		return
	}
	for _,val := range m.colleagues{
		val.Notify(message)
	}
}
func NewConcreteMediator()*ConcreteMedia{
	return &ConcreteMedia{}
}

