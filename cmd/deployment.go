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

	"github.com/spf13/cobra"
)

var configMap bool
var secret bool
var service bool
var all bool
var mode string

func createDeploymentYaml(...int) {

}

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "create a deployment yaml file",
	Long: `create a deployment yaml file with the given name. This will have 3 different mode flags: simple, standard, complex. 
	There are flags to create a configmap, secret, service, or all of them that are automatically added to the yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
	},
}

func init() {
	generateCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().BoolP("configMap", "c", false, "create and mount configMap")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
