# Golang Project to explain the Imports

In golang we cant import file we always import packages. 
<br>
So lets look at a project and understand how we can create and import packages for use
<br>

---
### **Make sure you are creating project inside your GO workspace under the src directory** 
---
### Directory Structure
Let us suppose we have a main.go file where the main code goes
<br>

`|- /lib`

`|- - /controllers`

`|- - - - controller.go`

`|- - /structs`

`|- - - struct.go`

`|- main.go`


so our structs are defined in the structs.go file inside structs.go folder.
And our controller where we can have some function can use structs 

so first thing is we need to make a package choose any name for the package i have used myproject

### struct.go

```go
package foo

type Exp struct {
	A int;
}
```

Here we created an experiment struct Exp  with  a property A . 
Since the name of the struct is in Capital letters it is exportable similarly the propery A is also in capital letter it is also public.

save it we will be using it later.


### Controller.go

```go
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
```
In our controller first we need to import the struct in order to so you need to provide the full path from root directorory of project to this directory so in this case it is `myproject/lib/structs`. Now once we have imported it we need to give it a alias to make thing recognizable. But you can also do 
<br>

`. "myproject/lib/structs" ` instead of 
<br>
`s "myproject/lib/structs"`

 in that case you would be directly able to create objects instances like 

` experiment := Exp{A:10} `

with this we are ready to use controller in our main file

### main.go

```go
package main

import (
	"fmt"
    //import controller and alias as 'controllers'
	controllers "myproject/lib/controllers"

    //import all structs and alias as 'structs' 
	structs "myproject/lib/structs"
)

func main(){

    //now we can use any public function inside main.
	controllers.SomeFunction();
	
    //similarly any struct
    somev :=  structs.Exp{A : 5}
	
    fmt.Print( somev)
}
```

like this you can use local packages you can as many directories as you want and nested levels the path will be always from root folder of the project.


### BUILD 
to build this in bin 
first set the gobin to current directory's bin 

```sh 
export GOBIN="$GOPATH/src/myproject/bin"
```

then run 
```sh
go install 
```

then execute it to test

```sh
./bin/myproject
```

