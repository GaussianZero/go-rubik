package main

import (
	"fmt"
	"./rubik"
	"./oll"
	"math/rand"
	"time"
)

const (
)

type Scramble struct {
	Moves []rubik.Move
	Cube rubik.Cube
	rep string
}

func randomTurn(c rubik.Cube, prevMove *rubik.Move) (rubik.Cube, rubik.Move) {
	n := rand.Int() % len(rubik.Moves)
	move := rubik.Moves[n]
	for prevMove != nil && rubik.IsCounter(move, *prevMove) {
		n = rand.Int() % len(rubik.Moves)
		move = rubik.Moves[n]
	}
	return c.Turn(move), move
}

func findScramble(n int) string {
	// Scramble from solved OLL state, see if we get any hits.
	reverseScrambles := []Scramble{}
	for i := 0; i < NUM_REVERSE_SCRAMBLES; i++ {
		caseN := oll.OLLCase(n)
		moves := []rubik.Move{}
		scramble := Scramble{}
		for j := 0; j < (rand.Int() % NUM_RANDOM_MOVES_HALF + NUM_RANDOM_MOVES_HALF); j++ {
			var prevMove *rubik.Move
			if len(moves) > 0 {
				prevMove = &moves[len(moves)-1]
			}
			newC, m := randomTurn(caseN, prevMove)
			caseN = newC
			moves = append(moves, m)
		}
		scramble.Moves = moves
		scramble.Cube = caseN.Copy()
		reverseScrambles = append(reverseScrambles, scramble)
		//fmt.Println("aww...")
	}
	//fmt.Println("Done generating cubes.")
	
	// Scramble a bunch of random cubes.
	possibleScrambles := []Scramble{}
	
	for i := 0; i < NUM_FORWARD_SCRAMBLES; i++ {
		if i % 10 == 0 {
			//fmt.Printf("searched %d cubes.\n", i)
		}
		c := rubik.NewSolvedCube()
		moves := []rubik.Move{}
		scramble := Scramble{}
		for j := 0; j < NUM_RANDOM_MOVES; j++ {
			var prevMove *rubik.Move
			if len(moves) > 0 {
				prevMove = &moves[len(moves)-1]
			}
			newC, m := randomTurn(c, prevMove)
			c = newC
			moves = append(moves, m)

			// Check if it matches any scrambles.
			for _, s := range reverseScrambles {
				if len(s.Cube) == 0 {
					panic("uh oh")
				}
				//fmt.Println(c, s.Cube)
				ollMatch, f2lMatch := oll.OLLDeepMatch(c, s.Cube)
				if ollMatch {
					fmt.Println("new cube:    ", c, moves)
					fmt.Println("OLL'ed cube: ", s.Cube, s.Moves)
				}
				if f2lMatch {
					oll.OLLDebugDeepMatch(c, s.Cube, true)
					fmt.Println("new cube:    ", c, moves)
					fmt.Println("OLL'ed cube: ", s.Cube, s.Moves)
					fmt.Println("woohoo!")
					printScramble(moves, s.Moves)
				}
			}
		}
		
		scramble.Moves = moves
		scramble.Cube = c.Copy()
		if len(scramble.Cube.String()) > 0 {
			// Huh, they can be empty?...
			possibleScrambles = append(possibleScrambles, scramble)
		}
		//fmt.Println(moves)
	}

	return "aww"
}

func printScramble(fMoves, bMoves []rubik.Move) {
	for _, m := range fMoves {
		fmt.Println(m)
	}
	for i := range bMoves {
		fmt.Println("reverse: ", bMoves[len(bMoves) - i - 1])
	}
}

func main() {
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	fmt.Println("Hello Cube")

	for i := 0; i < 100; i++ {
		scramble := findScramble(1)
		fmt.Println("found scramble: ", scramble)
	}
}
