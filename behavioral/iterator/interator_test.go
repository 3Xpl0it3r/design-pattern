package iterator

import (
	"fmt"
	"testing"
)

func TestItrator(t *testing.T){
	bg := BookGroup{}
	bg.Add(Book{name: "sdasda"})
	bg.Add(Book{"hehe"})

	it := bg.CreateInterator()

	for b := it.First(); b!= nil;b = it.Next(){
		fmt.Println(b)
	}
}
