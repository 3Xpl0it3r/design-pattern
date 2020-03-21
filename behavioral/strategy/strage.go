package strategy

import "errors"

type CashSuper interface {
	acceptCash(memory float32)float32
}

type CashNormal struct {
}
func(c *CashNormal)acceptCash(money float32)float32{
	return money
}

type CashRebate struct {
	memeyRebate float32
}

func(this *CashRebate)acceptCash(money float32)float32{
	return this.memeyRebate * money
}

type CashReturn struct {
	memeyCondition float32
	memReturn float32
}

func(this* CashReturn)acceptCash(money float32)float32{
	if money >= this.memeyCondition{
		return money - float32(int(money/this.memeyCondition*this.memReturn))
	}else {
		return money
	}
}

//
type Context struct {
	CashSuper
}

func NewCashContext(str string)(cash *Context,err error){
	cash = new(Context)
	switch str {
	case "正常收费":
		cash.CashSuper = &CashNormal{}
	case "满300返还100":
		cash.CashSuper = &CashReturn{300,100}
	case "8折优惠":
		cash.CashSuper = &CashRebate{0.98}
	default:
		err = errors.New("没有符合的")
	}
	return cash, nil
}