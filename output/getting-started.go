//this code was automatically generated. regenerating this file will overwrite any modifications!
package main
import (
	"fmt"
	"os"
	"bufio"
)	

type Prompt interface {
	Step()
}
	
type GitPull struct {}
func (g GitPull) Step() {
	fmt.Println("step 1")
	fmt.Println("pull the code back from github")
	fmt.Println(" ex. git clone https://github.com/caleb-noodahl/do-nothing.git")
}

type DependencyRestore struct {}
func (d DependencyRestore) Step() {
	fmt.Println("step 2")
	fmt.Println("download the dependencies specified in the go.mod file")
	fmt.Println(" ex. go mod download")
}

type Build struct {}
func (b Build) Step() {
	fmt.Println("step 3")
	fmt.Println("make a build creating a do-nothing.exe file")
	fmt.Println(" ex. go build .")
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	steps := []Prompt{}
	steps = append(steps, GitPull{})
	steps = append(steps, DependencyRestore{})
	steps = append(steps, Build{})
	for _, s := range steps {
		s.Step()
		input.Scan()
	}
}