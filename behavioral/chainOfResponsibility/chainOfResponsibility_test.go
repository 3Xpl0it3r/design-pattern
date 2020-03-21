package chainOfResponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T){
	common := NewCommonHandle(NoHelpTopic)
	application := NewApplicationHelpHandler(ApplicationTopic)
	dialog := NewDialogHelpHandler(DialogTopic)
	widget := NewWidgetHelpHandler(WidgetTopic)
	button := NewButtonHelpHandler(ButtonTopic)

	common.SetHandler(application)
	application.SetHandler(dialog)
	dialog.SetHandler(widget)
	widget.SetHandler(button)

	common.HandleHelp(ButtonTopic)

}
