package "show_defer"

import "fmt"

func ShowDeferBehaviour(arg string){
 fmt.Println(" Hello ",arg," you are in a defered function")
 return arg
}