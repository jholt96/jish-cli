package templates

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

const secretAPIVersion, configMapAPIVersion, serviceAPIVersion, namespace string = "v1", "v1", "v1", "<namespace>"
const configMapKind, secretKind, secretType, serviceKind string = "ConfigMap", "Secret", "Opaque", "Service"
const username, password, exampleData string = "user", "Pass", "world"
const port, targetPort, nodePort int = 80, 80, 8080

/*
creates the configmap template from the struct
It then converts the struct into yaml
Lastly it writes it to a yaml file using the naming convention of <name>-configmap.yaml
*/
func ValidateFlagValue(flag string, flagPassed string, possibleValues []string) (string, error) {

	for _, val := range possibleValues {
		if strings.ToLower(flagPassed) == strings.ToLower(val) {
			return flagPassed, nil
		}
	}

	return "", errors.New(("Incorrect value for flag: " + flag))
}

/*Creates a configmap template using the name used in the arg*/
func CreateConfigMap(name string) {

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
		fmt.Printf(writeFileErr.Error())
	} else {
		fmt.Printf("%s.yaml has been created in the current directory\n", name)
	}
}

/*
creates the secret template from the struct
It then converts the struct into yaml
Lastly it writes it to a yaml file using the naming convention of <name>-secret.yaml
*/
func CreateSecret(name string) {

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
		fmt.Printf(writeFileErr.Error())
	} else {
		fmt.Printf("%s.yaml has been created in the current directory\n", name)
	}
}

func CreatService(deploymentName string, serviceName, typeOfService string) {
	fmt.Printf("Creating Service...\n\n")

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
		fmt.Printf(writeFileErr.Error())
	} else {
		fmt.Printf("%s.yaml has been created in the current directory\n", serviceName)
	}

}
