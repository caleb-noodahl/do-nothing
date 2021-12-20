package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	
	fmt.Println("step 1. github")
	fmt.Println("clone back the code via github cli")

	fmt.Println(" ex. git clone https://github.com/caleb-noodahl/do-nothing.git")

	input.Scan()
	fmt.Println("step 2. dependencies")
	fmt.Println("install the dependencies referenced in the go.mod file")

	fmt.Println(" ex. go mod download")

	input.Scan()
	fmt.Println("step 3. build")
	fmt.Println("make a local build")

	fmt.Println(" ex. go build")

	input.Scan()
	fmt.Println("step 4. test")
	fmt.Println("try it out!")

	fmt.Println(" ex. do-nothing.exe create")

	input.Scan()
}