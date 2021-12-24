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
		fmt.Println("pull the code back from github")
}

type DependencyRestore struct {}
func (d DependencyRestore) Step() {
		fmt.Println("download the dependencies specified in the go.mod file")
}

type Build struct {}
func (b Build) Step() {
		fmt.Println("make a build creating a do-nothing.exe file")
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