package templates

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const secretAPIVersion, configMapAPIVersion, namespace string = "v1", "v1", "<namespace>"
const configMapKind, secretKind, secretType string = "ConfigMap", "Secret", "Opaque"
const username, password, exampleData string = "user", "Pass", "world"

/*
creates the configmap template from the struct
It then converts the struct into yaml
Lastly it writes it to a yaml file using the naming convention of <name>-configmap.yaml
*/
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
