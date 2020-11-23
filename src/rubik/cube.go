package rubik

import (
	"fmt"
	"strings"
)

// Representation of the Rubik's Cube.

// Faces colors.
type Color rune
const (
	BLUE   = 'b'
	WHITE  = 'w'
	YELLOW = 'y'
	ORANGE = 'o'
	GREEN  = 'g'
	RED    = 'r'
)

// Moves defined using Singmaster's notation.
// https://en.wikipedia.org/wiki/Rubik%27s_Cube#Move_notation
type Move string

var Moves = []Move{
	UP, UP_COUNTER, DOWN, DOWN_COUNTER, LEFT, LEFT_COUNTER,
	RIGHT, RIGHT_COUNTER, FRONT, FRONT_COUNTER, BACK, BACK_COUNTER,
}
var counterMoves = map[Move]Move {
	UP: UP_COUNTER,
	RIGHT: RIGHT_COUNTER,
	LEFT: LEFT_COUNTER,
	DOWN: DOWN_COUNTER,
	FRONT: FRONT_COUNTER,
	BACK: BACK_COUNTER,
}

func IsCounter(m1, m2 Move) bool {
	return counterMoves[m1] == m2 || counterMoves[m2] == m1
}

// Edge vs corner
type Shape string
const (
	EDGE = "edge"
	CORNER = "corner"
	CENTER = "center"
)

type Piece struct {
	Shape Shape
	Colors []Color
}

// Note - yellow is on bottom
var f2lPieces = []Piece{
	Piece{EDGE, []Color{GREEN, RED}},
	Piece{EDGE, []Color{GREEN, YELLOW}},
	Piece{EDGE, []Color{GREEN, ORANGE}},
	Piece{EDGE, []Color{RED,BLUE}},
	Piece{EDGE, []Color{RED,YELLOW}},
	Piece{EDGE, []Color{BLUE, ORANGE}},
	Piece{EDGE, []Color{BLUE, YELLOW}},
	Piece{EDGE, []Color{ORANGE, YELLOW}},
	Piece{CORNER, []Color{YELLOW,RED,GREEN}},
	Piece{CORNER, []Color{YELLOW,RED,BLUE}},
	Piece{CORNER, []Color{YELLOW,BLUE,ORANGE}},
	Piece{CORNER, []Color{YELLOW,ORANGE,GREEN}},
}

// Represents a "location" on the cube, at which a piece can be. Locations have Shapes too (edge, corner, center)
type Location struct {
	Shape Shape
	Indices []int // The array of indexes of this "location"
}


// Orientation is clockwise from facing W from the top, then facing Y after rotating like x2
var corners = []Location{
	// W corners
	Location{CORNER, []int{2,20,27}},
	Location{CORNER, []int{8,11,18}},
	Location{CORNER, []int{6,9,38}},
	Location{CORNER, []int{0,29,36}},
	// Y corners
	Location{CORNER, []int{17,24,47}},
	Location{CORNER, []int{26,33,53}},
	Location{CORNER, []int{35,42,51}},
	Location{CORNER, []int{15,44,45}},
}

var centers = []Location{
	
}

// listed in orientation of clockwise when facing each face. In order of faces W, G, R, B, O, Y. no dupes.
var edges = []Location{
	// W edges
	Location{EDGE, []int{1,28}},
	Location{EDGE, []int{5,19}},
	Location{EDGE, []int{7,10}},
	Location{EDGE, []int{3,37}},

	// G edges
	Location{EDGE, []int{14,21}},
	Location{EDGE, []int{16,46}},
	Location{EDGE, []int{12,41}},

	// R edges
	Location{EDGE, []int{23,30}},
	Location{EDGE, []int{25,50}},

	// B edges
	Location{EDGE, []int{32,39}},
	Location{EDGE, []int{34,52}},

	// G edges
	Location{EDGE, []int{43,48}},


	
}

const (
	UP            = "U"
	UP_COUNTER    = "U'"
	DOWN          = "D"
	DOWN_COUNTER  = "D'"
	LEFT          = "L"
	LEFT_COUNTER  = "L'"
	RIGHT         = "R"
	RIGHT_COUNTER = "R'"
	FRONT         = "F"
	FRONT_COUNTER = "F'"
	BACK          = "B"
	BACK_COUNTER  = "B'"
)

