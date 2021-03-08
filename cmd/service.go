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
	"log"
	"strings"

	"github.com/jholt96/jish-cli/templates"
	"github.com/spf13/cobra"
)

func createNewService(name string, serviceFlag string) {

	fileChannel := make(chan string, 1)

	serviceType := ""

	if serviceFlag != "" {

		switch strings.ToLower(serviceFlag) {
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

	templates.CreatService(name, serviceType, fileChannel)

	println(<-fileChannel)
}

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service <name> --type=<loadbalancer,clusterip,nodeport>",
	Short: "creates a service using the name provided",
	Long:  `Creates a service using the name provided. The default type will be clusterip but it can be changed using the 'type' flag`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		serviceType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = templates.ValidateFlagValue("service", serviceType)

		if err != nil {
			log.Fatalf(err.Error())
		}

		createNewService(args[0], serviceType)
	},
}

func init() {
	generateCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().String("type", "ClusterIP", "Determines the type of service to be created. valid values are: clusterip, loadbalancer, and nodeport")
}
