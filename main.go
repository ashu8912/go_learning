package main

import (
	"fmt"

	"github.com/ashu8912/go-learning/pkg/show_defer"
)
func evaluateSomething(arg string){
   fmt.Println("Evaluating something")
	return arg
}
func main(){
 for i:=0;i<5;i++{
	 defer fmt.Println(i)
 }

 defer show_defer.ShowDeferBehaviour(evaluateSomething("Ashu"))
}