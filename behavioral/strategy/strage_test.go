package strategy

import (
	"fmt"
	"testing"
)

func TestStrage(t *testing.T){
	context,err := NewCashContext("正常收费")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(1000)
		fmt.Println(result)
	}

	context,err = NewCashContext("满300返还100")
	if err!= nil {
		t.Error(err)
	} else {
		result := context.acceptCash(350)
		fmt.Println(result)
	}

	context,err = NewCashContext("8折优惠")
	if err!= nil {
		t.Error(err)
	} else {
		result := context.acceptCash(300)
		fmt.Println(result)
	}

}
