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
	templates "github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
)

func createNewSecret(name string) {

	fileChannel := make(chan string, 1)

	templates.CreateSecret(name, fileChannel)

	println(<-fileChannel)
}

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret <name>",
	Short: "Generates a Secret",
	Long:  `Generates an opaque secret to be used with kubectl `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createNewSecret(args[0])
	},
}

func init() {
	generateCmd.AddCommand(secretCmd)
}
