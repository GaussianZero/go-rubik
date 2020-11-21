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

func case2 (c rubik.Cube) rubik.Cube {
	c = c.Rs().U().Rc().Uc().Uc().R().Uc().Uc().Rsc().Uc().Uc().Rs().Uc().Rsc()
	return c
}

func case3 (c rubik.Cube) rubik.Cube {
	c = c.M().Uc().Rs().Uc().Uc().Rsc().Uc().R().Uc().Rc().Rc().Rs()
	return c
}

func case4 (c rubik.Cube) rubik.Cube {
	c = c.M().R().U().Rc().U().Rs().Uc().Uc().Rsc().U().Mc()
	return c
}

func case5 (c rubik.Cube) rubik.Cube {
	c = c.Lsc().Uc().L().Uc().Lc().Uc().Uc().Ls()
	return c
}

func case6 (c rubik.Cube) rubik.Cube {
	c = c.Rs().U().Rc().U().R().Uc().Uc().Rsc()
	return c
}

func case7 (c rubik.Cube) rubik.Cube {
	c = c.Rs().Uc().Uc().Rc().Uc().R().Uc().Rsc()
	return c
}
func case8 (c rubik.Cube) rubik.Cube {
	c = c.Lsc().Uc().Uc().L().U().Lc().U().Ls()
	return c
}

func case9 (c rubik.Cube) rubik.Cube {
	c = c.F().U().R().Uc().Rc().Rc().Fc().R().U().R().Uc().Rc()
	return c
}

func case10 (c rubik.Cube) rubik.Cube {
	c = c.R().Uc().Uc().Rc().F().Rc().Fc().R().Uc().R().Uc().Rc()
	return c
}
func case11 (c rubik.Cube) rubik.Cube {
	c = c.Rs().Uc().Uc().Rc().F().Rc().Fc().R().Uc().R().Uc().Rsc()
	return c
}
func case12 (c rubik.Cube) rubik.Cube {
	c = c.Rs().Rc().U().Rc().Uc().Uc().R().U().Rc().U().R().M()
	return c
}
func case13 (c rubik.Cube) rubik.Cube {
	c = c.R().U().Rc().Uc().Rc().F().Rc().Rc().U().Rc().Uc().Fc()
	return c
}
func case14 (c rubik.Cube) rubik.Cube {
	c = c.F().U().Fc().Rc().F().R().Uc().Rc().Fc().R()
	return c
}
func case15 (c rubik.Cube) rubik.Cube {
	c = c.Lsc().Uc().Ls().Uc().Lc().U().L().Lsc().U().Ls()
	return c
}
func case16 (c rubik.Cube) rubik.Cube {
	c = c.Rs().U().Rsc().U().R().Uc().Rc().Rs().Uc().Rsc()
	return c
}

func case17 (c rubik.Cube) rubik.Cube {
	c = c.M().U().R().U().Rc().Uc().Rs().Rc().Rc().F().R().Fc()
	return c
}
func case18 (c rubik.Cube) rubik.Cube {
	c = c.Rsc().Uc().Uc().R().U().Rc().U().Rsc().Rsc().Uc().Uc().Rc().Uc().R().Uc().Rsc()
	return c
}
func case19 (c rubik.Cube) rubik.Cube {
	c = c.F().Rc().Fc().R().M().U().R().Uc().Rc().Uc().Rc().Rs()
	return c
}
func case20 (c rubik.Cube) rubik.Cube {
	c = c.M().U().R().U().Rc().Uc().Mc().Mc().U().R().Uc().Rsc()
	return c
}
func case21 (c rubik.Cube) rubik.Cube {
	c = c.R().U().Rc().U().R().Uc().Rc().U().R().Uc().Uc().Rc()
	return c
}

func case22 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Uc().Uc().Rc().Rc().U().Rc().Rc().U().Rc().Rc().Uc().Uc().Rc()
	return c
}
func case23 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Uc().Uc().Rc().Dc().R().Uc().Uc().Rc().D().Rc().Rc()
	return c
}
func case24 (c rubik.Cube) rubik.Cube {
	c = c.F().Rc().Fc().Rs().U().R().Uc().Rsc()
	return c
}
func case25 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Fc().Rs().U().R().Uc().Rsc().F()
	return c
}