// See Cube.pdf for an illustrated version.
// The cube is a []rune where each face is stored as follows:
//
//              #3 BLU
//             35 34 33
//             32 31 30
//             29 28 27
//
//  #4 ORG      #0 WHT     #2 RED
// 42 39 36     0  1  2    20 23 26
// 43 40 37     3  4  5    19 22 25
// 44 41 38     6  7  8    18 21 24
//
//              #1 GRN
//              9 10 11
//             12 13 14
//             15 16 17
//
//              #5 YLW
//             45 46 47
//             48 49 50
//             51 52 53
type Cube []Color

// During a move, a facet is moved `From` `To`.
type Pair struct {
	From, To int
}

// Build a new Cube from the given string representation.
//
// For instance:
// "bwwbwwyyr orwogbygb gbbrrwrrw ooygbygbr oogoogyyb rrwgywgyo"
// with white, green, red, blue, orange and yellow faces as described earlier.
// "wwwwwwwww ggggggggg rrrrrrrrr bbbbbbbbb ooooooooo yyyyyyyyy"
// would correspond to a solved cube.
// No validity check is performed.
func NewCube(s string) Cube {
	stripped := strings.Replace(s, " ", "", -1)
	if len(stripped) != 9*6 {
		panic("Not a valid cube: " + s)
	}
	cube := make([]Color, 9*6)
	for i, c := range stripped {
		cube[i] = Color(c)
	}
	return cube
}

// Build a new solved Cube.
func NewSolvedCube() Cube {
	cube := make([]Color, 9*6)
	colors := []Color{WHITE, GREEN, RED, BLUE, ORANGE, YELLOW}
	for i, color := range colors {
		for j := 0; j < 9; j++ {
			cube[i*9+j] = color
		}
	}
	return cube
}

// String representation.
func (cube Cube) String() string {
	var s string
	for i, c := range cube {
		if i != 0 && i%9 == 0 {
			s = s + " "
		}
		s = s + string(c)
	}
	return s
}

// Return `true` if the given face is solved,
// aka made of a unique color.
func faceIsSolved(face []Color) bool {
	ref := face[0]
	for _, color := range face {
		if color != ref {
			return false
		}
	}
	return true
}

// Return `true` is this cube is solved.
func (cube Cube) IsSolved() bool {
	for f := 0; f < 6*9; f = f + 9 {
		face := cube[f : f+9]
		if !faceIsSolved(face) {
			return false
		}
	}
	return true
}

// Check if two cubes are equal.
func (cube Cube) Equals(other Cube) bool {
	for i, c := range cube {
		if c != other[i] {
			return false
		}
	}
	return true
}

// Clone a cube.
func (cube Cube) Copy() Cube {
	clone := make([]Color, len(cube))
	copy(clone, cube)
	return clone
}

// Turn the cube using the given move, and return a new Cube.
func (cube Cube) Turn(move Move) Cube {
	switch move {
	case UP:
		return cube.U()
	case UP_COUNTER:
		return cube.Uc()
	case DOWN:
		return cube.D()
	case DOWN_COUNTER:
		return cube.Dc()
	case LEFT:
		return cube.L()
	case LEFT_COUNTER:
		return cube.Lc()
	case RIGHT:
		return cube.R()
	case RIGHT_COUNTER:
		return cube.Rc()
	case FRONT:
		return cube.F()
	case FRONT_COUNTER:
		return cube.Fc()
	case BACK:
		return cube.B()
	case BACK_COUNTER:
		return cube.Bc()
	default:
		panic("Unknown move: " + move)
	}
}

// Return a new Cube after Pair.To becomes Pair.From.
func (cube Cube) exchange(pairs []*Pair) Cube {
	clone := cube.Copy()
	for _, pair := range pairs {
		clone[pair.To] = cube[pair.From]
	}
	return clone
}

// Rotate a face right, starting from `index`.
func (cube Cube) rotateRight(index int) Cube {
	clone := cube.Copy()
	displacements := []int{6, 3, 0, 7, 4, 1, 8, 5, 2}
	for i, d := range displacements {
		clone[i+index] = cube[d+index]
	}
	return clone
}

// Rotate a face left, starting from `index`.
func (cube Cube) rotateLeft(index int) Cube {
	clone := cube.Copy()
	displacements := []int{2, 5, 8, 1, 4, 7, 0, 3, 6}
	for i, d := range displacements {
		clone[i+index] = cube[d+index]
	}
	return clone
}

