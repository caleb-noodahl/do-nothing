/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caleb-noodahl/do-nothing/models"
	"github.com/caleb-noodahl/do-nothing/translator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `creates a do nothing script for a yaml definition`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path         string
			out          string
			lang         string
			yamlContents []byte
			steps        []models.Step
		)

		path, _ = cmd.Flags().GetString("yaml")
		split := strings.Split(path, ".")
		name := split[len(split)-2]
		out, _ = cmd.Flags().GetString("out")
		lang, _ = cmd.Flags().GetString("lang")
		yamlContents, _ = os.ReadFile(path)

		fmt.Printf("reading from: %s\nwriting to: %s\n", path, out)
		if err := yaml.Unmarshal(yamlContents, &steps); err != nil {
			log.Fatal(err)
		}

		switch lang {
		case "go":
			data, err := translator.TranslateGolang(steps)
			if err != nil {
				log.Fatal(err)
			}
			if err := os.WriteFile(out, data.Bytes(), 0644); err != nil {
				log.Fatal(err)
			}
		case "markdown":
			data, err := translator.TranslateMarkdown(name, steps)
			if err != nil {
				log.Fatal(err)
			}
			if err := os.WriteFile(out, data.Bytes(), 0644); err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("yaml", "y", "", "yaml file to read from")
	createCmd.PersistentFlags().StringP("out", "o", "", "output filepath")
	createCmd.PersistentFlags().StringP("lang", "l", "", "language to translate to")
}