func case26 (c rubik.Cube) rubik.Cube {
	c = c.R().U().Rc().U().R().Uc().Uc().Rc()
	return c
}
func case27 (c rubik.Cube) rubik.Cube {
	c = c.R().Uc().Uc().Rc().Uc().R().Uc().Rc()
	return c
}

func case28 (c rubik.Cube) rubik.Cube {
	c = c.R().U().Rc().Uc().Rc().Rs().U().R().Uc().Rsc()
	return c
}
func case29 (c rubik.Cube) rubik.Cube {
	c = c.R().Uc().Rc().Fc().U().F().R().U().Rc().U().R().Uc().Rc()
	return c
}
func case30 (c rubik.Cube) rubik.Cube {
	c = c.Fc().Fc().R().Uc().Rc().U().R().U().Rc().Rc().Fc().R().Fc()
	return c
}

func case31 (c rubik.Cube) rubik.Cube {
	c = c.Rc().F().R().U().Rc().Uc().Fc().U().R()
	return c
}

func case32 (c rubik.Cube) rubik.Cube {
	c = c.L().Fc().Lc().Uc().L().U().F().Uc().Lc()
	return c
}
func case33 (c rubik.Cube) rubik.Cube {
	c = c.F().Rc().Fc().R().U().R().Uc().Rc()
	return c
}

func case34 (c rubik.Cube) rubik.Cube {
	c = c.F().U().Rc().Uc().Rc().Fc().R().U().Rc().Rc().Uc().Rc()
	return c
}
func case35 (c rubik.Cube) rubik.Cube {
	c = c.R().Uc().Uc().Rc().F().Rc().Fc().Rc().Rc().Uc().Uc().Rc()
	return c
}
func case36 (c rubik.Cube) rubik.Cube {
	c = c.Fc().L().F().Lc().Uc().Lc().Uc().L().U().Lc().U().L()
	return c
}
func case37 (c rubik.Cube) rubik.Cube {
	c = c.R().U().Rc().Uc().Rc().F().R().Fc()
	return c
}
func case38 (c rubik.Cube) rubik.Cube {
	c = c.F().Rc().Fc().R().U().R().U().Rc().Uc().R().Uc().Rc()
	return c
}

func case39 (c rubik.Cube) rubik.Cube {
	c = c.L().U().Fc().Uc().Lc().U().L().F().Lc()
	return c
}

func case40 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Uc().F().U().R().Uc().Rc().Fc().R()
	return c
}
func case41 (c rubik.Cube) rubik.Cube {
	c = c.F().U().R().Uc().Rc().Fc().R().Uc().Uc().Rc().Uc().R().Uc().Rc()
	return c
}

func case42 (c rubik.Cube) rubik.Cube {
	c = c.F().U().R().Uc().Rc().Fc().Rc().Uc().Uc().R().U().Rc().U().R()
	return c
}

func case43 (c rubik.Cube) rubik.Cube {
	c = c.Fc().Lc().Uc().L().U().F()
	return c
}
func case44 (c rubik.Cube) rubik.Cube {
	c = c.F().R().U().Rc().Uc().Fc()
	return c
}
func case45 (c rubik.Cube) rubik.Cube {
	c = c.F().U().R().Uc().Rc().Fc()
	return c
}

func case46 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Uc().F().Rc().Fc().R().U().R()
	return c
}

func case47 (c rubik.Cube) rubik.Cube {
	c = c.Rc().Uc().F().Rc().Fc().R().F().Rc().Fc().R().U().R()
	return c
}

func case48 (c rubik.Cube) rubik.Cube {
	c = c.F().U().R().Uc().Rc().U().R().Uc().Rc().Fc()
	return c
}

func case49 (c rubik.Cube) rubik.Cube {
	c = c.Rsc().U().Rsc().Rsc().Uc().Rsc().Rsc().Uc().Rsc().Rsc().U().Rsc()
	return c
}
func case50 (c rubik.Cube) rubik.Cube {
	c = c.Rs().Uc().Rsc().Rsc().U().Rsc().Rsc().U().Rsc().Rsc().Uc().Rs()
	return c
}

