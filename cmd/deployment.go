/*
Copyright © 2021 Josh Holt jholt96@live.com

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
	"log"

	templates "github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
)

func getFlagsandArgs(cmd *cobra.Command, args []string) (string, string, string, string, string) {
	configMap, configErr := cmd.Flags().GetString("configmap")
	if configErr != nil {
		log.Fatalf(configErr.Error())
	}
	secret, secretErr := cmd.Flags().GetString("secret")
	if secretErr != nil {
		log.Fatalf(secretErr.Error())
	}
	service, serviceErr := cmd.Flags().GetString("service")
	if serviceErr != nil {
		log.Fatalf(serviceErr.Error())
	}
	mode, modeErr := cmd.Flags().GetString("mode")
	if modeErr != nil {
		log.Fatalf(modeErr.Error())
	}

	name := args[0]

	return configMap, secret, service, mode, name
}

func createDeploymentYaml(configMap, secret, service, mode, name string) {

	if secret != "" {
		templates.CreateSecret((name + "-secret"))
	}
	if configMap != "" {
		templates.CreateConfigMap((name + "-configmap"))
	}

	switch mode {
	case "simple":
		templates.CreateSimpleDeploymentYaml(name, configMap, secret, service)
	case "standard":
		templates.CreateSimpleDeploymentYaml(name, configMap, secret, service)
	case "complex":
		templates.CreateSimpleDeploymentYaml(name, configMap, secret, service)
	default:
		templates.CreateSimpleDeploymentYaml(name, configMap, secret, service)
	}
}

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment <name>",
	Short: "create a deployment yaml file",
	Long: `create a deployment yaml file with the given name. This will have 3 different mode flags: simple, standard, complex. 
	There are 3 additional flags to create a configmap, secret, and service that are automatically added to the yaml`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configMap, secret, service, mode, name := getFlagsandArgs(cmd, args)

		createDeploymentYaml(configMap, secret, service, mode, name)
	},
}

func init() {
	generateCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().String("configmap", "", "Create a configmap. takes the value of either env or mount to determine how to set up the template")
	deploymentCmd.Flags().String("secret", "", "Create a Secret. takes the value of either env or mount to determine how to set up the template")
	deploymentCmd.Flags().String("service", "", "Create a service. takes the value of either clusterip, nodeport,loadbalancer, or externalname to determine how to set up the template")
	deploymentCmd.Flags().String("mode", "", "Determines the verboseness of the yaml file. ")

}
