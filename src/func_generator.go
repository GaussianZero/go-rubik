package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	outputScramble := ""
	nextToken := ""
	outputTokens := []string{}
	//prevToken := ""
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		text := scanner.Text()
		fields := strings.Split(text, ":")
		caseNum := fields[0]
		//fmt.Println(caseNum)
		leftoverText := fields[1]

		outputScramble = fmt.Sprintf("func case%s (c rubik.Cube) rubik.Cube {\n\tc = c.", caseNum)
		
		for _, char := range  leftoverText + " " {
			//fmt.Printf("IN[%d] = %c\n", pos, char)
			switch (fmt.Sprintf("%c", char)) {
			case " ":
				//outputScramble += nextToken + "."
				outputTokens = append(outputTokens, nextToken + ".")
				nextToken = ""
				continue
			case "R":
				nextToken += "R()"
			case "U":
				nextToken += "U()"
			case "F":
				nextToken += "F()"
			case "D":
				nextToken += "D()"
			case "B":
				nextToken += "B()"
			case "M":
				nextToken += "M()"
			case "L":
				nextToken += "L()"
			case "r":
				nextToken += "Rs()"
			case "u":
				nextToken += "Us()"
			case "f":
				nextToken += "Fs()"
			case "d":
				nextToken += "Ds()"
			case "b":
				nextToken += "Bs()"
			case "l":
				nextToken += "Ls()"
			case "2":
				//fmt.Println("squaring...", nextToken)
				nextToken = nextToken + "." + nextToken
				//fmt.Println("squared: ", nextToken)
			case "'":
				//fmt.Println("reversing...", nextToken)
				nextToken = reverseToken(nextToken)
				//fmt.Println("reversed: ", nextToken)
			default:
				log.Fatalf("Unknown value! %c", char)
			}
		}
	}
	
	// reverse the algo
	// for _, token := range outputTokens {
	// 	fmt.Printf("%s", token)
	// }
	// fmt.Println()
	for i := range outputTokens {
		token := outputTokens[len(outputTokens) - i - 1]
		rToken := reverseToken(token)
		//fmt.Printf("reverse[ %s ] = %s\n", token, rToken)
		outputScramble += rToken
	}
	outputScramble = outputScramble[:len(outputScramble)-1]

	fmt.Println(outputScramble + "\n\treturn c\n}\n")
	
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}


func reverseToken(s string) string {
	if strings.Contains(s, "c") {
		s = strings.ReplaceAll(s, "c", "")
	} else {
		s = strings.ReplaceAll(s, "()", "c()")
	}
	return s
}