func case51 (c rubik.Cube) rubik.Cube {
	c = c.F().R().U().Rc().Uc().R().U().Rc().Uc().Fc()
	return c
}

func case52 (c rubik.Cube) rubik.Cube {
	c = c.R().B().U().Bc().U().Rc().Uc().R().Uc().Rc()
	return c
}

func case53 (c rubik.Cube) rubik.Cube {
	c = c.Lsc().Uc().L().Uc().Lc().U().L().Uc().Lc().Uc().Uc().Ls()
	return c
}

func case54 (c rubik.Cube) rubik.Cube {
	c = c.Rs().U().Rc().U().R().Uc().Rc().U().R().Uc().Uc().Rsc()
	return c
}

func case55 (c rubik.Cube) rubik.Cube {
	c = c.R().Uc().Rc().Uc().R().U().Rc().Rc().F().Rc().Rc().U().Rc().Uc().Rc().Fc().R()
	return c
}

func case56 (c rubik.Cube) rubik.Cube {
	c = c.Rsc().Uc().Rs().Rc().Uc().R().U().Rc().Uc().R().U().Rsc().U().Rs()
	return c
}
func case57 (c rubik.Cube) rubik.Cube {
	c = c.Rs().U().Rc().Uc().M().U().R().Uc().Rc()
	return c
}



var (
	ollCases = []rubik.Cube{
		rubik.NewSolvedCube(), // 0: Case 0 is the solved OLL
		case1(rubik.NewSolvedCube()),
		case2(rubik.NewSolvedCube()),
		case3(rubik.NewSolvedCube()),
		case4(rubik.NewSolvedCube()),
		case5(rubik.NewSolvedCube()),
		case6(rubik.NewSolvedCube()),
		case7(rubik.NewSolvedCube()),
		case8(rubik.NewSolvedCube()),
		case9(rubik.NewSolvedCube()),
		case10(rubik.NewSolvedCube()),
		case11(rubik.NewSolvedCube()),
		case12(rubik.NewSolvedCube()),
		case13(rubik.NewSolvedCube()),
		case14(rubik.NewSolvedCube()),
		case15(rubik.NewSolvedCube()),
		case16(rubik.NewSolvedCube()),
		case17(rubik.NewSolvedCube()),
		case18(rubik.NewSolvedCube()),
		case19(rubik.NewSolvedCube()),
		case20(rubik.NewSolvedCube()),
		case21(rubik.NewSolvedCube()),
		case22(rubik.NewSolvedCube()),
		case23(rubik.NewSolvedCube()),
		case24(rubik.NewSolvedCube()),
		case25(rubik.NewSolvedCube()),
		case26(rubik.NewSolvedCube()),
		case27(rubik.NewSolvedCube()),
		case28(rubik.NewSolvedCube()),
		case29(rubik.NewSolvedCube()),
		case30(rubik.NewSolvedCube()),
		case31(rubik.NewSolvedCube()),
		case32(rubik.NewSolvedCube()),
		case33(rubik.NewSolvedCube()),
		case34(rubik.NewSolvedCube()),
		case35(rubik.NewSolvedCube()),
		case36(rubik.NewSolvedCube()),
		case37(rubik.NewSolvedCube()),
		case38(rubik.NewSolvedCube()),
		case39(rubik.NewSolvedCube()),
		case40(rubik.NewSolvedCube()),
		case41(rubik.NewSolvedCube()),
		case42(rubik.NewSolvedCube()),
		case43(rubik.NewSolvedCube()),
		case44(rubik.NewSolvedCube()),
		case45(rubik.NewSolvedCube()),
		case46(rubik.NewSolvedCube()),
		case47(rubik.NewSolvedCube()),
		case48(rubik.NewSolvedCube()),
		case49(rubik.NewSolvedCube()),
		case50(rubik.NewSolvedCube()),
		case51(rubik.NewSolvedCube()),
		case52(rubik.NewSolvedCube()),
		case53(rubik.NewSolvedCube()),
		case54(rubik.NewSolvedCube()),
		case55(rubik.NewSolvedCube()),
		case56(rubik.NewSolvedCube()),
		case57(rubik.NewSolvedCube()),
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
