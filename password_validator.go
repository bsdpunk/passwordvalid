package main

import "fmt"
import "os"
import "bufio"
import "unicode"

func main() {
	var lines []string
	if len(os.Args[1]) > 0 {
		var err error
		lines, err = readLines(os.Args[1])
		if err != nil {
			fmt.Printf("Bad Weak Password File: %s", err)
			os.Exit(2)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//		fmt.Println(scanner.Text())
		switch pword := scanner.Text(); {
		case len(pword) < 8:
			fmt.Printf("%s -> Error: Too Short\n", pword)
		case isASCII(pword) != true:
			fmt.Printf("%s -> Error: Not ASCII\n", pword)
		case isCommon(lines, pword):
			fmt.Printf("%s -> Error: Common Password\n", pword)
		default:
			fmt.Printf("%s , Good!\n", pword)
		}
	}
}

//Stole this function
//Just cause it was simple
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

//Stole This Function
//I didn't want to write the boiler plate
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isCommon(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
