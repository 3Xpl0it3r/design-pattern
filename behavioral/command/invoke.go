package command

type Invoke struct {
	cmdList []Command
}
func NewInvoke()*Invoke{
	return &Invoke{cmdList: []Command{}}
}

func(i *Invoke)AddCommand(c Command){
	if i == nil {
		return
	}
	i.cmdList = append(i.cmdList, c)
}

func(i *Invoke)RemoveCommand(c Command){
	if i == nil{
		return
	}
	for j,v := range i.cmdList{
		if v == c {
			i.cmdList = append(i.cmdList[:j], i.cmdList[j+1:]...)
		}
	}
}

func(i *Invoke)ExecuteCommand(){
	for _,c := range i.cmdList{
		c.Execute()
	}
}