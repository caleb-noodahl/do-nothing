package translator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/caleb-noodahl/do-nothing/models"
)

func TranslateGolang(steps []models.Step) (bytes.Buffer, error) {
	out := bytes.Buffer{}
	stepDefs := bytes.Buffer{}
	stepDefs.WriteString("\tsteps := []Prompt{}\n")
	out.WriteString(header())
	for _, step := range steps {
		out.WriteString(fmt.Sprintf("\ntype %s struct {}\n", step.NameCamelCase()))
		out.WriteString(fmt.Sprintf("func (%s %s) Step() {\n\t\tfmt.Println(\"%s\")\n}\n", strings.ToLower(step.Name[0:1]), step.NameCamelCase(), step.Description))
		stepDefs.WriteString(fmt.Sprintf("\tsteps = append(steps, %s{})\n", step.NameCamelCase()))
	}
	out.WriteString("func main() {\n\tinput := bufio.NewScanner(os.Stdin)\n")
	out.WriteString(stepDefs.String())
	out.WriteString("\tfor _, s := range steps {\n\t\ts.Step()\n\t\tinput.Scan()\n\t}\n}")
	return out, nil
}

func header() string {
	return `//this code was automatically generated. regenerating this file will overwrite any modifications!
package main
import (
	"fmt"
	"os"
	"bufio"
)	

type Prompt interface {
	Step()
}
	`
}
