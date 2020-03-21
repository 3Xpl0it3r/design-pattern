package state

import "fmt"

type State interface {
	writePrograme(w *Work)
}


//
type MoringState struct {
}
func(m *MoringState)writePrograme(w *Work){
	if w.Hour() < 12 {
		fmt.Println("现在是上午十分", w.hour)
	} else {
		w.SetState(new(NoonState))
		w.writePrograme()
	}
}


// 中午
type NoonState struct {
}
func(n *NoonState)writePrograme(w *Work){
	if w.Hour() < 13 {
		fmt.Println("现在还是中午", w.Hour())
	} else {
		w.SetState(new(AfternoonState))
		w.writePrograme()
	}
}


type AfternoonState struct {
}
func(a *AfternoonState)writePrograme(work *Work){
	if work.Hour() < 17 {
		fmt.Println("现在是下午十分", work.Hour())
	} else {
		work.SetState(new(EveningState))
	}
}

type EveningState struct {
}
func(e *EveningState)writePrograme(work *Work){
	if work.Hour() < 21 {
		fmt.Println("现在是晚上时间", work.Hour())
	} else {
		fmt.Println("开始睡觉", work.Hour())
	}
}