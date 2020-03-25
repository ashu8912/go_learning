package main

import (
	"fmt"

	sample "github.com/ashu8912/go_learning/pkg/deferbehaviour"
)

func evaluateSomething(arg string) string {
	fmt.Println("Evaluating something")
	return arg
}
func main() {
	a:=[...]int{1,2,3,4,5}
	sl:=a[:]
	sl=append(sl,6,7,8,9)
	b:=make([]int,4)
	for i:=0;i<len(b);i++{
		b[i]=i*5
	}
	sl=append(sl,b...)
	fmt.Println(sl)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	defer sample.ShowDeferBehaviour(evaluateSomething(("ashu ghildiyal")))
}
