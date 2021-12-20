package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/caleb-noodahl/do-nothing/models"
	"github.com/caleb-noodahl/do-nothing/utils"
	"github.com/caleb-noodahl/do-nothing/wizard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "a do something script to create a do nothing script",
	Long: `
	if you wish to compile your do-nothing script make sure your file and step names are compliant (no spaces)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		creator := wizard.StepWizard{}
		fmt.Printf("starting the do nothing document wizard. reply to a prompt with #stop to end\n")
		namePrompt := creator.CreateStepPrompt(wizard.ProtoPrompt{
			Label:    "do nothing document name",
			ErrorMsg: "enter the name of the do nothing script you wish to create",
		})
		creator.RemoveStep(0)
		doc, _ := namePrompt.Run()
		path := fmt.Sprintf("%s/%s", viper.GetString("output_path"), doc)
		cancel := false
		i := 1
		for !cancel {
			next := creator.CreateStepPrompt(wizard.ProtoPrompt{
				Label: fmt.Sprintf("step %v name", i),
			})
			step := models.Step{
				Sequence: i,
			}

			step.Name, step.Err = next.Run()
			if step.Err != nil {
				log.Fatal(step.Err)
			}
			if step.Name == "#stop" {
				cancel = true
				continue
			}

			next = creator.CreateStepPrompt(wizard.ProtoPrompt{
				Label: fmt.Sprintf("step %v text", i),
			})
			step.Text, step.Err = next.Run()
			if step.Err != nil {
				log.Fatal(step.Err)
			}

			next = creator.CreateStepPromptNoValidate(wizard.ProtoPrompt{
				Label: fmt.Sprintf("step %v cmds?", i),
			})

			cmd, err := next.Run()
			if err != nil {
				log.Fatal(err)
			}
			step.CMDs = strings.Split(cmd, ",")

			creator.AddStep(&step)
			i++
		}
		//so now we've got a list of steps that we'll have to compile into
		if err := utils.WriteJson(fmt.Sprintf("%s.dn", path), creator.Steps); err != nil {
			log.Fatal(err)
		}
		//generate the compiled file
		out := creator.CompileToGolang()
		//write the finished file
		if err := utils.Write(fmt.Sprintf("%s.go", path), out.Bytes()); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("output", "o", "", "output file name")
}
