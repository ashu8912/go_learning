package main

import (
	"fmt"

	sample "github.com/ashu8912/go_learning/pkg/deferbehaviour"
)
func evaluateSomething(arg string) string{
   fmt.Println("Evaluating something")
	return arg
}
func main(){
 for i:=0;i<5;i++{
	 defer fmt.Println(i)
 }

 defer sample.ShowDeferBehaviour(evaluateSomething(("ashu ghildiyal"))
}