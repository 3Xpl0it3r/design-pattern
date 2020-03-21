package chainOfResponsibility

import (
	"fmt"
)

type WidgetHelpHandler struct {
	HelpHandler
}
func NewWidgetHelpHandler(topic Topic)*WidgetHelpHandler{
	return &WidgetHelpHandler{HelpHandler{
		successor: nil,
		topic:     topic,
	}}
}

func(wh *WidgetHelpHandler)HandleHelp(topic Topic)error{
	if wh.HasHelp(topic){
		fmt.Println("Widget help information")
		return nil
	} else if wh.successor != nil{
		return wh.successor.HandleHelp(topic)
	} else {
		return nil
	}
}




type ButtonHelpHandler struct {
	HelpHandler
}
func NewButtonHelpHandler(topic Topic)*ButtonHelpHandler{
	return &ButtonHelpHandler{HelpHandler{
		successor: nil,
		topic:     topic,
	}}
}

func(bh *ButtonHelpHandler)HandleHelp(topic Topic)error {
	if bh.HasHelp(topic){
		fmt.Println("This is button help information")
		return nil
	} else if bh.successor != nil{
		return bh.successor.HandleHelp(topic)
	} else {
		return nil
	}
}



type DialogHelpHandler struct {
	HelpHandler
}
func NewDialogHelpHandler(topic Topic)*DialogHelpHandler{
	return &DialogHelpHandler{HelpHandler{
		successor: nil,
		topic:     topic,
	}}
}

func(dh *DialogHelpHandler)HandleHelp(topic Topic)error{
	if dh.HasHelp(topic){
		fmt.Println("Dialog help information")
		return nil
	} else if dh.successor!= nil{
		return dh.successor.HandleHelp(topic)
	} else {
		return nil
	}
}


type ApplicationHelpHandler struct {
	HelpHandler
}
func NewApplicationHelpHandler(topic Topic)*ApplicationHelpHandler{
	return &ApplicationHelpHandler{HelpHandler{
		successor: nil,
		topic:     topic,
	}}
}

func(ah *ApplicationHelpHandler)HandleHelp(topic Topic)error{
	if ah.HasHelp(topic){
		fmt.Println("Application help information")
		return nil
	} else if ah.successor != nil {
		return ah.successor.HandleHelp(topic)
	} else {
		return nil
	}
}