package wizard

import (
	"github.com/manifoldco/promptui"
)

func promptTemplate() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
}

func noValidate(input string) error {
	return nil
}
