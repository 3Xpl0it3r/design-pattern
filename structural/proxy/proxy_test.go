package proxy

import "testing"

func TestProxy(t *testing.T) {
	ip := NewProxy()
	ip.Draw()
	ip.Save()
	ip.HandleMouse()
}