// Move Front face clockwise, and return a new Cube.
func (cube Cube) F() Cube {
	pairs := []*Pair{
		&Pair{6, 18}, &Pair{7, 21}, &Pair{8, 24},
		&Pair{18, 47}, &Pair{21, 46}, &Pair{24, 45},
		&Pair{45, 38}, &Pair{46, 41}, &Pair{47, 44},
		&Pair{38, 8}, &Pair{41, 7}, &Pair{44, 6},
	}

	return cube.exchange(pairs).rotateRight(9)
}

// Move Front face counter clockwise, and return a new Cube.
func (cube Cube) Fc() Cube {
	pairs := []*Pair{
		&Pair{6, 44}, &Pair{7, 41}, &Pair{8, 38},
		&Pair{38, 45}, &Pair{41, 46}, &Pair{44, 47},
		&Pair{45, 24}, &Pair{46, 21}, &Pair{47, 18},
		&Pair{18, 6}, &Pair{21, 7}, &Pair{24, 8},
	}
	return cube.exchange(pairs).rotateLeft(9)
}

// Move Back face clockwise, and return a new Cube.
func (cube Cube) B() Cube {
	pairs := []*Pair{
		&Pair{0, 42}, &Pair{1, 39}, &Pair{2, 36},
		&Pair{36, 51}, &Pair{39, 52}, &Pair{42, 53},
		&Pair{53, 20}, &Pair{52, 23}, &Pair{51, 26},
		&Pair{20, 0}, &Pair{23, 1}, &Pair{26, 2},
	}
	return cube.exchange(pairs).rotateRight(27)
}

// Move Back face counter clockwise, and return a new Cube.
func (cube Cube) Bc() Cube {
	pairs := []*Pair{
		&Pair{0, 20}, &Pair{1, 23}, &Pair{2, 26},
		&Pair{20, 53}, &Pair{23, 52}, &Pair{26, 51},
		&Pair{51, 36}, &Pair{52, 39}, &Pair{53, 42},
		&Pair{36, 2}, &Pair{39, 1}, &Pair{42, 0},
	}
	return cube.exchange(pairs).rotateLeft(27)
}

// Move Upper face clockwise, and return a new Cube.
func (cube Cube) U() Cube {
	pairs := []*Pair{
		&Pair{9, 36}, &Pair{10, 37}, &Pair{11, 38},
		&Pair{36, 27}, &Pair{37, 28}, &Pair{38, 29},
		&Pair{27, 18}, &Pair{28, 19}, &Pair{29, 20},
		&Pair{18, 9}, &Pair{19, 10}, &Pair{20, 11},
	}
	return cube.exchange(pairs).rotateRight(0)
}

// Move Upper face counter clockwise, and return a new Cube.
func (cube Cube) Uc() Cube {
	pairs := []*Pair{
		&Pair{9, 18}, &Pair{10, 19}, &Pair{11, 20},
		&Pair{18, 27}, &Pair{19, 28}, &Pair{20, 29},
		&Pair{27, 36}, &Pair{28, 37}, &Pair{29, 38},
		&Pair{36, 9}, &Pair{37, 10}, &Pair{38, 11},
	}
	return cube.exchange(pairs).rotateLeft(0)
}

// Move Down face clockwise, and return a new Cube.
func (cube Cube) D() Cube {
	pairs := []*Pair{
		&Pair{15, 24}, &Pair{16, 25}, &Pair{17, 26},
		&Pair{24, 33}, &Pair{25, 34}, &Pair{26, 35},
		&Pair{33, 42}, &Pair{34, 43}, &Pair{35, 44},
		&Pair{42, 15}, &Pair{43, 16}, &Pair{44, 17},
	}
	return cube.exchange(pairs).rotateRight(45)
}

// Move Down face counter clockwise, and return a new Cube.
func (cube Cube) Dc() Cube {
	pairs := []*Pair{
		&Pair{15, 42}, &Pair{16, 43}, &Pair{17, 44},
		&Pair{42, 33}, &Pair{43, 34}, &Pair{44, 35},
		&Pair{33, 24}, &Pair{34, 25}, &Pair{35, 26},
		&Pair{24, 15}, &Pair{25, 16}, &Pair{26, 17},
	}
	return cube.exchange(pairs).rotateLeft(45)
}

