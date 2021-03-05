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
	"github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
)

// configmapCmd represents the configmap command
var configmapCmd = &cobra.Command{
	Use:   "configmap <name>",
	Short: "creates a configMap",
	Long:  `Creates a base configMap that can be used with kubectl`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templates.CreateConfigMap(args[0])
	},
}

func init() {
	generateCmd.AddCommand(configmapCmd)
}
