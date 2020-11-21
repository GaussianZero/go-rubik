package main

import (
	"fmt"
	"./rubik"
)

func case1 (c rubik.Cube) rubik.Cube {
	//	c = c.R().U().U().R().R().F().R().Fc().U().U().Rc().F().R().Fc()
	c = c.F().Rc().Fc().R().Uc().Uc().F().Rc().Fc().Rc().Rc().Uc().Uc().Rc()
	return c
}

func case27(c rubik.Cube) rubik.Cube{
	c = c.R().U().Rc().U().R().U().U().Rc()
	return c
}

var (
	ollCases = []rubik.Cube{
		rubik.NewSolvedCube(), // 0: Case 0 is the solved OLL
		case1(rubik.NewSolvedCube()),
		case27(rubik.NewSolvedCube()),
	}
)

func matchesCase(c, d rubik.Cube) bool {
	s1, s2 := string(c), string(d)
	
	for i, c := range s1 {
		if c == 111 {
			if s2[i] != 111 {
				return false
			}
		}
		//fmt.Println(c, i, s2[i])
		//fmt.Printf("%c, %d, %c\n", c, i, s2[i])
	}
	return true
}

func isOLLCase(c rubik.Cube, n int) bool {
	if n == 27 {
		n = 2
	}
	return matchesCase(c, ollCases[n])
}

func main() {
	cube := rubik.NewCube("sssssssssqqqqqqqqqsssssssssqqqqqqqqqsssssssssqqqqqqqqq")
	fmt.Printf("Cube:\n%s %v\n", cube, cube.IsSolved())

	solved := rubik.NewSolvedCube()
	//fmt.Printf("Solved cube\n%s %v\n", solved, solved.IsSolved())

	copied := solved.Copy().Fs().Fsc()
	fmt.Printf("copied:\n%s %v\n", copied, copied.IsSolved())

	c := PLL(solved.Copy())
	fmt.Printf("PLL'ed:\n%s %v\n", c, c.IsSolved())
	fmt.Println("is OLL case? ", isOLLCase(c, 1))
}

func PLL(c rubik.Cube) rubik.Cube {
	// J Perm
	c = c.R().U().Rc().Fc().R().U().Rc().Uc().Rc().F().R().R().Uc().Rc().Uc()

	// Sune
	//c = c.R().U().Rc().U().R().U().U().Rc()
	c = c.F().Rc().Fc().R().U().U().F().Rc().Fc().R().R().U().U().Rc()

	return c
}

// Cannot be solved by this program. Probably about 5 moves to solution.
// func TestSolveFunnyLooking(t *testing.T) {
// 	origin := NewCube("ggggwgggg rrrrgrrrr wwwwrwwww ooooboooo yyyyoyyyy bbbbybbbb")
// 	moves := Solve(origin)
// 	fmt.Printf("%s -> %s\n", origin, moves)
// }