// Move Left face clockwise, and return a new Cube.
func (cube Cube) L() Cube {
	pairs := []*Pair{
		&Pair{0, 9}, &Pair{3, 12}, &Pair{6, 15},
		&Pair{9, 45}, &Pair{12, 48}, &Pair{15, 51},
		&Pair{45, 35}, &Pair{48, 32}, &Pair{51, 29},
		&Pair{29, 6}, &Pair{32, 3}, &Pair{35, 0},
	}
	return cube.exchange(pairs).rotateRight(36)
}

// Move Left face counter clockwise, and return a new Cube.
func (cube Cube) Lc() Cube {
	pairs := []*Pair{
		&Pair{0, 35}, &Pair{3, 32}, &Pair{6, 29},
		&Pair{29, 51}, &Pair{32, 48}, &Pair{35, 45},
		&Pair{45, 9}, &Pair{48, 12}, &Pair{51, 15},
		&Pair{9, 0}, &Pair{12, 3}, &Pair{15, 6},
	}
	return cube.exchange(pairs).rotateLeft(36)
}

// Move Right face clockwise, and return a new Cube.
func (cube Cube) R() Cube {
	pairs := []*Pair{
		&Pair{2, 33}, &Pair{5, 30}, &Pair{8, 27},
		&Pair{27, 53}, &Pair{30, 50}, &Pair{33, 47},
		&Pair{47, 11}, &Pair{50, 14}, &Pair{53, 17},
		&Pair{11, 2}, &Pair{14, 5}, &Pair{17, 8},
	}
	return cube.exchange(pairs).rotateRight(18)
}

// Move Right face counter clockwise, and return a new Cube.
func (cube Cube) Rc() Cube {
	pairs := []*Pair{
		&Pair{2, 11}, &Pair{5, 14}, &Pair{8, 17},
		&Pair{11, 47}, &Pair{14, 50}, &Pair{17, 53},
		&Pair{47, 33}, &Pair{50, 30}, &Pair{53, 27},
		&Pair{27, 8}, &Pair{30, 5}, &Pair{33, 2},
	}
	return cube.exchange(pairs).rotateLeft(18)
}

// Represents a "Middle" move
func (cube Cube) M() Cube {
	pairs := []*Pair{
		&Pair{1, 10}, &Pair{4,13}, &Pair{7, 16}, // done
		&Pair{10, 46}, &Pair{13, 49}, &Pair{16, 52}, // done
		&Pair{46, 34}, &Pair{49, 31}, &Pair{52, 28}, // done
		&Pair{28, 7}, &Pair{31, 4}, &Pair{34, 1}, // done
	}
	return cube.exchange(pairs)
}

// Represents a "Middle" move
func (cube Cube) Mc() Cube {
	pairs := []*Pair{
		&Pair{1, 34}, &Pair{4, 31}, &Pair{7, 28},
		&Pair{28, 52}, &Pair{31, 49}, &Pair{34, 46},
		&Pair{46, 10}, &Pair{49, 13}, &Pair{52, 16},
		&Pair{10, 1}, &Pair{13, 4}, &Pair{16, 7},
	}
	return cube.exchange(pairs)
}

func (cube Cube) S() Cube {
	pairs := []*Pair{
		&Pair{3, 19}, &Pair{4, 22}, &Pair{5, 25}, // done
		&Pair{19, 50}, &Pair{22, 49}, &Pair{25, 48}, // done
		&Pair{48, 37}, &Pair{49, 40}, &Pair{50, 43}, // done
		&Pair{37, 5}, &Pair{40, 4}, &Pair{43, 3}, // done
	}
	return cube.exchange(pairs)
}

func (cube Cube) Sc() Cube {
	pairs := []*Pair{
		&Pair{3, 43}, &Pair{4, 40}, &Pair{5, 37}, // done
		&Pair{43, 50}, &Pair{40, 49}, &Pair{37, 48}, // done
		&Pair{48, 25}, &Pair{49, 22}, &Pair{50, 19}, // done
		&Pair{25, 5}, &Pair{22, 4}, &Pair{19, 3}, // done
	}
	return cube.exchange(pairs)
}

