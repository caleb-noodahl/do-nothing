package wizard

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/caleb-noodahl/do-nothing/models"
	"github.com/manifoldco/promptui"
)

type StepWizard struct {
	Steps   []*models.Step
	Outputs []string
}

type ProtoPrompt struct {
	Label    string `json:"label"`
	ErrorMsg string `json:"err_msg"`
}

func (p *ProtoPrompt) Validate(input string) error {
	if len(input) <= 0 {
		return errors.New(p.ErrorMsg)
	}
	return nil
}

func (s *StepWizard) CreateStepPrompt(p ProtoPrompt) *promptui.Prompt {
	out := &promptui.Prompt{
		Label:       p.Label,
		Validate:    p.Validate,
		HideEntered: true,
	}
	return out
}

func (s *StepWizard) CreateStepPromptNoValidate(p ProtoPrompt) *promptui.Prompt {
	out := &promptui.Prompt{
		Label:       p.Label,
		Validate:    noValidate,
		HideEntered: true,
	}
	return out
}

func (s *StepWizard) AddStep(step *models.Step) int {
	s.Steps = append(s.Steps, step)
	return len(s.Steps)
}

func (s *StepWizard) RemoveStep(index int) {
	if len(s.Steps) > index {
		s.Steps = append(s.Steps[:index], s.Steps[index+1:]...)
	} else {
		s.Steps = []*models.Step{}
	}

}

func (s *StepWizard) CompileToGolang() bytes.Buffer {
	buf := bytes.Buffer{}
	buf.WriteString(`package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	`)

	for _, s := range s.Steps {
		buf.WriteString(fmt.Sprintf(`
	fmt.Println("step %v. %s")
	fmt.Println("%s")
`, s.Sequence, s.Name, s.Text))
		for _, cmd := range s.CMDs {
			buf.WriteString(fmt.Sprintf(`
	fmt.Println(" ex. %s")
`, cmd))
		}
		buf.WriteString(`
	input.Scan()`)
	}
	buf.WriteString(`
}`)
	return buf
}
