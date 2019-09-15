package main

import (
	"fmt"
	controllers "myproject/lib/controllers"
	structs "myproject/lib/structs"
)

func main(){
	controllers.SomeFunction();
	somev :=  structs.Exp{A : 5}
	fmt.Print( somev)
}