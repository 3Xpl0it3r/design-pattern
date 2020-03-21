package chainOfResponsibility

import "fmt"

type Topic int

const (
	NoHelpTopic = iota
	WidgetTopic
	ButtonTopic
	DialogTopic
	ApplicationTopic

)

type IHelpHandler interface {
	HasHelp(topic Topic)bool
	SetHandler(handler IHelpHandler)
	HandleHelp(topic Topic)error
}

type HelpHandler struct {
	successor IHelpHandler
	topic Topic
}

func NewCommonHandle(topic Topic)*HelpHandler{
	return &HelpHandler{
		successor: nil,
		topic:     topic,
	}
}

func(h *HelpHandler)HasHelp(topic Topic)bool{
	return h.topic == topic
}

func(h *HelpHandler)SetHandler(handler IHelpHandler){
	h.successor = handler
}
func(h *HelpHandler)HandleHelp(topic Topic)error{
	fmt.Println("this is common ", topic, h.topic)
	if h.topic == topic && topic == NoHelpTopic{
		fmt.Println("this is default help")
		return nil
	} else if h.successor!= nil{
		fmt.Println("我处理不了，传递给下一个")
		return h.successor.HandleHelp(topic)
	} else {
		fmt.Println("模特你")
		return nil
	}
}