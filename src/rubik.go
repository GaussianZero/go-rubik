package main

import (
	"fmt"
	"./rubik"
	"./oll"
)

func main() {
	cube := rubik.NewCube("sssssssssqqqqqqqqqsssssssssqqqqqqqqqsssssssssqqqqqqqqq")
	fmt.Printf("Cube:\n%s %v\n", cube, cube.IsSolved())

	solved := rubik.NewSolvedCube()
	//fmt.Printf("Solved cube\n%s %v\n", solved, solved.IsSolved())

	copied := solved.Copy().Fs().Fsc()
	fmt.Printf("copied:\n%s %v\n", copied, copied.IsSolved())

	c := PLL(solved.Copy())
	fmt.Printf("PLL'ed:\n%s %v\n", c, c.IsSolved())
	fmt.Println("is OLL case? ", oll.IsOLLCase(c, 1))
}

func PLL(c rubik.Cube) rubik.Cube {
	// J Perm
	c = c.R().U().Rc().Fc().R().U().Rc().Uc().Rc().F().R().R().Uc().Rc().Uc()

	// Sune
	//c = c.R().U().Rc().U().R().U().U().Rc()

	// OLL Case 1
	c = c.F().Rc().Fc().R().U().U().F().Rc().Fc().R().R().U().U().Rc()

	return c
}

// Cannot be solved by this program. Probably about 5 moves to solution.
// func TestSolveFunnyLooking(t *testing.T) {
// 	origin := NewCube("ggggwgggg rrrrgrrrr wwwwrwwww ooooboooo yyyyoyyyy bbbbybbbb")
// 	moves := Solve(origin)
// 	fmt.Printf("%s -> %s\n", origin, moves)
// }
