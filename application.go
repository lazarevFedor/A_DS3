package main

import (
	Binary "A_DS3/BinaryTree"
	RedBlack "A_DS3/RedBlackTree"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

// clearInputBuffer reads and discards any unread input from the standard input buffer.
func clearInputBuffer() {
	// Создаем новый Reader для работы с буфером ввода
	reader := bufio.NewReader(os.Stdin)
	// Читаем и игнорируем все символы до конца строки
	_, _ = reader.ReadString('\n')
}

// clearInputBuffer reads and discards any unread input from the standard input buffer.
func waitForKeyPress() {
	fmt.Println("Press any key to continue...")
	clearInputBuffer()
	var waitFor string
	fmt.Scanln(&waitFor)
}

func removeExtraSpaces(input string) string {
	// Разбиваем строку на части по пробелам
	words := strings.Fields(input)
	// Соединяем части обратно, вставляя ровно один пробел между ними
	result := strings.Join(words, " ")
	return result
}

// ClearScreen clears the console screen on Windows systems.
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

// getStringFromFile reads the contents of a file and returns it as a string.
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

// parseToTree parses an arithmetic expression in infix notation into a binary tree.
func parseToTree(expression string) (*Binary.Node, error, int) {
	var digit string
	leftSonExpected := false
	rightSonExpected := false
	bracketBalance := 0
	sonIdxStart := 0
	sonIdxEnd := 0
	twoSons := false
	var err error
	node := &Binary.Node{}
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

// redBlackTreeApplication - red-black tree application.
func redBlackTreeApplication(travers string) {
	var choise int
	var output string
	tree := RedBlack.NewRBTree()
	for {
		fmt.Println("Red-Black Tree Application Menu")
		fmt.Println("1. Create tree")
		fmt.Println("2. Insert element")
		fmt.Println("3. Delete element")
		fmt.Println("4. Search element")
		fmt.Println("5. In-order traversal")
		fmt.Println("6. Pre-order traversal")
		fmt.Println("7. Post-order traversal")
		fmt.Println("8. Level-Order traversal")
		fmt.Println("9. Clear screen")
		fmt.Println("10. Print tree")
		fmt.Println("11. Clear tree")
		fmt.Println("12. Back")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choise)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		switch choise {
		case 1:
			var numbers string
			fmt.Print("\n1 - make from Binary Tree\n2 - make from console output\nEnter your choice: ")
			_, err = fmt.Scan(&choise)
			clearInputBuffer()
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if choise == 1 {
				numbers = travers
			} else if choise == 2 {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter binary tree (space separated values): ")
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid binary tree.")
					continue
				}
				numbers = removeExtraSpaces(input)
			} else {
				fmt.Println("Invalid choice. Please try again.")
				continue
			}
			var digit string
			var num int
			tree.Clear()
			for _, v := range numbers {
				if unicode.IsDigit(v) {
					digit += string(v)
				} else if digit != "" {
					num, _ = strconv.Atoi(digit)
					tree.Insert(num)
					digit = ""
				}
			}
			if digit != "" {
				num, _ = strconv.Atoi(digit)
				tree.Insert(num)
				digit = ""
			}
			fmt.Println("Red-Black Tree successfully created.")
			waitForKeyPress()
		case 2:
			var key int
			fmt.Print("Enter key to insert: ")
			_, err := fmt.Scan(&key)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			tree.Insert(key)
			fmt.Printf("Element %d inserted successfully.\n", key)
			waitForKeyPress()
		case 3:
			var key int
			fmt.Print("Enter key to delete: ")
			_, err := fmt.Scan(&key)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			err = tree.Delete(key)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Printf("Element %d deleted successfully.\n", key)
			}
			waitForKeyPress()
		case 4:
			var key int
			fmt.Print("Enter key to search: ")
			_, err := fmt.Scan(&key)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if tree.Search(key) != nil {
				fmt.Printf("Element %d found in the tree.\n", key)
			} else {
				fmt.Printf("Element %d not found in the tree.\n", key)
			}
			waitForKeyPress()
		case 5:
			fmt.Println("In-order traversal:")
			RedBlackTreeTravers := tree.InOrderTravers(tree.Root)
			fmt.Println(RedBlackTreeTravers)
			waitForKeyPress()
		case 6:
			fmt.Println("Pre-order traversal:")
			RedBlackTreeTravers := tree.PreOrderTravers(tree.Root)
			RedBlackTreeTravers = removeExtraSpaces(RedBlackTreeTravers)
			fmt.Println(RedBlackTreeTravers)
			waitForKeyPress()
		case 7:
			fmt.Println("Post-order traversal:")
			RedBlackTreeTravers := tree.PostOrderTravers(tree.Root)
			RedBlackTreeTravers = removeExtraSpaces(RedBlackTreeTravers)
			fmt.Println(RedBlackTreeTravers)
			waitForKeyPress()
		case 8:
			fmt.Println("Level-order traversal:")
			RedBlackTreeTravers := tree.LevelOrderTravers(tree.Root)
			RedBlackTreeTravers = removeExtraSpaces(RedBlackTreeTravers)
			fmt.Println(RedBlackTreeTravers)
			waitForKeyPress()
		case 9:
			ClearScreen()
			continue
		case 10:
			ClearScreen()
			RedBlack.Output(tree.Root, "", false, &output)
			fmt.Print(output)
			output = ""
			waitForKeyPress()
		case 11:
			tree.Clear()
		case 12:
			return
		}
	}
}

// binaryTreeApplication handles the binary tree creation process.
func binaryTreeApplication() string {
	tree := Binary.NewBinTree()
	bracketBalance := 0
	//Extracting string with tree from the file
	var filename string
	// File named as expression.txt
	fmt.Print("Enter filename: ")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println("Error reading filename")
		return ""
	}
	content, err := getStringFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	//Parsing string into the tree
	tree.Root, err, bracketBalance = parseToTree(content)
	if err != nil || bracketBalance != 0 {
		fmt.Printf("Error parsing string to tree: %v\n", err)
		return ""
	}
	fmt.Println("Binary Tree successfully created.")

	travers := tree.PreOrderTravers(tree.Root)
	travers = removeExtraSpaces(travers)
	fmt.Println("Binary Tree:", travers)
	return travers
}

// application - This function serves as the main entry point for the application.
func application() {
	var choice int
	var travers string
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
			travers = ""
			travers = binaryTreeApplication()
			break
		case 2:
			ClearScreen()
			redBlackTreeApplication(travers)
			travers = ""
			break
		case 3:
			return
		}
	}
}
