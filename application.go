package main

import (
	BinTree "A_DS3/BinaryTree"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"unicode"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

func getStringFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	var content string
	scanner.Scan()
	content += scanner.Text()
	for scanner.Scan() {
		content += "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func parseToTree(expression string) (*BinTree.Node, error, int) {
	var digit string
	leftSonExpected := false
	rightSonExpected := false
	bracketBalance := 0
	sonIdxStart := 0
	sonIdxEnd := 0
	twoSons := false
	var err error
	node := &BinTree.Node{}
	// parse the expression
	for i, token := range expression {
		// parsing SubTree from the expression
		if bracketBalance != 1 && (leftSonExpected || rightSonExpected) {
			if token == '(' {
				bracketBalance++
			} else if token == ')' {
				bracketBalance--
				// subtree found, recursively parse subtree
				if bracketBalance == 1 {
					sonIdxEnd = i
					if leftSonExpected {
						leftSonExpected = false
						node.Left, err, _ = parseToTree(expression[sonIdxStart : sonIdxEnd+1])
						rightSonExpected = true
						if err != nil {
							return nil, err, 0
						}
					} else if rightSonExpected {
						rightSonExpected = false
						node.Right, err, _ = parseToTree(expression[sonIdxStart : sonIdxEnd+1])
						twoSons = true
						if err != nil {
							return nil, err, 0
						}
					}
				}
			}
			continue
		}
		// find subtrees & Node key
		switch {
		case unicode.IsDigit(token):
			digit += string(token)
		case token == ' ':
			if digit != "" && !leftSonExpected && !rightSonExpected {
				node.Key, _ = strconv.Atoi(digit)
				digit = ""
			} else {
				continue
			}
		case token == '(':
			if expression[i+1] == '(' {
				return nil, fmt.Errorf("incorrect entry of brackets"), 0
			}
			bracketBalance++
			if twoSons {
				return nil, fmt.Errorf("too many clindren"), 0
			}
			if digit != "" && !leftSonExpected && !rightSonExpected {
				node.Key, _ = strconv.Atoi(digit)
				digit = ""
			}
			if rightSonExpected && bracketBalance == 2 {
				sonIdxStart = i
			} else if !rightSonExpected && bracketBalance == 2 {
				leftSonExpected = true
				sonIdxStart = i
			}
		case token == ')':
			if digit != "" && !leftSonExpected && !rightSonExpected {
				node.Key, _ = strconv.Atoi(digit)
				digit = ""
			}
			bracketBalance--
		default:
			return nil, fmt.Errorf("invalid character: %c", token), 0
		}
	}
	if expression[len(expression)-1] != ')' {
		return nil, fmt.Errorf("missing closing bracket"), 0
	}
	return node, nil, bracketBalance
}

// Applications
func redBlackTreeApplication() {
}

func binaryTreeApplication() {
	//Extracting string with tree from the file
	var filename string
	// File named as expression.txt
	fmt.Print("Enter filename: ")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println("Error reading filename")
		return
	}
	content, err := getStringFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	//Parsing string into the tree
	_, err, bracketBalance := parseToTree(content)
	if err != nil || bracketBalance != 0 {
		fmt.Printf("Error parsing string to tree: %v\n", err)
		return
	}
}

// main application
func application() {
	var choice int
	for {
		fmt.Print("Main menu:\n")
		fmt.Println("1. Make Binary Tree")
		fmt.Println("2. Make Red-Black Tree")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		switch choice {
		case 1:
			ClearScreen()
			binaryTreeApplication()
			break
		case 2:
			redBlackTreeApplication()
			break
		case 3:
			return
		}
	}
}
