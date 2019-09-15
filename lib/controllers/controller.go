package foo

import (
	"fmt"
	s "myproject/lib/structs"
)

func SomeFunction(  ){
	experiment := s.Exp{A:10}
	experiment.A += 10;	
	fmt.Print(experiment)
	fmt.Print("/n")

}