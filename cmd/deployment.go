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
	"strings"

	templates "github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
)

func getFlagsandArgs(cmd *cobra.Command, args []string) (string, string, string, string, string) {
	configMap, configErr := cmd.Flags().GetString("configmap")
	if configErr != nil {
		log.Fatalf(configErr.Error())
	}

	_, err := templates.ValidateFlagValue("configmap", configMap)

	if err != nil {
		log.Fatalf(err.Error())
	}

	secret, secretErr := cmd.Flags().GetString("secret")
	if secretErr != nil {
		log.Fatalf(secretErr.Error())
	}
	_, err = templates.ValidateFlagValue("secret", secret)

	if err != nil {
		log.Fatalf(err.Error())
	}
	service, serviceErr := cmd.Flags().GetString("service")
	if serviceErr != nil {
		log.Fatalf(serviceErr.Error())
	}

	_, err = templates.ValidateFlagValue("service", service)

	if err != nil {
		log.Fatalf(err.Error())
	}

	mode, modeErr := cmd.Flags().GetString("mode")
	if modeErr != nil {
		log.Fatalf(modeErr.Error())
	}

	_, err = templates.ValidateFlagValue("mode", mode)

	if err != nil {
		log.Fatalf(err.Error())
	}

	name := args[0]

	return configMap, secret, service, mode, name
}

func createDeploymentYaml(configMap, secret, service, mode, name string) {
	count := 0

	fileChannel := make(chan string, 1)

	if secret != "" {
		count++
		go templates.CreateSecret((name + "-secret"), fileChannel)
	}
	if configMap != "" {
		count++
		go templates.CreateConfigMap((name + "-configmap"), fileChannel)
	}

	if service != "" {
		count++

		serviceType := ""

		switch strings.ToLower(service) {
		case "clusterip":
			serviceType = "ClusterIP"
		case "nodeport":
			serviceType = "NodePort"
		case "loadbalancer":
			serviceType = "LoadBalancer"
		default:
			serviceType = "ClusterIP"
		}
		go templates.CreatService(name, serviceType, fileChannel)
	}

	count++
	switch mode {
	case "simple":
		go templates.CreateSimpleDeploymentYaml(name, configMap, secret, service, fileChannel)
	case "standard":
		go templates.CreateStandardDeploymentYaml(name, configMap, secret, service, fileChannel)
	case "complex":
		go templates.CreateSimpleDeploymentYaml(name, configMap, secret, service, fileChannel)
	default:
		go templates.CreateStandardDeploymentYaml(name, configMap, secret, service, fileChannel)
	}

	for i := 0; i < count; i++ {
		println(<-fileChannel)
	}

	close(fileChannel)
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
