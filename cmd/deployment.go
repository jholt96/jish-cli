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

	templates "github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var configMap bool
var secret bool
var service bool
var all bool
var mode string

func createDeploymentYaml(args []string) {

	name := args[0]

	newDeploy := templates.CreateSimpleDeploymentYaml(name, "", "")

	newYaml, err := yaml.Marshal(newDeploy)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(string(newYaml))
}

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "create a deployment yaml file",
	Long: `create a deployment yaml file with the given name. This will have 3 different mode flags: simple, standard, complex. 
	There are 3 additional flags to create a configmap, secret, and service that are automatically added to the yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
		createDeploymentYaml(args)
	},
}

func init() {
	generateCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().String("configmap", "", "Create a configmap. takes the value of either env or mount to determine how to set up the template")
	deploymentCmd.Flags().String("secret", "", "Create a Secret. takes the value of either env or mount to determine how to set up the template")
	deploymentCmd.Flags().String("service", "", "Create a service. takes the value of either clusterip, nodeport,loadbalancer, or externalname to determine how to set up the template")

}
