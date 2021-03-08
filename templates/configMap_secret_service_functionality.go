package templates

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

var configMapVals = []string{"env", "val", ""}
var secretVals = []string{"env", "val", ""}
var serviceVals = []string{"clusterip", "loadbalancer", "nodeport", ""}
var modeVals = []string{"simple", "standard", "complex", ""}

/*
creates the configmap template from the struct
It then converts the struct into yaml
Lastly it writes it to a yaml file using the naming convention of <name>-configmap.yaml
*/
func ValidateFlagValue(flag string, flagPassed string) (string, error) {
	var possibleValues []string

	if flag == "configmap" {
		possibleValues = configMapVals
	} else if flag == "secret" {
		possibleValues = secretVals
	} else if flag == "service" {
		possibleValues = serviceVals
	} else if flag == "mode" {
		possibleValues = modeVals
	}

	for _, val := range possibleValues {
		if strings.ToLower(flagPassed) == strings.ToLower(val) {
			return flagPassed, nil
		}
	}

	return "", errors.New(("Incorrect value for flag: " + flag))
}

/*Creates a configmap template using the name used in the arg*/
func CreateConfigMap(name string, ch chan<- string) {

	fmt.Printf("Creating ConfigMap...\n\n")

	configMap := &configMap{}

	configMap.ApiVersion = configMapAPIVersion
	configMap.Kind = configMapKind
	configMap.Metadata.Name = name
	configMap.Metadata.Namespace = namespace

	newConfigMapYaml, err := yaml.Marshal(configMap)

	if err != nil {
		log.Fatal(err.Error())
	}

	writeFileErr := ioutil.WriteFile((name + ".yaml"), newConfigMapYaml, 0755)

	if writeFileErr != nil {
		ch <- writeFileErr.Error()
	} else {
		ch <- name + ".yaml has been created in the current directory\n"
	}
}

/*
creates the secret template from the struct
It then converts the struct into yaml
Lastly it writes it to a yaml file using the naming convention of <name>-secret.yaml
*/
func CreateSecret(name string, ch chan<- string) {

	fmt.Printf("Creating Secret...\n\n")

	secret := &secret{}

	secret.ApiVersion = secretAPIVersion
	secret.Kind = secretKind
	secret.Metadata.Name = name
	secret.Metadata.Namespace = namespace
	secret.StringData.Username = username
	secret.StringData.Password = password
	secret.Type = secretType

	newSecretYaml, err := yaml.Marshal(secret)

	if err != nil {
		log.Fatal(err.Error())
	}

	writeFileErr := ioutil.WriteFile((name + ".yaml"), newSecretYaml, 0755)

	if writeFileErr != nil {
		ch <- writeFileErr.Error()
	} else {
		ch <- name + ".yaml has been created in the current directory\n"
	}
}

func CreatService(deploymentName string, typeOfService string, ch chan<- string) {
	fmt.Printf("Creating Service...\n\n")

	serviceName := deploymentName + "-service"

	service := &service{}

	service.APIVersion = serviceAPIVersion
	service.Kind = serviceKind
	service.Metadata.Name = serviceName
	service.Metadata.Namespace = namespace
	service.Spec.Ports = make([]struct {
		Port       int "yaml:\"port\""
		TargetPort int "yaml:\"targetPort\""
	}, 1)

	service.Spec.Selector.App = deploymentName

	service.Spec.Ports[0].Port = port
	service.Spec.Ports[0].TargetPort = targetPort
	service.Spec.Type = typeOfService

	newServiceYaml, err := yaml.Marshal(service)
	if err != nil {
		log.Fatal(err.Error())
	}
	writeFileErr := ioutil.WriteFile((serviceName + ".yaml"), newServiceYaml, 0755)

	if writeFileErr != nil {
		ch <- writeFileErr.Error()
	} else {
		ch <- serviceName + ".yaml has been created in the current directory\n"
	}

}
