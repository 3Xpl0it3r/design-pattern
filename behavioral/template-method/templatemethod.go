package template_method

import "fmt"

type getfood interface {
	first()
	second()
	third()
}

type template struct {
	g getfood
}

func(b *template)getSomeFood(){
	if b == nil{return }
	b.g.first()
	b.g.second()
	b.g.third()
}


type bingA struct {
	template
}

func NewBingA()*bingA{
	return &bingA{}
}
func(b *bingA)first(){
	fmt.Println("打开冰箱")
}
func(b *bingA)second(){
	fmt.Println("取出东西")
}
func(b *bingA)third(){
	fmt.Println("关上冰箱")
}

type Guo struct {
	template
}

func NewGuo()*Guo{
	return &Guo{}
}
func(b *Guo)first(){
	fmt.Println("揭开锅盖")
}
func(b *Guo)second(){
	fmt.Println("盛碗饭")
}
func(b *Guo)third(){
	fmt.Println("关闭锅盖")
}