func (cube Cube) E() Cube {
	pairs := []*Pair{
		&Pair{12, 21}, &Pair{13, 22}, &Pair{14, 23}, // done
		&Pair{21, 30}, &Pair{22, 31}, &Pair{23, 32}, // done
		&Pair{30, 39}, &Pair{31, 40}, &Pair{32, 41}, // done
		&Pair{39, 12}, &Pair{40, 13}, &Pair{41, 14}, // done
	}
	return cube.exchange(pairs).rotateRight(45)
}

func (cube Cube) Ec() Cube {
	pairs := []*Pair{
		&Pair{15, 42}, &Pair{16, 43}, &Pair{17, 44},
		&Pair{42, 33}, &Pair{43, 34}, &Pair{44, 35},
		&Pair{33, 24}, &Pair{34, 25}, &Pair{35, 26},
		&Pair{24, 15}, &Pair{25, 16}, &Pair{26, 17},
	}
	return cube.exchange(pairs).rotateLeft(45)
}



// Represents a "small r" move.
func (cube Cube) Rs() Cube {
	return cube.R().Mc()
}

// Represents a "small r'" move.
func (cube Cube) Rsc() Cube {
	return cube.Rc().M()
}

// Represents a "small l" move.
func (cube Cube) Ls() Cube {
	return cube.L().M()
}

// Represents a "small l'" move.
func (cube Cube) Lsc() Cube {
	return cube.Lc().Mc()
}

// Represents a "small f" move.
func (cube Cube) Fs() Cube {
	return cube.F().S()
}

// Represents a "small f'" move.
func (cube Cube) Fsc() Cube {
	return cube.Fc().Sc()
}

func (cube Cube) X() Cube {
	return cube.R().Mc().Lc()
}

func (cube Cube) Xc() Cube {
	return cube.Rc().M().L()
}

func (cube Cube) X() Cube {
	return cube.R().Mc().Lc()
}



// Below is F2L stuff, probably move to it's own file


func (cube Cube) ColorAt(n int) Color {
	// TODO Implement
	return cube[n]
}

func pieceMatches(e Location, p Piece, c Cube) bool {
	//
	for _, index := range e.Indices {
		//fmt.Println("Checking if piece at index ", index, " matches")
		colorFound := false
		for _, color := range p.Colors {
			if c.ColorAt(index) == color {
				colorFound = true
			}
		}
		if !colorFound {
			return false
		}
	}
	//fmt.Println("matches!")
	return true
}

func isInSameSlotAndOrientation(p Piece, c, d Cube) bool {
	fmt.Println("check piece!")
	var pieces []Location
	switch(p.Shape) {
	case EDGE:
		pieces = edges
	case CORNER:
		pieces = corners
	case CENTER:
		pieces = centers
	}
	for _, e := range pieces {
		//fmt.Println("checking pieces")
		if !pieceMatches(e, p, c) {
			continue
		}

		// We found our piece in cube C, check if location and orientation of the piece match in C and D.
		colorsMatch := true
		for _, location := range e.Indices {
			//fmt.Println("checking loc: ", location)
			if c.ColorAt(location) != d.ColorAt(location) {
				colorsMatch = false
			}
		}
		if !colorsMatch {
			return false
		}
	}
	return true
}

func isCornerInSameSlot(p Piece, c, d Cube) bool {
	return false
}

func isCenterInSameSlot(p Piece, c, d Cube) bool {
	return false
}

func (c Color) String() string{
	return fmt.Sprintf("%c", c)
}

// F2L related stuff. TODO: move to its own library
func F2LMatches(c, d Cube) bool {
	numMatchingPieces := 0
	for _, piece := range f2lPieces {
		//fmt.Println("piece: ", piece)
		if !isInSameSlotAndOrientation(piece, c, d) {
			fmt.Printf("Pieces Don't Match: %s\n\t%s\n\t%s\n", piece, c, d)
			fmt.Printf("%d pieces did match.\n", numMatchingPieces)
			return false
		} else {
			numMatchingPieces++
		}
	}
	fmt.Println("WHOA!!!.")
	return true
}

func (cube Cube) DebugEdges() {
	for i, edge := range edges {
		fmt.Printf("\t%d, {%d,%d} = {%c,%c}\n",
			i,
			edge.Indices[0], edge.Indices[1],
			cube.ColorAt(edge.Indices[0]),
			cube.ColorAt(edge.Indices[1]),
		)
	}
}
