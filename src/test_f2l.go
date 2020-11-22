package main

import (
	"fmt"
	"./rubik"
	"./oll"
)

func main() {
	fmt.Println("Test f2l")
	c := rubik.NewSolvedCube()
	//c = c.Rc().D().L().Dc().R().D().Lc().Dc()
	//c = c.M().M().S().S().U().U()
	
	d := oll.OLLCase(27)//rubik.NewSolvedCube()
	//d = d.B().B().L().Uc()
	fmt.Printf("matches?\n%s\n%s\n\t%v\n", c, d, rubik.F2LMatches(c,d))
	//d = d.Rc().Rc().Rc()
	//fmt.Printf("matches?\n%s\n%s\n\t%v\n", c, d, rubik.F2LMatches(c,d))
}
