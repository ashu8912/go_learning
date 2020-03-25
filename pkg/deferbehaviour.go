package deferbehaviour

import "fmt"

//this funtion is a helper function to show defer behaviour
func ShowDeferBehaviour(arg string) string{
 fmt.Println(" Hello ",arg," you are in a defered function")
 return arg
}