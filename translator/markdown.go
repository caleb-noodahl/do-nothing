package translator

import (
	"bytes"
	"fmt"

	"github.com/caleb-noodahl/do-nothing/models"
)

func TranslateMarkdown(name string, steps []models.Step) (bytes.Buffer, error) {
	out := bytes.Buffer{}
	out.WriteString(fmt.Sprintf("# %s\n", name))
	for i, step := range steps {
		out.WriteString(fmt.Sprintf("## step %v\n\n%s\n%s\n", i+1, step.Name, step.Description))
		out.WriteString("\n```bash\n")
		for _, cmd := range step.Cmds {
			out.WriteString(fmt.Sprintf("%s\n", cmd))
		}
		out.WriteString("```\n\n")
	}
	return out, nil
}
