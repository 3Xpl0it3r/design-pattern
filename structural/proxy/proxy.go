package proxy

import (
	"time"
)

type ImageProxy struct {
	image *Image
}

func(i *ImageProxy)Draw(){
	time.Sleep(2* time.Second)
	i.image.Draw()
}
func(i *ImageProxy)HandleMouse(){
	time.Sleep(1 * time.Second)
	i.image.HandleMouse()
}
func(i *ImageProxy)Load(){
	time.Sleep(1 * time.Second)
	i.image.Load()
}
func(i *ImageProxy)Save(){
	i.image.Save()
}


func NewProxy()Graphic{
	return &ImageProxy{image: &Image{}}
}